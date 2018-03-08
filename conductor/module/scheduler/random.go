package scheduler

type RandomAlgorithm struct {
}

func GetRandomAlgorithm() *RandomAlgorithm {
	return &RandomAlgorithm{}
}

func (ra RandomAlgorithm) GetPlayerId() string {
	return ""
}
