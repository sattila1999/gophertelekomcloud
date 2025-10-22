package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateBaseLogConfigurationOpts struct {
	// These parameters are passed to the logs.UpdateBaseLogs function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket"`
}

type UpdateRealTimeLogConfigurationOpts struct {
	// These parameters are passed to the logs.UpdateRealTimeLogs function.
	// Index prefix for storing logs.
	IndexPrefix string `json:"index_prefix"`
	// Log retention duration.
	KeepDays int `json:"keep_days"`
	// Specifies the target cluster for saving logs.
	TargetClusterId string `json:"target_cluster_id"`
}

// UpdateBaseLogs will change the base log collect configurations.
func UpdateBaseLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateBaseLogConfigurationOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	queryParam := getOpts{
		Action: "base_log_collect",
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "settings").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}

// UpdateRealTimeLogs will change the real time log collect configurations.
func UpdateRealTimeLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateRealTimeLogConfigurationOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	queryParam := getOpts{
		Action: "real_time_log_collect",
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "settings").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}

// // UpdateLogs will change the cluster logging configurations based on UpdateLogConfigurationOpts.
// func updateLogs(client *golangsdk.ServiceClient, clusterID string, opts UpdateLogConfigurationOpts, action *string) error {
// 	b, err := build.RequestBody(opts, "")
// 	if err != nil {
// 		return err
// 	}

// 	queryParam := getOpts{
// 		Action: *action,
// 	}

// 	url, err := golangsdk.NewURLBuilder().
// 		WithEndpoints("clusters", clusterID, "logs", "settings").
// 		WithQueryParams(&queryParam).Build()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
// 		OkCodes: []int{200},
// 	})
// 	return err
// }
