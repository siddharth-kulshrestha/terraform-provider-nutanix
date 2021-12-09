package nutanix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const (
	PERMISSIONNAME = "Delete_ACP"
	PERMISSINOUUID = "e6eae070-76e4-40ab-9ed9-53753ffd2739"
)

func TestAccNutanixPermissionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionDataSourceConfig(PERMISSINOUUID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "name", PERMISSIONNAME),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "operation", "delete"),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "fields.0.field_mode", "NONE"),
				),
			},
		},
	})
}

func testAccPermissionDataSourceConfig(uuid string) string {
	return fmt.Sprintf(`
data "nutanix_permission" "test" {
	permission_id = "%s"
}
`, uuid)
}

func TestAccNutanixPermissionDataSource_basicByName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionDataSourceConfigByName(PERMISSIONNAME),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "name", PERMISSIONNAME),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "operation", "delete"),
					resource.TestCheckResourceAttr(
						"data.nutanix_permission.test", "fields.0.field_mode", "NONE"),
				),
			},
		},
	})
}

func testAccPermissionDataSourceConfigByName(name string) string {
	return fmt.Sprintf(`
data "nutanix_permission" "test" {
	permission_name = "%s"
}
`, name)
}
