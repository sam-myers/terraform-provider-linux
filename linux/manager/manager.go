package manager

import (
	"fmt"
	"github.com/hashicorp/terraform/communicator"
	"github.com/sam-myers/terraform-provider-linux/linux/sshconnection"
	"sync"
)

var manCreated sync.Once
var man *manager

func Manager() *manager {
	manCreated.Do(func() {
		man = &manager{
			lock:        sync.Mutex{},
			connections: make(map[string]sshconnection.SSHConnection, 0),
		}
	})
	return man
}

type manager struct {
	lock        sync.Mutex
	connections map[string]sshconnection.SSHConnection
}

func (m *manager) AddConnection(conn sshconnection.SSHConnection) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.connections[conn.ID()] = conn
}

func (m *manager) GetConnection(id string) (conn sshconnection.SSHConnection, found bool) {
	conn, found = m.connections[id]
	return
}

func (m *manager) GetCommunicator(id string) (communicator.Communicator, error) {
	conn, found := m.connections[id]
	if !found {
		return nil, fmt.Errorf("no communicator found with id %s", id)
	}
	return conn.Communicator()
}
