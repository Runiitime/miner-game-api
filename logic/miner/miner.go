package miner

import (
	"context"
	"miner-game-api/logic/miner/worker"
	"miner-game-api/models"
	"sync"
)

type MinerMethods interface {
	Run(ctx context.Context, ch chan int64)
	Info() models.WorkerInfo
}

type Miner struct {
	Miner MinerMethods
	Coal  int64
}

func Hire(level string) *Miner {
	if level == "junior" {
		return newMiner(worker.NewJuniorWorker())
	}

	if level == "middle" {
		return newMiner(worker.NewMiddleWorker())
	}

	return newMiner(worker.NewSeniorWorker())
}

func newMiner(miner MinerMethods) *Miner {
	return &Miner{
		Miner: miner,
	}
}

func (m *Miner) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ch := make(chan int64)

	wg.Add(1)
	go func() {
		defer wg.Done()
		go m.Miner.Run(ctx, ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			m.Coal += v
		}
	}()
}

func (m *Miner) Info() models.WorkerInfo {
	return m.Miner.Info()
}
