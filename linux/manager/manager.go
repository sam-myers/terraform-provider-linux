package manager

import (
	"fmt"
	"github.com/hashicorp/terraform/communicator"
	"github.com/sam-myers/terraform-provider-linux/linux/sshconnection"
	"sync"
)

var manCreated sync.Once
var man *Manager

// Get the global manager instance
func GetManager() *Manager {
	manCreated.Do(func() {
		man = &Manager{
			lock:        sync.Mutex{},
			connections: make(map[string]sshconnection.SSHConnection, 0),
		}
	})
	return man
}

type Manager struct {
	lock        sync.Mutex
	connections map[string]sshconnection.SSHConnection
}

func (m *Manager) AddConnection(conn sshconnection.SSHConnection) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.connections[conn.ID()] = conn
}

func (m *Manager) GetConnection(id string) (conn sshconnection.SSHConnection, found bool) {
	conn, found = m.connections[id]
	return
}

func (m *Manager) GetCommunicator(id string) (communicator.Communicator, error) {
	conn, found := m.connections[id]
	if !found {
		return nil, fmt.Errorf("no communicator found with id %s", id)
	}
	return conn.Communicator()
}
