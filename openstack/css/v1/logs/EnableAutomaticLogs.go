package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableAutomaticLogsOptions struct {
	// This parameter passed to the logs.EnableAutomaticLogs function.
	// Period is the start time of a backup job.
	Period string `json:"period" required:"true"`
}

// EnableAutomaticLogs will enable the automatic log backup for the cluster based on EnableAutomaticLogsOptions.
func EnableAutomaticLogs(client *golangsdk.ServiceClient, clusterID string, opts EnableAutomaticLogsOptions) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "logs", "policy", "update")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
