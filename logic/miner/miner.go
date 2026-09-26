package miner

import (
	"context"
	"fmt"
	"sync"
)

type WorkerInfo struct {
	Id        string
	Salary    int
	Energy    int
	Coal      int64
	BreakTime int
	LevelUp   bool
}

type MinerMethods interface {
	Run(ctx context.Context, ch chan int64)
	Info() WorkerInfo
}

type Miner struct {
	Miner MinerMethods
	Coal  int64
}

func NewMiner(miner MinerMethods) *Miner {
	return &Miner{
		Miner: miner,
	}
}

func (m *Miner) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Miner run is done")

	ch := make(chan int64)

	wg.Add(1)
	go func() {
		defer wg.Done()
		go m.Miner.Run(ctx, ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Jun sum done")
		for v := range ch {
			m.Coal += v
		}
	}()
}

func (m *Miner) Info() WorkerInfo {
	return m.Miner.Info()
}
