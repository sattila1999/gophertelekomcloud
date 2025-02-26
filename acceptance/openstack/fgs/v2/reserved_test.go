package v2

import (
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/reserved"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphListReserved(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to LIST FUNCGRAPH RESERVED INSTANCE CONFIGURATION")
	listReservedConfigs, err := reserved.ListReservedInstConfigs(client, reserved.ListConfigOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listReservedConfigs)

	t.Logf("Attempting to LIST FUNCGRAPH RESERVED INSTANCES")
	listReserved, err := reserved.ListReservedInst(client, reserved.ListOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listReserved)
}

func TestFunctionGraphReservedLifecycle(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	t.Logf("Attempting to UPDATE FUNCGRAPH RESERVED INSTANCES")
	updateResp, err := reserved.Update(client, reserved.UpdateOpts{
		FuncUrn: funcUrn,
		Count:   pointerto.Int(1),
	})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updateResp)
}
