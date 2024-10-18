package clusters

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ClusterFlavorOptsBuilder interface {
}

type ClusterFlavorOpts struct {
	// Indicates whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	NewFlavorID string `json:"newFlavorId" required:"true"`
}

type ClusterNodeFlavorOpts struct {
	// Indicates whether to verify replicas.
	NeedCheckReplica bool `json:"needCheckReplica"`
	// ID of the new flavor.
	NewFlavorID string `json:"newFlavorId" required:"true"`
	// Type of the cluster node to modify.
	NodeType string `json:"type" required:"true"`
}

func UpdateClusterFlavor(client *golangsdk.ServiceClient, clusterID string, opts ClusterFlavorOptsBuilder) error {
	url := ""

	switch options := opts.(type) {
	case ClusterFlavorOpts:
		url = client.ServiceURL("clusters", clusterID, "flavor")
	case ClusterNodeFlavorOpts:
		url = client.ServiceURL("clusters", clusterID, options.NodeType, "flavor")
	default:
		return fmt.Errorf("invalid options type provided: %T", opts)
	}

	b, err := build.RequestBody(opts, "")
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
