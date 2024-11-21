package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddClusterNodes(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	nodeType := "ess-client"

	clusterID := getEnvVar("CSS_CLUSTER_ID")

	_, err = clusters.AddClusterNodes(client, clusterID, nodeType, clusters.AddNodesOpts{
		Flavor:     "css.xlarge.2",
		NodeSize:   1,
		VolumeType: "HIGH",
	})
	th.AssertNoErr(t, err)

	timeout := 1200

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
