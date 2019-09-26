package sshconnection

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/communicator"
	"github.com/hashicorp/terraform/helper/schema"
	"sync"
)

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

	comm     communicator.Communicator
	commErr  error
	commOnce sync.Once
}

func (s *SSHConnection) ID() string {
	bytes, _ := json.Marshal(s)
	hash := md5.New()
	hash.Write(bytes)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *SSHConnection) Communicator() (communicator.Communicator, error) {
	s.commOnce.Do(func() {
		id := uuid.New().String()
		d := schema.ResourceData{}
		d.SetConnInfo(s.ToMap())
		d.SetId(id)
		s.comm, s.commErr = communicator.New(d.State())
		d.SetId("")

		if s.commErr != nil {
			return
		}

		s.commErr = backoff.Retry(func() error {
			return s.comm.Connect(nil)
		}, backoff.NewExponentialBackOff())
	})
	return s.comm, s.commErr
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
