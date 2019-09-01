package linux

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceLinuxSSHConnection() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the resource to connect to",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "root",
				Description: "The user that we should use for the connection",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     22,
				Description: "The port to connect to",
			},
			"password": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The password we should use for the connection",
				ConflictsWith: []string{"private_key"},
			},
			"private_key": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The contents of an SSH key to use for the connection",
				ConflictsWith: []string{"password"},
			},
		},
	}
}
