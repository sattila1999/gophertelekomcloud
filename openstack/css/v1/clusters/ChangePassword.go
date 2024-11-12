package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type ChangePasswordOpts struct {
	// DisplayName contains options for new name
	// This object is passed to the snapshots.ChangeClusterName function.
	NewPassword string `json:"newpassword" required:"true"`
}

func ChangePassword(client *golangsdk.ServiceClient, opts ChangePasswordOpts, clusterId string) (err error) {
	// ChangeClusterName will change cluster name based on ChangeClusterNameOpts
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "password/reset"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
