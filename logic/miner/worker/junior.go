package worker

import (
	"context"
	"fmt"
	"miner-game-api/helpers"
	"miner-game-api/logic/miner"
	"time"
)

type Junior struct {
	worker miner.WorkerInfo
}

func NewJuniorWorker() *Junior {
	return &Junior{
		worker: miner.WorkerInfo{
			Id:        helpers.GenerateId(),
			Salary:    5,
			Energy:    3,
			Coal:      1,
			BreakTime: 2,
			LevelUp:   false,
		},
	}
}

func (j *Junior) Run(ctx context.Context, ch chan int64) {
	jCtx, cancel := context.WithCancel(ctx)
	defer close(ch)
	defer cancel()
	defer fmt.Println("Jun run done")
	for {
		select {
		case <-ctx.Done():
			return
		case <-jCtx.Done():
			return
		case <-time.After(time.Duration(j.Info().BreakTime) * time.Millisecond):
			if j.worker.Energy == 0 {
				cancel()
				return
			}

			ch <- j.worker.Coal
			j.DecreaseEnergy()
		}
	}
}

func (j *Junior) DecreaseEnergy() {
	j.worker.Energy--
}

func (j *Junior) Info() miner.WorkerInfo {
	return j.worker
}
