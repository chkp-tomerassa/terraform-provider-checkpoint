package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"strings"
	"testing"
)

func TestAccCheckpointManagementVoipDomainSipProxy_basic(t *testing.T) {

	var voipDomainSipProxyMap map[string]interface{}
	resourceName := "checkpoint_management_voip_domain_sip_proxy.test"
	objName := "tfTestManagementVoipDomainSipProxy_" + acctest.RandString(6)

	context := os.Getenv("CHECKPOINT_CONTEXT")
	if context != "web_api" {
		t.Skip("Skipping management test")
	} else if context == "" {
		t.Skip("Env CHECKPOINT_CONTEXT must be specified to run this acc test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckpointManagementVoipDomainSipProxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccManagementVoipDomainSipProxyConfig(objName, "new_group_sip", "new_host_sip"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCheckpointManagementVoipDomainSipProxyExists(resourceName, &voipDomainSipProxyMap),
					testAccCheckCheckpointManagementVoipDomainSipProxyAttributes(&voipDomainSipProxyMap, objName, "new_group_sip", "new_host_sip"),
				),
			},
		},
	})
}

func testAccCheckpointManagementVoipDomainSipProxyDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*checkpoint.ApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "checkpoint_management_voip_domain_sip_proxy" {
			continue
		}
		if rs.Primary.ID != "" {
			res, _ := client.ApiCall("show-voip-domain-sip-proxy", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
			if res.Success {
				return fmt.Errorf("VoipDomainSipProxy object (%s) still exists", rs.Primary.ID)
			}
		}
		return nil
	}
	return nil
}

func testAccCheckCheckpointManagementVoipDomainSipProxyExists(resourceTfName string, res *map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceTfName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceTfName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VoipDomainSipProxy ID is not set")
		}

		client := testAccProvider.Meta().(*checkpoint.ApiClient)

		response, err := client.ApiCall("show-voip-domain-sip-proxy", map[string]interface{}{"uid": rs.Primary.ID}, client.GetSessionID(), true, client.IsProxyUsed())
		if !response.Success {
			return err
		}

		*res = response.GetData()

		return nil
	}
}

func testAccCheckCheckpointManagementVoipDomainSipProxyAttributes(voipDomainSipProxyMap *map[string]interface{}, name string, endpointsDomain string, installedAt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		voipDomainSipProxyName := (*voipDomainSipProxyMap)["name"].(string)
		if !strings.EqualFold(voipDomainSipProxyName, name) {
			return fmt.Errorf("name is %s, expected %s", name, voipDomainSipProxyName)
		}
		voipDomainSipProxyEndpointsDomain := (*voipDomainSipProxyMap)["endpoints-domain"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainSipProxyEndpointsDomain, endpointsDomain) {
			return fmt.Errorf("endpointsDomain is %s, expected %s", endpointsDomain, voipDomainSipProxyEndpointsDomain)
		}
		voipDomainSipProxyInstalledAt := (*voipDomainSipProxyMap)["installed-at"].(map[string]interface{})["name"].(string)
		if !strings.EqualFold(voipDomainSipProxyInstalledAt, installedAt) {
			return fmt.Errorf("installedAt is %s, expected %s", installedAt, voipDomainSipProxyInstalledAt)
		}
		return nil
	}
}

func testAccManagementVoipDomainSipProxyConfig(name string, endpointsDomain string, installedAt string) string {
	return fmt.Sprintf(`
resource "checkpoint_management_group" "group1" {
  name = "%s"
}

resource "checkpoint_management_host" "host1" {
  name = "%s"
  ipv4_address = "192.0.2.8"
}

resource "checkpoint_management_voip_domain_sip_proxy" "test" {
        name = "%s"
        endpoints_domain = "${checkpoint_management_group.group1.name}"
        installed_at = "${checkpoint_management_host.host1.name}"
}
`, endpointsDomain, installedAt, name)
}
