package linux

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sam-myers/terraform-provider-linux/linux/manager"
	"github.com/sam-myers/terraform-provider-linux/linux/sshconnection"
)

var linuxDataSourceSSHConnectionSchema = map[string]*schema.Schema{
	"host": {
		Type:     schema.TypeString,
		Required: true,
	},
	"user": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "root",
	},
	"port": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  22,
	},

	"password": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"private_key": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"certificate": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"host_key": {
		Type:     schema.TypeString,
		Optional: true,
	},

	"agent": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"agent_identity": {
		Type:     schema.TypeString,
		Optional: true,
	},

	"bastion_host": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"bastion_user": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"bastion_password": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"bastion_private_key": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"bastion_certificate": {
		Type:     schema.TypeString,
		Optional: true,
	},

	"timeout": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"script_path": {
		Type:     schema.TypeString,
		Optional: true,
	},

	// Computed
	"json": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func linuxDataSourceSSHConnection() *schema.Resource {
	return &schema.Resource{
		Read:   linuxDataSourceSSHConnectionRead,
		Schema: linuxDataSourceSSHConnectionSchema,
	}
}

func linuxDataSourceSSHConnectionRead(d *schema.ResourceData, meta interface{}) error {
	connection := sshconnection.SSHConnection{
		Host: d.Get("host").(string),
		User: d.Get("user").(string),
		Port: d.Get("port").(int),

		Password:    d.Get("password").(string),
		PrivateKey:  d.Get("private_key").(string),
		Certificate: d.Get("certificate").(string),
		HostKey:     d.Get("host_key").(string),

		Agent:         d.Get("agent").(string),
		AgentIdentity: d.Get("agent_identity").(string),

		BastionHost:        d.Get("bastion_host").(string),
		BastionUser:        d.Get("bastion_user").(string),
		BastionPassword:    d.Get("bastion_password").(string),
		BastionPrivateKey:  d.Get("bastion_private_key").(string),
		BastionCertificate: d.Get("bastion_certificate").(string),

		Timeout:    d.Get("timeout").(string),
		ScriptPath: d.Get("script_path").(string),
	}

	bytes, err := json.Marshal(connection)
	if err != nil {
		return fmt.Errorf("encoding to JSON: %s", err)
	}
	setOrPanic(d, "json", string(bytes))

	hash := md5.New()
	hash.Write(bytes)
	d.SetId(connection.ID())
	manager.Manager().AddConnection(connection)

	return nil
}
