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
