package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableLogBackupOpts struct {
	// These parameters are passed to the logs.EnableLogBackup function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency" required:"true"`
	// Indicates whether to enable automatic backup.
	AutoEnable bool `json:"autoEnable"`
	// BasePath is the obs path where the logs should be stored for the css cluster.
	BasePath string `json:"logBasePath"`
	// Bucket is the obs bucket name to store the logs for the css cluster.
	Bucket string `json:"logBucket" required:"true"`
	// Start time of automatic log backup.
	Period string `json:"period"`
}

type EnableLogIngestionOpts struct {
	// These parameters are passed to the logs.EnableLogIngestion function.
	// Prefix of the index for saving logs.
	IndexPrefix string `json:"index_prefix"`
	// Log retention duration.
	KeepDays int `json:"keep_days"`
	// ID of the target cluster where logs are saved.
	TargetClusterId string `json:"target_cluster_id"`
}

// EnableLogBackup function is used to enable the log backup switch based on EnableLogBackupOpts.
func EnableLogBackup(client *golangsdk.ServiceClient, clusterID string, opts EnableLogBackupOpts) error {
	b, err := build.RequestBody(opts, "")
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
	})
	return err
}

// EnableLogIngestion function is used to enable the log ingestion switch based on EnableLogIngestionOpts.
func EnableLogIngestion(client *golangsdk.ServiceClient, clusterID string, opts EnableLogIngestionOpts) error {
	b, err := build.RequestBody(opts, "")
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
	})
	return err
}
