package executor

type DockerExecutor struct {
}

func (d *DockerExecutor) Install() error {

	return nil;
}

func (d *DockerExecutor) Start() error {

	return nil;
}

func (d *DockerExecutor) Stop() error {

	return nil;
}

func (d *DockerExecutor) Remove() error {

	return nil;
}

func GetDockerExecutor() *DockerExecutor {
	return &DockerExecutor{}
}
