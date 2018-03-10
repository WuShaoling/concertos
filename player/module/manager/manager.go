package manager

import "sync"

func (m *Manager) Start() {
}

var manager *Manager
var once sync.Once

type Manager struct {
	PlayerManager *PlayerManager
}

func GetManage() *Manager {
	once.Do(func() {
		manager = &Manager{
			PlayerManager: GetPlayerManage(),
		}
	})
	return manager
}
