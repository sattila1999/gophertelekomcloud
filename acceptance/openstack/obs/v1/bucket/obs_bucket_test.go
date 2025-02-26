package bucket

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/obs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestObsBucketLifecycle(t *testing.T) {
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

	_, err = client.SetBucketEncryption(&obs.SetBucketEncryptionInput{
		Bucket: bucketName,
		BucketEncryptionConfiguration: obs.BucketEncryptionConfiguration{
			SSEAlgorithm: "kms",
		},
	})
	th.AssertNoErr(t, err)

	bucketHead, err := client.GetBucketMetadata(&obs.GetBucketMetadataInput{
		Bucket: bucketName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, bucketHead.FSStatus, obs.FSStatusDisabled)
	th.AssertEquals(t, bucketHead.Version, "3.0")
}

func TestObsParralelFSBucketLifecycle(t *testing.T) {
	client, err := clients.NewOBSClient()
	th.AssertNoErr(t, err)

	bucketName := strings.ToLower(tools.RandomString("obs-sdk-test", 5))

	_, err = client.CreateBucket(&obs.CreateBucketInput{
		Bucket:            bucketName,
		IsFSFileInterface: true,
	})
	t.Cleanup(func() {
		_, err = client.DeleteBucket(bucketName)
		th.AssertNoErr(t, err)
	})
	th.AssertNoErr(t, err)

	bucketHead, err := client.GetBucketMetadata(&obs.GetBucketMetadataInput{
		Bucket: bucketName,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, bucketHead.FSStatus, obs.FSStatusEnabled)
	th.AssertEquals(t, bucketHead.Version, "3.0")
}
