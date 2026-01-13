package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// RunManualBackup function will run back up of logs, if log backup is enabled for a cluster
func RunManualBackup(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Post(client.ServiceURL("clusters", clusterID, "logs", "collect"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}