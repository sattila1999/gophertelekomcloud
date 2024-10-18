package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateClusterFlavor(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined")
	}

	err = clusters.UpdateClusterFlavor(client, clusterID, clusters.ClusterFlavorOpts{
		NeedCheckReplica: false,
		// css.medium.8: ced8d1a7-eff8-4e30-a3de-cd9578fd518f
		// css.xlarge.2: d9dc06ae-b9c4-4ef4-acd8-953ef4205e27
		NewFlavorID: "d9dc06ae-b9c4-4ef4-acd8-953ef4205e27",
	})
	th.AssertNoErr(t, err)

	timeout := 1200

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestUpdateClusterNodeFlavor(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined")
	}

	err = clusters.UpdateClusterFlavor(client, clusterID, clusters.ClusterNodeFlavorOpts{
		NeedCheckReplica: false,
		NewFlavorID:      "ced8d1a7-eff8-4e30-a3de-cd9578fd518f",
		NodeType:         "ess",
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
