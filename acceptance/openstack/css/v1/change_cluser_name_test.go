package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestChangeClusterNameWorkflow(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := os.Getenv("CSS_CLUSTER_ID")

	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` need to be defined")
	}

	opts := clusters.ChangeClusterNameOpts{
		DisplayName: tools.RandomString("changed-css-", 4),
	}
	err = clusters.ChangeClusterName(client, opts, clusterID)
	th.AssertNoErr(t, err)
}
