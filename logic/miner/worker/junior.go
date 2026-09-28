package worker

import (
	"context"
	"miner-game-api/models"
	"sync"
	"time"
)

type Junior struct {
	mtx    sync.Mutex
	worker models.WorkerInfo
}

func NewJuniorWorker() *Junior {
	return &Junior{
		mtx: sync.Mutex{},
		worker: models.WorkerInfo{
			Id:        "junior",
			Salary:    5,
			Energy:    30,
			Coal:      1,
			BreakTime: 3,
		},
	}
}

func (j *Junior) Run(ctx context.Context, ch chan int64) {
	jCtx, cancel := context.WithCancel(ctx)
	defer close(ch)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		case <-jCtx.Done():
			return
		case <-time.After(time.Duration(j.Info().BreakTime) * time.Second):
			j.mtx.Lock()
			energy := j.worker.Energy
			j.mtx.Unlock()
			if energy == 0 {
				return
			}

			ch <- j.worker.Coal
			j.decreaseEnergy()
		}
	}
}

func (j *Junior) decreaseEnergy() {
	j.mtx.Lock()
	defer j.mtx.Unlock()
	j.worker.Energy--
}

func (j *Junior) Info() models.WorkerInfo {
	return j.worker
}
