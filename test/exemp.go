package test

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const energy = 7
const miners = 3

var mtx sync.Mutex
var coal = 0

/*TODO Удалить позже. Тестовый код*/
func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	mineWg := new(sync.WaitGroup)
	sumWg := new(sync.WaitGroup)
	incWg := new(sync.WaitGroup)
	ch := make(chan int)

	mineWg.Add(1)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= miners; i++ {
			fmt.Println("Miner №", i, "start working")
			wg.Add(1)
			go mine(ch, wg, i)
		}
	}(ch, mineWg)

	incWg.Add(1)
	go inc(incWg, ctx)

	sumWg.Add(1)
	go sum(ch, sumWg)

	mineWg.Wait()
	close(ch)
	sumWg.Wait()
	ctxCancel()
	incWg.Wait()
	fmt.Println(coal)
}

func inc(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
			mtx.Lock()
			fmt.Println("INC added")
			coal += 1
			mtx.Unlock()
		}
	}
}

func sum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		mtx.Lock()
		coal += n
		mtx.Unlock()
	}
}

func mine(ch chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	defer fmt.Println("Miner №", n, "finished work")
	for i := 1; i <= n; i++ {
		time.Sleep(2500 * time.Millisecond)
		ch <- n
	}
}
