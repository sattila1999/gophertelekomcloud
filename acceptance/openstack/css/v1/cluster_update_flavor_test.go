package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateClusterFlavor(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := getEnvVar("CSS_CLUSTER_ID")
	flavor := getEnvVar("CSS_NEW_FLAVOR")

	err = clusters.UpdateClusterFlavor(client, clusterID, flavor, clusters.ClusterFlavorOpts{
		NeedCheckReplica: false,
		Flavor:           flavor,
	})
	th.AssertNoErr(t, err)

	timeout := 600
	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestUpdateClusterNodeFlavor(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := getEnvVar("CSS_CLUSTER_ID")
	flavor := getEnvVar("CSS_NEW_FLAVOR")
	nodeType := getEnvVar("CSS_NODE_TYPE")

	err = clusters.UpdateClusterFlavor(client, clusterID, flavor, clusters.ClusterNodeFlavorOpts{
		NeedCheckReplica: false,
		Flavor:           flavor,
		Type:             nodeType,
	})
	th.AssertNoErr(t, err)

	timeout := 600
	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
