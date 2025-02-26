package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Rename function is used to modify the name of a DDM instance.
// name is the name of a DDM instance, which:
// Can include 4 to 64 characters.
// Must start with a letter.
// Can contain only letters, digits, and hyphens (-).
func Rename(client *golangsdk.ServiceClient, instanceId string, name string) (*string, error) {
	b, err := build.RequestBody(name, "name")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/instances/{instance_id}/modify-name
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "modify-name"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ModifyNameResponse
	err = extract.Into(raw.Body, &res)
	return &res.Name, err
}

type ModifyNameResponse struct {
	Name string `json:"name"`
}
