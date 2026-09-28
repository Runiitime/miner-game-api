package main

import (
	"context"
	"miner-game-api/handlers"
	"miner-game-api/logic/facility"
	"miner-game-api/logic/info"
	"sync"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	inf := info.NewInfo()
	facil := facility.NewFacility(ctx, wg)
	hdls := handlers.NewHTTPHandlers(facil, inf)
	server := handlers.NewHTTPServer(hdls)

	server.Start("9091")
	//r := gin.Default()
	//
	//r.GET("/aboba", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "ABOBA",
	//	})
	//})
	//
	//if err := r.Run(":8080"); err != nil {
	//	panic(err)
	//}
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
