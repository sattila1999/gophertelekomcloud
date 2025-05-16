package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableLogsOptions struct {
	Agency   string `json:"agency" required:"true"`
	BasePath string `json:"logBasePath" required:"true"`
	Bucket   string `json:"logBucket" required:"true"`
}

func EnableLogs(client *golangsdk.ServiceClient, clusterID string, opts EnableLogsOptions) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "logs", "open")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
