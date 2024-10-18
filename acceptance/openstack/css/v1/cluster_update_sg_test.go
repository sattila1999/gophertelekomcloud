package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateSecurityGroup(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := os.Getenv("CSS_CLUSTER_ID")
	sgID := os.Getenv("SECURITY_GROUP_ID")
	if clusterID == "" || sgID == "" {
		t.Skip("Both CSS_CLUSTER_ID and SECURITY_GROUP_ID needs to be defined")
	}

	err = clusters.UpdateSecurityGroup(client, clusterID, clusters.SecurityGroupOpts{
		SecurityGroupID: sgID,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
