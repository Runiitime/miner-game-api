package worker

import (
	"context"
	"miner-game-api/helpers"
	"miner-game-api/logic/miner"
	"time"
)

type Middle struct {
	worker miner.WorkerInfo
}

func NewMiddleWorker() *Middle {
	return &Middle{
		worker: miner.WorkerInfo{
			Id:        helpers.GenerateId(),
			Salary:    50,
			Energy:    45,
			Coal:      3,
			BreakTime: 2,
		},
	}
}

func (m *Middle) Run(ctx context.Context, ch chan int64) {
	mCtx, cancel := context.WithCancel(ctx)
	defer close(ch)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		case <-mCtx.Done():
			return
		case <-time.After(time.Duration(m.worker.BreakTime) * time.Millisecond):
			if m.worker.Energy == 0 {
				return
			}

			ch <- m.worker.Coal
			m.decreaseEnergy()
		}
	}
}

func (m *Middle) decreaseEnergy() {
	m.worker.Energy--
}

func (m *Middle) Info() miner.WorkerInfo {
	return m.worker
}
