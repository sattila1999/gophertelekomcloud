package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type ChangeClusterNameOpts struct {
	// DisplayName contains options for new name
	// This object is passed to the snapshots.ChangeClusterName function.
	DisplayName string `json:"displayName" required:"true"`
}

func ChangeClusterName(client *golangsdk.ServiceClient, opts ChangeClusterNameOpts, clusterId string) (err error) {
	// ChangeClusterName will change cluster name based on ChangeClusterNameOpts
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "changename"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
