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
	flavor := os.Getenv("CSS_NEW_FLAVOR")
	if clusterID == "" || flavor == "" {
		t.Skip("CSS_CLUSTER_ID and CSS_NEW_FALVOR need to be defined")
	}

	err = clusters.UpdateClusterFlavor(client, clusterID, flavor, clusters.ClusterFlavorOpts{
		NeedCheckReplica: false,
		Flavor:           flavor,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestUpdateClusterNodeFlavor(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := os.Getenv("CSS_CLUSTER_ID")
	flavor := os.Getenv("CSS_NEW_FLAVOR")
	nodeType := os.Getenv("CSS_NODE_TYPE")
	if clusterID == "" || flavor == "" || nodeType == "" {
		t.Skip("CSS_CLUSTER_ID, CSS_NEW_FLAVOR, and CSS_NODE_TYPE need to be defined")
	}

	err = clusters.UpdateClusterFlavor(client, clusterID, flavor, clusters.ClusterNodeFlavorOpts{
		NeedCheckReplica: false,
		Flavor:           flavor,
		Type:             nodeType,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
