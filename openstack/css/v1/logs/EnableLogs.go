package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type BaseLogOpts struct {
	// These parameters are passed to the logs.EnableBaseLogs function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket"`
}

type RealTimeLogOpts struct {
	// These parameters are passed to the logs.EnableRealTimeLogs function.
	// Index prefix for saving logs.
	IndexPrefix string `json:"index_prefix"`
	// Log retention duration.
	KeepDays int `json:"keep_days"`
	// Specifies the target cluster for saving logs.
	TargetClusterId string `json:"target_cluster_id"`
}

// type getOpst struct {
// 	Action string `q:"action"`
// }

// EnableBaseLogs will enable log switch for the CSS cluster.
func EnableBaseLogs(client *golangsdk.ServiceClient, clusterID string, logSwitchOpts *BaseLogOpts) error {
	b, err := build.RequestBody(*logSwitchOpts, "")
	if err != nil {
		return err
	}

	queryParam := getOpts{
		Action: "base_log_collect",
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "open").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}

// EnableRealTimeLogs will enable log ingestion for the CSS cluster.
func EnableRealTimeLogs(client *golangsdk.ServiceClient, clusterID string, logSwitchOpts *RealTimeLogOpts) error {
	b, err := build.RequestBody(*logSwitchOpts, "")
	if err != nil {
		return err
	}

	queryParam := getOpts{
		Action: "real_time_log_collect",
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "open").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}

// // EnableLogs function is used to enable the "Log Backup" or the "Log Ingestion" switch of a CSS cluster based on EnableLogsOpts.
// func enableLogs(client *golangsdk.ServiceClient, clusterID string, logSwitchOpts *LogsOpts, action *string) error {
// 	b, err := build.RequestBody(*logSwitchOpts, "")
// 	if err != nil {
// 		return err
// 	}

// 	queryParam := getOpts{
// 		Action: *action,
// 	}

// 	url, err := golangsdk.NewURLBuilder().
// 		WithEndpoints("clusters", clusterID, "logs", "open").
// 		WithQueryParams(&queryParam).Build()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
// 		OkCodes: []int{200},
// 		MoreHeaders: map[string]string{
// 			"Content-Type": "application/json",
// 		},
// 	})

// 	return err
// }
