package main

import (
	"context"
	"fmt"
	"miner-game-api/logic/miner"
	"miner-game-api/logic/miner/worker"
	"sync"
)

var energy = 5

const coal int64 = 1

var sum int64 = 0
var ctx, cancel = context.WithCancel(context.Background())

func main() {
	wg := new(sync.WaitGroup)

	var workers []*miner.Miner

	for i := 0; i < 4; i++ {
		jun := worker.NewJuniorWorker()
		mnr := miner.NewMiner(jun)

		workers = append(workers, mnr)
	}

	for _, m := range workers {
		wg.Add(1)
		go m.Run(ctx, wg)
	}

	wg.Wait()

	for _, m := range workers {
		sum += m.Coal
	}

	fmt.Println("All coal:", sum)
}
