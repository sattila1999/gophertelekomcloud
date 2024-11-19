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
	flavorID := os.Getenv("CSS_NEW_FLAVOR_ID")
	if clusterID == "" || flavorID == "" {
		t.Skip("CSS_CLUSTER_ID and CSS_NEW_FALVOR_ID need to be defined")
	}

	err = clusters.UpdateClusterFlavor(client, clusterID, flavorID, clusters.ClusterFlavorOpts{
		NeedCheckReplica: false,
		FlavorID:         flavorID,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestUpdateClusterNodeFlavor(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := os.Getenv("CSS_CLUSTER_ID")
	flavorID := os.Getenv("CSS_NEW_FLAVOR_ID")
	if clusterID == "" || flavorID == "" {
		t.Skip("CSS_CLUSTER_ID and CSS_FLAVOR_ID need to be defined")
	}

	err = clusters.UpdateClusterFlavor(client, clusterID, flavorID, clusters.ClusterNodeFlavorOpts{
		NeedCheckReplica: false,
		FlavorID:         flavorID,
		Type:             "ess-master",
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
