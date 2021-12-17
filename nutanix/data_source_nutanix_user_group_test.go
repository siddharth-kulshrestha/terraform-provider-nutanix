package nutanix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccNutanixUserGroupDataSource_basic(t *testing.T) {
	distinguishedName := "cn=sspadmins,cn=users,dc=qa,dc=nucalm,dc=io"
	displayName := "sspadmins"
	uuid := "a8df1950-f144-4716-bc7f-274aeffa8449"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserGroupDataSourceConfig(uuid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "display_name", displayName),
					resource.TestCheckResourceAttrSet("data.nutanix_user_group.test", "directory_service_user_group.#"),
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "directory_service_user_group.0.distinguished_name", distinguishedName),
				),
			},
		},
	})
}

func testAccUserGroupDataSourceConfig(uuid string) string {
	return fmt.Sprintf(`
data "nutanix_user_group" "test" {
	user_group_id = "%s"
}
`, uuid)
}
func TestAccNutanixUserGroupDataSource_ByName(t *testing.T) {
	distinguishedName := "cn=sspadmins,cn=users,dc=qa,dc=nucalm,dc=io"
	displayName := "sspadmins"
	uuid := "a8df1950-f144-4716-bc7f-274aeffa8449"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserGroupDataSourceConfigByName(displayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "display_name", displayName),
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "id", uuid),
					resource.TestCheckResourceAttrSet("data.nutanix_user_group.test", "directory_service_user_group.#"),
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "directory_service_user_group.0.distinguished_name", distinguishedName),
				),
			},
		},
	})
}

func testAccUserGroupDataSourceConfigByName(dn string) string {
	return fmt.Sprintf(`
data "nutanix_user_group" "test" {
	user_group_name = "%s"
}
`, dn)
}

func TestAccNutanixUserGroupDataSource_ByDistinguishedName(t *testing.T) {
	distinguishedName := "cn=sspadmins,cn=users,dc=qa,dc=nucalm,dc=io"
	displayName := "sspadmins"
	uuid := "a8df1950-f144-4716-bc7f-274aeffa8449"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserGroupDataSourceConfigByDistinguishedName(distinguishedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "display_name", displayName),
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "id", uuid),
					resource.TestCheckResourceAttrSet("data.nutanix_user_group.test", "directory_service_user_group.#"),
					resource.TestCheckResourceAttr(
						"data.nutanix_user_group.test", "directory_service_user_group.0.distinguished_name", distinguishedName),
				),
			},
		},
	})
}

func testAccUserGroupDataSourceConfigByDistinguishedName(dn string) string {
	return fmt.Sprintf(`
data "nutanix_user_group" "test" {
	user_group_distinguished_name = "%s"
}
`, dn)
}
