package main

import (
	"context"
	"miner-game-api/handlers"
	"miner-game-api/logic/equipments"
	"miner-game-api/logic/facility"
	"miner-game-api/logic/miner"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	mnrHR := miner.NewMinersHR()
	eqStore := equipments.NewEquipmentsStore()

	facil := facility.NewFacility(ctx)
	hdls := handlers.NewHTTPHandlers(facil, eqStore, mnrHR)
	server := handlers.NewHTTPServer(hdls)

	go facil.Start()
	server.Start("9091")
}

//
//func main() {
//	var ctx, _ = context.WithCancel(context.Background())
//	fwg := new(sync.WaitGroup)
//	facil := facility.NewFacility(ctx, fwg)
//
//	fwg.Add(1)
//	go func() {
//		defer fwg.Done()
//		for i := 1; i <= 3; i++ {
//			fwg.Add(1)
//			jun := worker.NewJuniorWorker()
//			mnr := miner.NewMiner(jun)
//
//			go facil.HireMiner(mnr)
//		}
//	}()
//
//	fwg.Add(1)
//	go func() {
//		defer fwg.Done()
//		for i := 1; i <= 2; i++ {
//			fwg.Add(1)
//			mid := worker.NewMiddleWorker()
//			mnr := miner.NewMiner(mid)
//
//			go facil.HireMiner(mnr)
//		}
//	}()
//
//	fwg.Add(1)
//	go func() {
//		defer fwg.Done()
//		for i := 1; i <= 1; i++ {
//			fwg.Add(1)
//			sen := worker.NewSeniorWorker()
//			mnr := miner.NewMiner(sen)
//
//			go facil.HireMiner(mnr)
//		}
//	}()
//
//	fwg.Wait()
//
//	fmt.Println("done")
//	fmt.Println("Facility balance:", facil.GetBalance())
//}
