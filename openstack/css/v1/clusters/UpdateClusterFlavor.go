package clusters

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/flavors"
)

const (
	DefaultNodeType = "ess"
)

type ClusterFlavorOptsBuilder interface{}

// ClusterFlavorOpts for updating the flavor of a cluster.
type ClusterFlavorOpts struct {
	// Whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	Flavor string `json:"-" required:"true"`
}

// ClusterNodeFlavorOpts for updating the flavor of a specific cluster node type.
type ClusterNodeFlavorOpts struct {
	// Whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	Flavor string `json:"-" required:"true"`
	// Type of the cluster node to modify.
	Type string `json:"type" required:"true"`
}

type apiFlavorOpts struct {
	// Whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	NewFlavorID string `json:"newFlavorId"`
}

// getFlavorIDByName retrieves the flavor ID for a given flavor name and node type.
func getFlavorIDByName(client *golangsdk.ServiceClient, flavorName, nodeType string) (string, error) {
	versions, err := flavors.List(client)
	if err != nil {
		return "", err
	}

	// Search for the flavor by name and node type.
	findFlavorByName := flavors.FindFlavor(versions, flavors.FilterOpts{
		FlavorName: flavorName,
		Type:       nodeType,
	})

	if findFlavorByName == nil {
		return flavorName, nil
	}

	return findFlavorByName.FlavorID, nil
}

// UpdateClusterFlavor updates the flavor of a cluster or a specific cluster node type.
func UpdateClusterFlavor(client *golangsdk.ServiceClient, clusterID, flavor string, opts ClusterFlavorOptsBuilder) error {
	var (
		url              string
		flavorID         string
		needCheckReplica bool
		err              error
	)

	switch options := opts.(type) {
	case ClusterFlavorOpts:
		// Update cluster flavor.
		url = client.ServiceURL("clusters", clusterID, "flavor")
		needCheckReplica = options.NeedCheckReplica
		flavorID, err = getFlavorIDByName(client, flavor, DefaultNodeType)
	case ClusterNodeFlavorOpts:
		// Update specific node type flavor.
		url = client.ServiceURL("clusters", clusterID, options.Type, "flavor")
		needCheckReplica = options.NeedCheckReplica
		flavorID, err = getFlavorIDByName(client, flavor, options.Type)
	default:
		return fmt.Errorf("invalid options type provided: %T", opts)
	}

	if err != nil {
		return err
	}

	// Construct the API payload.
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
