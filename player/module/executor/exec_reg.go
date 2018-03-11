package executor

type RegExecutor struct {
}

func (d *RegExecutor) Install() error {
	return nil;
}

func (d *RegExecutor) Start() error {
	return nil;
}

func (d *RegExecutor) Stop() error {
	return nil;
}

func (d *RegExecutor) Remove() error {
	return nil;
}

func GetRegExecutor() *RegExecutor {
	return &RegExecutor{}
}
