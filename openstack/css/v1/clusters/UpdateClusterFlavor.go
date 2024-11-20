package clusters

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/flavors"
)

type ClusterFlavorOptsBuilder interface {
}

type ClusterFlavorOpts struct {
	// Indicates whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	Flavor string `json:"-" required:"true"`
}

type ClusterNodeFlavorOpts struct {
	// Indicates whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	Flavor string `json:"-" required:"true"`
	// Type of the cluster node to modify.
	Type string `json:"type" required:"true"`
}

type apiFlavorOpts struct {
	// Indicates whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	NewFlavorID string `json:"newFlavorId"`
}

func getFlavorIDByName(client *golangsdk.ServiceClient, flavorName string, nodeType string) (string, error) {
	var flavorID string
	versions, err := flavors.List(client)

	if err != nil {
		return "", err
	}

	findFlavorByName := flavors.FindFlavor(versions, flavors.FilterOpts{
		FlavorName: flavorName,
		Type:       nodeType,
	})

	if findFlavorByName == nil {
		flavorID = flavorName
	} else {
		flavorID = findFlavorByName.FlavorID
	}

	return flavorID, err
}

func UpdateClusterFlavor(client *golangsdk.ServiceClient, clusterID string, flavor string, opts ClusterFlavorOptsBuilder) error {
	var url string
	var flavorID string
	var needCheckReplica bool
	var err error

	switch options := opts.(type) {
	case ClusterFlavorOpts:
		url = client.ServiceURL("clusters", clusterID, "flavor")
		needCheckReplica = options.NeedCheckReplica
		flavorID, err = getFlavorIDByName(client, flavor, "ess")
		if err != nil {
			return err
		}
	case ClusterNodeFlavorOpts:
		url = client.ServiceURL("clusters", clusterID, options.Type, "flavor")
		needCheckReplica = options.NeedCheckReplica
		flavorID, err = getFlavorIDByName(client, flavor, options.Type)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid options type provided: %T", opts)
	}

	apiOpts := apiFlavorOpts{
		NeedCheckReplica: needCheckReplica,
		NewFlavorID:      flavorID,
	}

	b, err := build.RequestBody(apiOpts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
