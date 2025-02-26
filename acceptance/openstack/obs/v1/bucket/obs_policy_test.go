package bucket

import (
	"fmt"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestObsPolicyLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket: bucketName,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	objectName := tools.RandomString("test-obs-", 5)

	_, err = client.PutObject(&obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    objectName,
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		_, err = client.DeleteObject(&obs.DeleteObjectInput{
			Bucket: bucketName,
			Key:    objectName,
		})
		th.AssertNoErr(t, err)
	})

	policy := fmt.Sprintf(
		`{
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "ID": [
          "*"
        ]
      },
      "Action": [
        "*"
      ],
      "Resource": [
        "%[1]s/*",
        "%[1]s"
      ]
    }
  ]
}`, bucketName)

	policyInput := &obs.SetBucketPolicyInput{
		Bucket: bucketName,
		Policy: policy,
	}
	_, err = client.SetBucketPolicy(policyInput)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		_, err = client.DeleteBucketPolicy(bucketName)
		th.AssertNoErr(t, err)
	})

	_, err = client.GetBucketPolicy(bucketName)
	th.AssertNoErr(t, err)
}
