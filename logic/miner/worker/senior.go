package worker

import (
	"context"
	"miner-game-api/models"
	"sync"
	"time"
)

type Senior struct {
	worker models.WorkerInfo
}

func NewSeniorWorker() *Senior {
	return &Senior{
		worker: models.WorkerInfo{
			Id:        "senior",
			Salary:    450,
			Energy:    60,
			Coal:      10,
			BreakTime: 1,
		},
	}
}

func (s *Senior) Run(ctx context.Context, ch chan int64) {
	sCtx, cancel := context.WithCancel(ctx)
	mtx := sync.Mutex{}
	defer close(ch)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case <-sCtx.Done():
			return
		case <-time.After(time.Duration(s.worker.BreakTime) * time.Second):
			if s.worker.Energy == 0 {
				cancel()
				return
			}

			mtx.Lock()
			ch <- s.worker.Coal
			mtx.Unlock()
			s.decreaseEnergy()
			s.levelUp(&mtx) // TODO возможность не по ссылке
		}
	}
}

func (s *Senior) levelUp(mtx *sync.Mutex) {
	mtx.Lock()
	s.worker.Coal += 1
	mtx.Unlock()
}

func (s *Senior) decreaseEnergy() {
	s.worker.Energy--
}

func (s *Senior) Info() models.WorkerInfo {
	return s.worker
}
