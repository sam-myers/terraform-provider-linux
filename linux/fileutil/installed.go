package fileutil

import (
	"fmt"
	"github.com/hashicorp/terraform/communicator"
	"github.com/hashicorp/terraform/communicator/remote"
)

func Installed(comm communicator.Communicator, program string) (bool, error) {
	command := remote.Cmd{
		Command: fmt.Sprintf(`which "%s"`, program),
	}

	err := comm.Start(&command)
	if err != nil {
		return false, fmt.Errorf("starting installed command: %s", err)
	}

	err = command.Wait()
	if err != nil {
		return false, nil
	}

	return true, nil
}
