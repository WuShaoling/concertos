package executor

type Docker struct {
	Base Container
	Type int
}

func (d Docker) Install() bool {
	return false;
}

func (d Docker) Start() bool {
	return false;
}

func (d Docker) Stop() bool {
	return false;
}

func (d Docker) Pause() bool {
	return false;
}

func (d Docker) Restart() bool {
	return false;
}
