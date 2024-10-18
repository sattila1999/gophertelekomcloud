package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestUpdateSecurityModeEnableAll(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)
	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined.")
	}

	cssCluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	httpsEnable := cssCluster.HttpsEnabled
	authEnable := cssCluster.AuthorityEnabled

	if httpsEnable == false && authEnable == false {
		httpsEnable = !httpsEnable
		authEnable = !authEnable
	} else {
		t.Skip("The HTTPS and the Authority is already enabled.")
	}

	var adminPWD string = "P4s77gfhz!+%"

	err = clusters.UpdateSecurityMode(client, clusterID, clusters.SecurityModeOpts{
		AuthorityEnable: authEnable,
		AdminPwd:        adminPWD,
		HttpsEnable:     httpsEnable,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterOperationSucces(client, clusterID, timeout))

}

func TestUpdateSecurityModeDisableAll(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)
	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined.")
	}

	cssCluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	httpsEnable := cssCluster.HttpsEnabled
	authEnable := cssCluster.AuthorityEnabled

	if httpsEnable == true && authEnable == true {
		httpsEnable = !httpsEnable
		authEnable = !authEnable
	} else {
		t.Skip("The HTTPS and the Authority is already disabled.")
	}

	err = clusters.UpdateSecurityMode(client, clusterID, clusters.SecurityModeOpts{
		AuthorityEnable: authEnable,
		HttpsEnable:     httpsEnable,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterOperationSucces(client, clusterID, timeout))

}

func TestUpdateSecurityModeEnableHttps(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)
	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined.")
	}

	cssCluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	httpsEnable := cssCluster.HttpsEnabled

	if httpsEnable == false {
		httpsEnable = !httpsEnable
	} else {
		t.Skip("HTTPS is already enabled.")
	}

	var adminPWD string = "P4s77gfhz!+%"
	var authEnable bool = true

	err = clusters.UpdateSecurityMode(client, clusterID, clusters.SecurityModeOpts{
		AuthorityEnable: authEnable,
		AdminPwd:        adminPWD,
		HttpsEnable:     httpsEnable,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterOperationSucces(client, clusterID, timeout))

}

func TestUpdateSecurityModeEnableAuthority(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)
	clusterID := os.Getenv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("CSS_CLUSTER_ID needs to be defined.")
	}

	cssCluster, err := clusters.Get(client, clusterID)
	th.AssertNoErr(t, err)

	authEnable := cssCluster.AuthorityEnabled

	if authEnable == false {
		authEnable = !authEnable
	} else {
		t.Skip("Authority is already enabled.")
	}

	var adminPWD string = "P4s77gfhz!+%"
	var httpsEnable bool = false

	err = clusters.UpdateSecurityMode(client, clusterID, clusters.SecurityModeOpts{
		AuthorityEnable: authEnable,
		AdminPwd:        adminPWD,
		HttpsEnable:     httpsEnable,
	})
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForClusterOperationSucces(client, clusterID, timeout))

}
