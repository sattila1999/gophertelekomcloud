package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type LogsOpts struct {
	// These parameters are passed to the logs.EnableLogs function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket"`
	// Index prefix for saving logs.
	IndexPrefix string `json:"index_prefix"`
	// Log retention duration.
	KeepDays int `json:"keep_days"`
	// Specifies the target cluster for saving logs.
	TargetClusterId string `json:"target_cluster_id"`
}

// EnableBaseLogs will enable log switch for the CSS cluster.
func EnableBaseLogs(client *golangsdk.ServiceClient, clusterID string, logSwitchOpts *LogsOpts) error {
	action := "base_log_collect"
	return enableLogs(client, clusterID, logSwitchOpts, &action)
}

// EnableRealTimeLogs will enable log ingestion for the CSS cluster.
func EnableRealTimeLogs(client *golangsdk.ServiceClient, clusterID string, logSwitchOpts *LogsOpts) error {
	action := "real_time_log_collect"
	return enableLogs(client, clusterID, logSwitchOpts, &action)
}

// EnableLogs function is used to enable the "Log Backup" or the "Log Ingestion" switch of a CSS cluster based on EnableLogsOpts.
func enableLogs(client *golangsdk.ServiceClient, clusterID string, logSwitchOpts *LogsOpts, action *string) error {
	b, err := build.RequestBody(*logSwitchOpts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "logs", "open")

	if action != nil && *action != "" {
		url += "?action=" + *action
	}

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
