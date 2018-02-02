package executor

type ContainerApi interface {
	Install() bool
	Start() bool
	Stop() bool
	Pause() bool
	Restart() bool
}
