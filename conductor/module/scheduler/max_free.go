package scheduler

type MaxFreeAlgorithm struct {
}

func GetMaxFreeAlgorithm() *MaxFreeAlgorithm {
	return &MaxFreeAlgorithm{
	}
}

func (mfa *MaxFreeAlgorithm) GetPlayerId() string {
	return ""
}
