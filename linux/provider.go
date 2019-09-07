package linux

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"linux_file": linuxFile(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"linux_ssh_connection": linuxDataSourceSSHConnection(),
		},
	}
}
