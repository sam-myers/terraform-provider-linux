package linux

import (
	"bytes"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/communicator"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sam-myers/terraform-provider-linux/linux/manager"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func setOrPanic(d *schema.ResourceData, key string, value interface{}) {
	err := d.Set(key, value)
	if err != nil {
		panic(fmt.Sprintf("invariant broken, bug in provider. trying to set key `%s` with value `%+v` failed: %s", key, value, err))
	}
}

func getCommunicator(d *schema.ResourceData) (communicator.Communicator, error) {
	var state *terraform.InstanceState

	err := SetConnectionInfo(d)
	if err != nil {
		return nil, fmt.Errorf("set connection info: %s", err)
	}

	if d.Id() == "" {
		id := uuid.New().String()
		d.SetId(id)
		state = d.State()
		d.SetId("")
	} else {
		state = d.State()
	}

	if state == nil {
		return nil, fmt.Errorf("no state")
	}

	comm, err := communicator.New(state)
	if err != nil {
		return nil, fmt.Errorf("creating communicator: %s", err)
	}

	err = backoff.Retry(func() error {
		return comm.Connect(nil)
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return comm, fmt.Errorf("connecting: %s", err)
	}

	err = comm.Connect(nil)
	if err != nil {
		return comm, fmt.Errorf("connecting: %s", err)
	}

	return comm, nil
}

func SetConnectionInfo(d *schema.ResourceData) error {
	connectionId := d.Get("connection_id").(string)
	conn, found := manager.GetManager().GetConnection(connectionId)
	if !found {
		return fmt.Errorf("no connection of id: %s", connectionId)
	}

	d.SetConnInfo(conn.ToMap())
	return nil
}

func getFixture(path string) string {
	file, err := os.Open("./fixtures/" + path)
	if err != nil {
		panic(fmt.Sprintf("invalid path %s: %s", path, err))
	}

	bodyBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("couldn't ready template: %s", err))
	}

	t, err := template.New(path).Parse(string(bodyBytes))
	if err != nil {
		panic(fmt.Sprintf("couldn't parse template: %s", err))
	}

	privateKeyPath, err := filepath.Abs("./fixtures/id_rsa")
	if err != nil {
		panic(fmt.Sprintf("couldn't find pub key: %s", err))
	}

	publicKeyPath, err := filepath.Abs("./fixtures/id_rsa.pub")
	if err != nil {
		panic(fmt.Sprintf("couldn't find pub key: %s", err))
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, struct {
		PrivateKeyPath string
		PublicKeyPath  string
	}{
		PrivateKeyPath: privateKeyPath,
		PublicKeyPath:  publicKeyPath,
	})
	if err != nil {
		panic(fmt.Sprintf("couldn't execute template: %s", err))
	}

	return buff.String()
}
