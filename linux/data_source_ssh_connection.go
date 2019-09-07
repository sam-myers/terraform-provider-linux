package linux

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func linuxDataSourceSSHConnection() *schema.Resource {
	return &schema.Resource{
		Read: linuxDataSourceSSHConnectionRead,
		Schema: map[string]*schema.Schema{
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
		},
	}
}

type SSHConnection struct {
	Host string `json:"host,omitempty"`
	User string `json:"user,omitempty"`
	Port int    `json:"port,omitempty"`

	Password    string `json:"password,omitempty"`
	PrivateKey  string `json:"private_key,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	HostKey     string `json:"host_key,omitempty"`

	Agent         string `json:"agent,omitempty"`
	AgentIdentity string `json:"agent_identity,omitempty"`

	BastionHost        string `json:"bastion_host,omitempty"`
	BastionUser        string `json:"bastion_user,omitempty"`
	BastionPassword    string `json:"bastion_password,omitempty"`
	BastionPrivateKey  string `json:"bastion_private_key,omitempty"`
	BastionCertificate string `json:"bastion_certificate,omitempty"`

	Timeout    string `json:"timeout,omitempty"`
	ScriptPath string `json:"script_path,omitempty"`
}

func linuxDataSourceSSHConnectionRead(d *schema.ResourceData, meta interface{}) error {
	connection := SSHConnection{
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
	SetOrPanic(d, "json", string(bytes))

	hash := md5.New().Sum(bytes)
	d.SetId(string(hash))

	return nil
}

func (s *SSHConnection) ToMap() map[string]string {
	commInfo := make(map[string]string)

	commInfo["host"] = s.Host
	commInfo["user"] = s.User
	commInfo["port"] = fmt.Sprintf("%d", s.Port)

	commInfo["password"] = s.Password
	commInfo["private_key"] = s.PrivateKey
	commInfo["certificate"] = s.Certificate
	commInfo["host_key"] = s.HostKey

	commInfo["agent"] = s.Agent
	commInfo["agent_identity"] = s.AgentIdentity

	commInfo["bastion_host"] = s.BastionHost
	commInfo["bastion_user"] = s.BastionUser
	commInfo["bastion_password"] = s.BastionPassword
	commInfo["bastion_private_key"] = s.BastionPrivateKey
	commInfo["bastion_certificate"] = s.BastionCertificate

	commInfo["timeout"] = s.Timeout
	commInfo["script_path"] = s.ScriptPath

	return commInfo
}

func SetConnectionInfo(d *schema.ResourceData) error {
	connectionJSON := d.Get("connection_json").(string)
	var sshConnection SSHConnection

	err := json.Unmarshal([]byte(connectionJSON), &sshConnection)
	if err != nil {
		return fmt.Errorf("deconding connection JSON: %s", err)
	}

	d.SetConnInfo(sshConnection.ToMap())
	return nil
}
