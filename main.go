package main

import (
	"context"
	"fmt"
	"miner-game-api/logic/facility"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	fac := facility.NewFacility(ctx, &wg)

	go func() {
		time.Sleep(30 * time.Second)
		cancel()
	}()
	wg.Add(1)
	fac.Start()

	wg.Wait()

	fmt.Println("Предприятие завершило работу")
	fmt.Println("Баланс предприятия составляет:", fac.GetBalance())

}
