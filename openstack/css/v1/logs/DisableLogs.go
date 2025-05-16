package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DisableLogs(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterID, "logs", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
