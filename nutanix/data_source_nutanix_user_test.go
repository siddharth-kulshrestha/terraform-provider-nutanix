package nutanix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccNutanixUserDataSource_basic(t *testing.T) {
	principalName := "user3@qa.nucalm.io"
	expectedDisplayName := "user3"
	directoryServiceUUID := "3d227657-118d-4819-b4c0-432d54fd0687"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNutanixUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(principalName, directoryServiceUUID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_user.user", "display_name", expectedDisplayName),
					resource.TestCheckResourceAttrSet("data.nutanix_user.user", "directory_service_user.#"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig(pn, dsuuid string) string {
	return fmt.Sprintf(`
resource "nutanix_user" "user" {
	directory_service_user {
		user_principal_name = "%s"
		directory_service_reference {
		uuid = "%s"
		}
	}
}

data "nutanix_user" "user" {
	user_id = nutanix_user.user.id
}
`, pn, dsuuid)
}

func TestAccNutanixUserDataSource_byName(t *testing.T) {
	principalName := "user4@qa.nucalm.io"
	expectedDisplayName := "user4"
	directoryServiceUUID := "3d227657-118d-4819-b4c0-432d54fd0687"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNutanixUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfigByName(principalName, directoryServiceUUID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_user.user", "display_name", expectedDisplayName),
					resource.TestCheckResourceAttrSet("data.nutanix_user.user", "directory_service_user.#"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfigByName(pn, dsuuid string) string {
	return fmt.Sprintf(`
resource "nutanix_user" "user" {
	directory_service_user {
		user_principal_name = "%s"
		directory_service_reference {
		uuid = "%s"
		}
	}
}

data "nutanix_user" "user" {
	user_name = nutanix_user.user.name
}
`, pn, dsuuid)
}
