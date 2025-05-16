package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type LogConfiguration struct {
	ID         string `json:"id"`
	ClusteID   string `json:"clusterId"`
	ObsBucket  string `json:"obsBucket"`
	Agency     string `json:"agency"`
	UpdateAt   int    `json:"updateAt"`
	BasePath   string `json:"basePath"`
	AutoEnable bool   `json:"autoEnable"`
	Period     string `json:"period"`
	LogSwitch  bool   `json:"logSwitch"`
}

func GetLogConfiguration(client *golangsdk.ServiceClient, clusterID string) (*LogConfiguration, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "logs", "settings"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res LogConfiguration
	err = extract.IntoStructPtr(raw.Body, &res, "logConfiguration")
	return &res, err
}
