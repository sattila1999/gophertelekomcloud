package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/checkpoint"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/vaults"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackupLifecycle(t *testing.T) {
	if os.Getenv("RUN_CBR") == "" {
		t.Skip("too long to run in ci")
	}
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	vault, aOpts, optsVault, checkp := CreateCBR(t, client)
	th.AssertEquals(t, vault.ID, checkp.Vault.ID)
	th.AssertEquals(t, optsVault.Parameters.Description, checkp.ExtraInfo.Description)
	th.AssertEquals(t, optsVault.Parameters.Name, checkp.ExtraInfo.Name)
	th.AssertEquals(t, aOpts.Resources[0].Type, checkp.Vault.Resources[0].Type)

	checkpointGet, err := checkpoint.Get(client, checkp.ID)
	th.AssertNoErr(t, err)
	// Checks are disabled due to STO-10008 bug
	// th.AssertEquals(t, description, checkpointGet.ExtraInfo.Description)
	// th.AssertEquals(t, checkName, checkpointGet.ExtraInfo.Name)
	th.AssertEquals(t, "available", checkpointGet.Status)
	th.AssertEquals(t, vault.ID, checkpointGet.Vault.ID)
	th.AssertEquals(t, aOpts.Resources[0].Type, checkp.Vault.Resources[0].Type)

	allBackups, err := backups.List(client, backups.ListOpts{VaultID: vault.ID})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = backups.Delete(client, allBackups[0].ID)
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, waitForBackupDelete(client, 600, allBackups[0].ID))
	})

	bOpts := backups.RestoreBackupOpts{
		VolumeID: allBackups[0].ResourceID,
	}
	restoreErr := RestoreBackup(t, client, allBackups[0].ID, bOpts)
	th.AssertNoErr(t, restoreErr)
}

func TestBackupListing(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	// Create Vault for further backup
	opts := vaults.CreateOpts{
		Billing: &vaults.BillingCreate{
			ConsistentLevel: "crash_consistent",
			ObjectType:      "disk",
			ProtectType:     "backup",
			Size:            100,
		},
		Description: "gophertelemocloud testing vault",
		Name:        tools.RandomString("cbr-test-", 5),
		Resources:   []vaults.ResourceCreate{},
	}
	vault, err := vaults.Create(client, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, vaults.Delete(client, vault.ID))
	})

	// Create Volume
	volume := openstack.CreateVolume(t)
	t.Cleanup(func() {
		openstack.DeleteVolume(t, volume.ID)
	})

	// Associate server to the vault
	aOpts := vaults.AssociateResourcesOpts{
		Resources: []vaults.ResourceCreate{
			{
				ID:   volume.ID,
				Type: "OS::Cinder::Volume",
			},
		},
	}
	associated, err := vaults.AssociateResources(client, vault.ID, aOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(associated))

	// Create vault checkpoint
	optsVault := checkpoint.CreateOpts{
		VaultID: vault.ID,
		Parameters: checkpoint.CheckpointParam{
			Description: "go created backup",
			Incremental: true,
			Name:        tools.RandomString("go-checkpoint", 5),
		},
	}
	checkp := CreateCheckpoint(t, client, optsVault)
	th.AssertEquals(t, vault.ID, checkp.Vault.ID)
	th.AssertEquals(t, optsVault.Parameters.Description, checkp.ExtraInfo.Description)
	th.AssertEquals(t, optsVault.Parameters.Name, checkp.ExtraInfo.Name)
	th.AssertEquals(t, aOpts.Resources[0].Type, checkp.Vault.Resources[0].Type)

	listOpts := backups.ListOpts{VaultID: vault.ID}
	th.AssertNoErr(t, err)

	allBackups, err := backups.List(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(associated))

	cases := map[string]backups.ListOpts{
		"ID": {
			ID: allBackups[0].ID,
		},
		"Name": {
			Name: allBackups[0].Name,
		},
	}
	for _, cOpts := range cases {
		list, err := backups.List(client, cOpts)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, 1, len(list))
		th.AssertEquals(t, allBackups[0].ID, list[0].ID)

	}
	errBack := backups.Delete(client, allBackups[0].ID)
	th.AssertNoErr(t, errBack)
	th.AssertNoErr(t, waitForBackupDelete(client, 600, allBackups[0].ID))
}

func TestBackupSharingLifecycle(t *testing.T) {
	if os.Getenv("RUN_CBR") == "" {
		t.Skip("too long to run in ci")
	}
	destProjectID := os.Getenv("OS_PROJECT_ID_2")
	if destProjectID == "" {
		t.Skip("OS_PROJECT_ID_2 are mandatory for this test!")
	}
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	vault, aOpts, optsVault, checkp := CreateCBR(t, client)
	th.AssertEquals(t, vault.ID, checkp.Vault.ID)
	th.AssertEquals(t, optsVault.Parameters.Description, checkp.ExtraInfo.Description)
	th.AssertEquals(t, optsVault.Parameters.Name, checkp.ExtraInfo.Name)
	th.AssertEquals(t, aOpts.Resources[0].Type, checkp.Vault.Resources[0].Type)

	checkpointGet, err := checkpoint.Get(client, checkp.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "available", checkpointGet.Status)
	th.AssertEquals(t, vault.ID, checkpointGet.Vault.ID)
	th.AssertEquals(t, aOpts.Resources[0].Type, checkp.Vault.Resources[0].Type)

	allBackups, err := backups.List(client, backups.ListOpts{VaultID: vault.ID})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = backups.Delete(client, allBackups[0].ID)
		th.AssertNoErr(t, err)
		th.AssertNoErr(t, waitForBackupDelete(client, 600, allBackups[0].ID))
	})

	member, err := backups.AddSharingMember(client, allBackups[0].ID, backups.MembersOpts{
		Members: []string{destProjectID},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, member[0].DestProjectId, destProjectID)

	members, err := backups.ListSharingMembers(client, allBackups[0].ID, backups.ListMemberOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(members))

	t.Cleanup(func() {
		err = backups.DeleteSharingMember(client, allBackups[0].ID, destProjectID)
		th.AssertNoErr(t, err)
	})

	newCloud := os.Getenv("OS_CLOUD_2")
	if newCloud != "" {
		err = os.Setenv("OS_CLOUD", newCloud)
		th.AssertNoErr(t, err)
		_, err := clients.EnvOS.Cloud(newCloud)
		th.AssertNoErr(t, err)
		newClient, err := clients.NewCbrV3Client()
		th.AssertNoErr(t, err)

		// Create Vault for further accept member
		vault2Opts := vaults.CreateOpts{
			Billing: &vaults.BillingCreate{
				ConsistentLevel: "crash_consistent",
				ObjectType:      "disk",
				ProtectType:     "backup",
				Size:            100,
			},
			Description: "gophertelemocloud testing vault",
			Name:        tools.RandomString("cbr-test-", 5),
			Resources:   []vaults.ResourceCreate{},
		}
		vault2, err := vaults.Create(client, vault2Opts)
		th.AssertNoErr(t, err)
		t.Cleanup(func() {
			th.AssertNoErr(t, vaults.Delete(client, vault2.ID))
		})

		_, err = backups.UpdateSharingMember(newClient, destProjectID, backups.UpdateOpts{
			BackupID: allBackups[0].ID,
			Status:   "accepted",
			VaultId:  vault2.ID,
		})
		th.AssertNoErr(t, err)

		newShare, err := backups.GetSharingMember(client, allBackups[0].ID, destProjectID)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, "accepted", newShare.Status)
	}
}
