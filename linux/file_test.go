package linux

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-docker/docker"
	"github.com/terraform-providers/terraform-provider-local/local"
	"testing"
)

func TestAccFile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"docker": docker.Provider(),
			"linux":  Provider(),
			"local":  local.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: getFixture("file/basic/1_setup.tf"),
				Check:  nil,
			},
			{
				Config: getFixture("file/basic/2_create_file.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("linux_file.test_txt", "hash_md5", "47058241a059e3bd04cf358f958d6929"),
				),
			},
			{
				Config: getFixture("file/basic/3_change_contents.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("linux_file.test_txt", "hash_md5", "0869757ea0e2d93f81447990d5421526"),
				),
			},
		},
	})
}
