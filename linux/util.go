package linux

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func SetOrPanic(d *schema.ResourceData, key string, value interface{}) {
	err := d.Set(key, value)
	if err != nil {
		panic(fmt.Sprintf("invariant broken, bug in provider. trying to set key `%s` with value `%+v` failed: %s", key, value, err))
	}
}
