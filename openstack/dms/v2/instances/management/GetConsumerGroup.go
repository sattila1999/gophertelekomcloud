package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetConsumerGroup is used to query a specific consumer group.
// Send GET /v2/{project_id}/instances/{instance_id}/groups/{group}
func GetConsumerGroup(client *golangsdk.ServiceClient, instanceId, groupId string) (*GetConsumerGropusResp, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "groups", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetConsumerGropusResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetConsumerGropusResp struct {
	// Consumer group information.
	Group Group `json:"group"`
}
