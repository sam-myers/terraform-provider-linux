package fileutil

import (
	"fmt"
	"github.com/hashicorp/terraform/communicator"
	"github.com/hashicorp/terraform/communicator/remote"
)

func Exists(comm communicator.Communicator, path string) (bool, error) {
	command := remote.Cmd{
		Command: fmt.Sprintf(`stat "%s"`, path),
	}

	err := comm.Start(&command)
	if err != nil {
		return false, fmt.Errorf("starting file exist command: %s", err)
	}

	err = command.Wait()
	if err != nil {
		return false, nil
	}

	return true, nil
}
