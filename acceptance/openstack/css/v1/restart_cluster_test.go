package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRestartClusterWorkflow(t *testing.T) {

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	cssCluserID := os.Getenv("CSS_CLUSTER_ID")

	if cssCluserID == "" {
		t.Skip("`CSS_CLUSTER_ID` need to be defined")
	}

	err = clusters.RestartCluster(client, cssCluserID)
	th.AssertNoErr(t, err)
}
