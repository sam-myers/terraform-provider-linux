package fileutil

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform/communicator"
	"github.com/hashicorp/terraform/communicator/remote"
	"io/ioutil"
)

func HashMD5(comm communicator.Communicator, path string) (string, error) {
	md5Installed, err := Installed(comm, "md5sum")
	if err != nil {
		return "", fmt.Errorf("checking if md5sum installed: %s", err)
	} else if !md5Installed {
		return "", fmt.Errorf("md5sum not installed")
	}

	var stdOutBuffer bytes.Buffer
	command := remote.Cmd{
		Command: fmt.Sprintf(`md5sum "%s"`, path),
		Stdout:  &stdOutBuffer,
	}

	err = comm.Start(&command)
	if err != nil {
		return "", fmt.Errorf("starting md5 hash command: %s", err)
	}

	err = command.Wait()
	if err != nil {
		return "", fmt.Errorf("running md5 hash command: %s", err)
	}

	stdOutBytes, err := ioutil.ReadAll(&stdOutBuffer)
	return string(stdOutBytes[:32]), nil
}
