package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type LogBackupResp struct {
	// The agency name.
	Agency string `json:"agency"`
	// Indicates whether to enable automatic backup.
	AutoEnable bool `json:"autoEnable"`
	// Storage path of backup logs in the OBS bucket.
	BasePath string `json:"basePath"`
	// CSS cluster ID.
	ClusterID string `json:"clusterId"`
	// Log backup ID.
	ID string `json:"id"`
	// Indicates whether to enable the log function.
	LogSwitch bool `json:"logSwitch"`
	// The bucket where the logs should be stored.
	ObsBucket string `json:"obsBucket"`
	// Start time of automatic log backup.
	Period string `json:"period"`
	// Update time.
	UpdateAt int `json:"updateAt"`
}

type LogIngestionResp struct {
	// CSS cluster ID.
	ClusterID string `json:"clusterId"`
	// Start time of a real-time log collection task.
	CreateAt int64 `json:"createAt"`
	// Log backup ID.
	ID string `json:"id"`
	// Prefix of the index for saving logs.
	IndexPrefix string `json:"indexPrefix"`
	// Log retention duration.
	KeepDays int `json:"keepDays"`
	// Status of a real-time log collection task.
	Status string `json:"status"`
	// ID of the target cluster where logs are saved.
	TargetClusterId string `json:"targetClusterId"`
	// Update time.
	UpdateAt int `json:"updateAt"`
}

type RealTimeLogConfiguration struct {
	// Log backup ID.
	ID string `json:"id"`
	// CSS cluster ID.
	ClusterID string `json:"clusterId"`
	// Prefix of the index for saving logs.
	IndexPrefix string `json:"indexPrefix"`
	// Log retention duration.
	KeepDays int `json:"keepDays"`
	// ID of the target cluster where logs are saved.
	TargetClusterId string `json:"targetClusterId"`
	// Status of a real-time log collection task.
	Status string `json:"status"`
	// Start time of a real-time log collection task.
	CreateAt int64 `json:"createAt"`
	// Update time.
	UpdateAt int `json:"updateAt"`
}

// GetConfiguration function will query the details of CSS cluster logging and returns a LogConfiguration object.
func GetLogBackupConfiguration(client *golangsdk.ServiceClient, clusterID string) (*LogBackupResp, error) {
	queryParam := getOpts{
		Action: "base_log_collect",
	}
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "settings").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res LogBackupResp
	err = extract.IntoStructPtr(raw.Body, &res, "logConfiguration")
	return &res, err
}

// GetRealTimeConfiguration function will query the details of CSS cluster logging and returns a LogConfiguration object.
func GetLogIngestionConfiguration(client *golangsdk.ServiceClient, clusterID string) (*LogIngestionResp, error) {

	queryParam := getOpts{
		Action: "real_time_log_collect",
	}
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "settings").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res LogIngestionResp
	err = extract.IntoStructPtr(raw.Body, &res, "realTimeLogCollectRecord")
	return &res, err
}
