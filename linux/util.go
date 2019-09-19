package linux

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/communicator"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func SetOrPanic(d *schema.ResourceData, key string, value interface{}) {
	err := d.Set(key, value)
	if err != nil {
		panic(fmt.Sprintf("invariant broken, bug in provider. trying to set key `%s` with value `%+v` failed: %s", key, value, err))
	}
}

func GetCommunicator(d *schema.ResourceData) (communicator.Communicator, error) {
	var state *terraform.InstanceState

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

	err = comm.Connect(nil)
	if err != nil {
		return comm, fmt.Errorf("connecting: %s", err)
	}

	return comm, nil
}
