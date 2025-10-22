package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// type getOpts struct {
// 	Action string `q:"action"`
// }

// DisableBaseLogs will disable log switch.
func DisableBaseLogs(client *golangsdk.ServiceClient, clusterID string) error {
	action := "base_log_collect"
	return disableLogs(client, clusterID, &action)
}

// DisableRealTimeLogs will disable log ingestion.
func DisableRealTimeLogs(client *golangsdk.ServiceClient, clusterID string) error {
	action := "real_time_log_collect"
	return disableLogs(client, clusterID, &action)
}

// DisableLogs function will disable the log option for a CSS cluster.
func disableLogs(client *golangsdk.ServiceClient, clusterID string, action *string) error {
	queryParam := getOpts{
		Action: *action,
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "close").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
