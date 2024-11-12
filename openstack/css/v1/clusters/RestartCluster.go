package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func RestartCluster(client *golangsdk.ServiceClient, clusterId string) (err error) {
	_, err = client.Post(client.ServiceURL("clusters", clusterId, "restart"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
