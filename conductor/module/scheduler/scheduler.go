package scheduler

import "sync"

type Scheduler struct {
	RandomAlgorithm  *RandomAlgorithm
	MaxFreeAlgorithm *MaxFreeAlgorithm
}

var scheduler *Scheduler
var once sync.Once

func GetScheduler() *Scheduler {
	once.Do(func() {
		scheduler = &Scheduler{
			RandomAlgorithm:  GetRandomAlgorithm(),
			MaxFreeAlgorithm: GetMaxFreeAlgorithm(),
		}
	})
	return scheduler
}
