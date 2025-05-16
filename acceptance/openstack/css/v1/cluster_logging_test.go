package v1

import (
	"fmt"
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCSSLogging(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	agency := clients.EnvOS.GetEnv("AGENCY_NAME")
	if agency == "" {
		t.Skipf("OS_AGENCY_NAME is required for this test")
	}
	bucketName := clients.EnvOS.GetEnv("BUCKET_NAME")
	if bucketName == "" {
		t.Skipf("OS_BUCKET_NAME is required for this test")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := logs.GetLogConfiguration(client, clusterID)
	th.AssertNoErr(t, err)
	// log.Print("Creating cluster, ID: ", got)
	if got.LogSwitch {
		fmt.Print("The logs are already enabled.")
		// err = logs.DisableLogs(client, clusterID)
		// th.AssertNoErr(t, err)

		// print("Cluster logging disabled.")

	} else {

		log.Println("The logs are not enabled.")

		basicOpts := logs.EnableLogsOptions{
			Agency:   agency,
			Bucket:   bucketName,
			BasePath: "css/log",
		}

		err = logs.EnableLogs(client, clusterID, basicOpts)
		th.AssertNoErr(t, err)

		log.Println("Cluster logging enabled.")

		th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

		err = logs.DisableLogs(client, clusterID)
		th.AssertNoErr(t, err)

		log.Println("Cluster logging disabled.")

	}

}
