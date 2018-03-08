package scheduler

import "sync"

type Scheduler struct {
	randomAlgorithm  *RandomAlgorithm
	maxFreeAlgorithm *MaxFreeAlgorithm
}

var scheduler *Scheduler
var once sync.Once

func GetScheduler() *Scheduler {
	once.Do(func() {
		scheduler = &Scheduler{
			randomAlgorithm:  GetRandomAlgorithm(),
			maxFreeAlgorithm: GetMaxFreeAlgorithm(),
		}
	})
	return scheduler
}
