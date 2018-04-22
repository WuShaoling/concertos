package executor

import "sync"

type Executor struct {
	DockerExecutor *DockerExecutor
}

var executor *Executor
var once sync.Once

func GetExecutor() *Executor {
	once.Do(func() {
		executor = &Executor{
			DockerExecutor: GetDockerExecutor(),
		}
	})
	return executor
}
