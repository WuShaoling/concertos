package executor

import "sync"

type Executor struct {
	DockerExecutor *DockerExecutor
	RegExecutor    *RegExecutor
}

var executor *Executor
var once sync.Once

func GetExecutor() *Executor {
	once.Do(func() {
		executor = &Executor{
			DockerExecutor: GetDockerExecutor(),
			RegExecutor:    GetRegExecutor(),
		}
	})
	return executor
}
