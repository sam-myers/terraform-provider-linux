package manager

import (
	"github.com/sam-myers/terraform-provider-linux/linux/sshconnection"
	"sync"
)

var manCreated sync.Once
var man *manager

func GetManager() *manager {
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
