package executor

import "sync"

type Executor struct {
	dockerExecutor *DockerExecutor
	regExecutor    *RegExecutor
}

var executor *Executor
var once sync.Once

func GetExecutor() *Executor {
	once.Do(func() {
		executor = &Executor{
			dockerExecutor: GetDockerExecutor(),
			regExecutor:    GetRegExecutor(),
		}
	})
	return executor
}
