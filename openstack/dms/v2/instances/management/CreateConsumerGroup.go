package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateConsumerGroupOpts struct {
	// Consumer group name.
	GroupName string `json:"group_name" required:"true"`
	// Consumer group description.
	// Minimum: 0
	// Maximum: 200
	Description string `json:"group_desc,omitempty"`
}

// CreateConsumerGroup is used to create a consumer group.
// Send POST /v2/{project_id}/kafka/instances/{instance_id}/group
func CreateConsumerGroup(client *golangsdk.ServiceClient, instanceId string, opts CreateConsumerGroupOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("kafka", "instances", instanceId, "group"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
