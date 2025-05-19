package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangeLogConfigurationOpts struct {
	// Agency is the agency name used for the css cluster.
	// BasePath is the obs path where the logs should be stored for the css cluster.
	// Bucket is the obs bucket name to store the logs for the css cluster.
	// These objects are passed to the logs.UpdateLogs function.
	Agency   string `json:"agency" required:"true"`
	BasePath string `json:"logBasePath" required:"true"`
	Bucket   string `json:"logBucket" required:"true"`
}

// UpdateLogs function is used to change the log configurations of a csscluster,
func UpdateLogs(client *golangsdk.ServiceClient, clusterID string, opts ChangeLogConfigurationOpts) error {
	// UpdateLogs will change the cluster logging configurations based on ChangeLogConfigurationOpts.
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterID, "logs", "settings"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
