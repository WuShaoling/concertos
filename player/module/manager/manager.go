package manager

import "sync"

func (m *Manager) Start() {
	go m.ContainerManager.KeepContainerAlive()
	m.PlayerManager.KeepAlive()
}

var manager *Manager
var once sync.Once

type Manager struct {
	ContainerManager *ContainerManager
	PlayerManager    *PlayerManager
}

func GetManage() *Manager {
	once.Do(func() {
		manager = &Manager{
			ContainerManager: GetContainerManager(),
			PlayerManager:    GetPlayerManager(),
		}
	})
	return manager
}
