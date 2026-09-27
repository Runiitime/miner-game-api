package main

import (
	"context"
	"fmt"
	"miner-game-api/logic/facility"
	"miner-game-api/logic/miner"
	"miner-game-api/logic/miner/worker"
	"sync"
)

func main() {
	var ctx, _ = context.WithCancel(context.Background())
	fwg := new(sync.WaitGroup)
	facil := facility.NewFacility(ctx, fwg)

	fwg.Add(1)
	go func() {
		defer fwg.Done()
		for i := 1; i <= 3; i++ {
			fwg.Add(1)
			jun := worker.NewJuniorWorker()
			mnr := miner.NewMiner(jun)

			go facil.HireMiner(mnr)
		}
	}()

	fwg.Add(1)
	go func() {
		defer fwg.Done()
		for i := 1; i <= 2; i++ {
			fwg.Add(1)
			mid := worker.NewMiddleWorker()
			mnr := miner.NewMiner(mid)

			go facil.HireMiner(mnr)
		}
	}()

	fwg.Add(1)
	go func() {
		defer fwg.Done()
		for i := 1; i <= 1; i++ {
			fwg.Add(1)
			sen := worker.NewSeniorWorker()
			mnr := miner.NewMiner(sen)

			go facil.HireMiner(mnr)
		}
	}()

	fwg.Wait()

	fmt.Println("done")
	fmt.Println("Facility balance:", facil.GetBalance())
}
