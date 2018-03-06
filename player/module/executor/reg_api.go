package executor

type Reg struct {
	Base Container
	Type int
}

func (d Reg) Install() bool {
	return false;
}

func (d Reg) Start() bool {
	return false;
}

func (d Reg) Stop() bool {
	return false;
}

func (d Reg) Pause() bool {
	return false;
}

func (d Reg) Restart() bool {
	return false;
}
