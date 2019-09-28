package manager

import (
	"fmt"
	"github.com/hashicorp/terraform/communicator"
	"github.com/sam-myers/terraform-provider-linux/linux/sshconnection"
	"sync"
)

var manCreated sync.Once
var man *manager

// getManager gets the global manager instance
func getManager() *manager {
	manCreated.Do(func() {
		man = &manager{
			lock:        sync.Mutex{},
			connections: make(map[string]sshconnection.SSHConnection, 0),
		}
	})
	return man
}

// manager allows resources to access useful information about related resources
type manager struct {
	lock        sync.Mutex
	connections map[string]sshconnection.SSHConnection
}

// AddConnection tells the manager about an SSH connection
func AddConnection(conn sshconnection.SSHConnection) {
	m := getManager()
	m.lock.Lock()
	defer m.lock.Unlock()

	m.connections[conn.ID()] = conn
}

// GetConnection allows resources to get an SSH connection by its ID
func GetConnection(id string) (conn sshconnection.SSHConnection, found bool) {
	m := getManager()
	m.lock.Lock()
	defer m.lock.Unlock()

	conn, found = m.connections[id]
	return
}

// GetCommunicator gets a communicator, used to directly interface with a Linux machine,
// from a connection ID
func GetCommunicator(connectionID string) (communicator.Communicator, error) {
	m := getManager()
	m.lock.Lock()
	defer m.lock.Unlock()

	conn, found := m.connections[connectionID]
	if !found {
		return nil, fmt.Errorf("no communicator found with connectionID %s", connectionID)
	}
	return conn.Communicator()
}
