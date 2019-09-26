package linux

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/communicator/remote"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sam-myers/terraform-provider-linux/linux/fileutil"
	"github.com/sam-myers/terraform-provider-linux/linux/manager"
	"io"
	"strings"
)

func linuxFile() *schema.Resource {
	return &schema.Resource{
		Read:   linuxFileRead,
		Create: linuxFileCreate,
		Delete: linuxFileDelete,
		Schema: map[string]*schema.Schema{
			"connection_id": {
				Type:      schema.TypeString,
				ForceNew:  true,
				Required:  true,
				Sensitive: true,
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The content to copy to the destination",
			},
			"destination": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The destination path. It must be specified as an absolute path",
			},

			// Computed
			"hash_md5": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "MD5 hash",
			},
		},
	}
}

func linuxFileCreate(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)
	content := d.Get("content").(string)
	contentReader := strings.NewReader(content)

	comm, err := manager.GetManager().GetCommunicator(d.Get("connection_id").(string))
	if err != nil {
		return err
	}

	err = comm.Upload(destination, contentReader)
	if err != nil {
		return fmt.Errorf("uploading file: %s", err)
	}

	d.SetId(uuid.New().String())

	hash := md5.New()
	_, _ = io.WriteString(hash, content)
	setOrPanic(d, "hash_md5", fmt.Sprintf("%x", hash.Sum(nil)))

	return nil
}

func linuxFileDelete(d *schema.ResourceData, meta interface{}) error {
	comm, err := manager.GetManager().GetCommunicator(d.Get("connection_id").(string))

	destination := d.Get("destination").(string)
	rmCmd := fmt.Sprintf(`rm -f "%s"`, destination)

	command := remote.Cmd{
		Command: rmCmd,
	}
	err = comm.Start(&command)
	if err != nil {
		return fmt.Errorf("deleting file: %s", err)
	}

	d.SetId("")
	return nil
}

func linuxFileRead(d *schema.ResourceData, meta interface{}) error {
	comm, err := manager.GetManager().GetCommunicator(d.Get("connection_id").(string))
	if err != nil {
		// Don't change state if read fails
		return nil
	}

	destination := d.Get("destination").(string)
	exists, err := fileutil.Exists(comm, destination)
	if err != nil {
		return err
	}

	// File is deleted, so destroy the resource
	if !exists {
		d.SetId("")
	}

	oldHash := d.Get("hash_md5").(string)
	newHash, err := fileutil.HashMD5(comm, destination)
	if err != nil {
		return fmt.Errorf("getting md5: %s", err)
	}

	changedOnRemoteMessage := "changed on remote"
	if oldHash != newHash && d.Get("content").(string) == changedOnRemoteMessage {
		setOrPanic(d, "content", changedOnRemoteMessage+" :) nice try")

	} else if oldHash != newHash {
		setOrPanic(d, "content", changedOnRemoteMessage)
	}

	return nil
}
