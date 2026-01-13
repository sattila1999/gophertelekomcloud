package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// TestTargetClusterConnection function will test the connectivity between two clusters.
func TestTargetClusterConnection(client *golangsdk.ServiceClient, clusterID string, targetClusterID string) error {
	body := map[string]interface{}{
		"target_cluster_id": targetClusterID,
	}
	b, err := build.RequestBody(body, "")
	if err != nil {
		return err
	}
	url := client.ServiceURL("clusters", clusterID, "logs", "connectivity")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		return err
	}
	return nil
}
