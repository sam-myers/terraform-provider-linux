package linux

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
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
