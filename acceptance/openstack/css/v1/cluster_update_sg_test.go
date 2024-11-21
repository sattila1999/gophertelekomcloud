package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateSecurityGroup(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := getEnvVar("CSS_CLUSTER_ID")
	sgID := getEnvVar("SECURITY_GROUP_ID")

	err = clusters.UpdateSecurityGroup(client, clusterID, clusters.SecurityGroupOpts{
		SecurityGroupID: sgID,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}
