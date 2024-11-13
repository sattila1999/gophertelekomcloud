package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAddClusterNodes(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	nodeType := "ess-client"

	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined")
	}

	_, err = clusters.AddClusterNodes(client, clusterID, nodeType, clusters.AddNodesOpts{
		// css.medium.8: ced8d1a7-eff8-4e30-a3de-cd9578fd518f
		// css.xlarge.2: d9dc06ae-b9c4-4ef4-acd8-953ef4205e27
		Flavor:     "d9dc06ae-b9c4-4ef4-acd8-953ef4205e27",
		NodeSize:   1,
		VolumeType: "HIGH",
	})
	th.AssertNoErr(t, err)

	timeout := 1200

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
