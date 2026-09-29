package facility

import (
	"context"
	"fmt"
	"miner-game-api/helpers"
	"miner-game-api/logic/balance"
	"miner-game-api/logic/miner"
	"slices"
	"sync"
	"time"
)

type Facility struct {
	equipments []string
	balance    *balance.Balance
	ctx        context.Context
	stop       context.CancelFunc
	wg         *sync.WaitGroup
	miners     []*miner.Miner
	mtx        *sync.Mutex
}

func NewFacility(ctx context.Context) *Facility {
	fCtx, cancel := context.WithCancel(ctx)

	return &Facility{
		equipments: make([]string, 0),
		balance:    balance.NewBalance(),
		ctx:        fCtx,
		stop:       cancel,
		wg:         new(sync.WaitGroup),
		miners:     make([]*miner.Miner, 0),
		mtx:        new(sync.Mutex),
	}
}

/*
Start Запуск предприятия. Пассивное увеличение баланса предприятия каждую секунду
*/
func (f *Facility) Start() {
	fmt.Println("Предприятие начало работу")
	for {
		select {
		case <-f.ctx.Done():
			return
		case <-time.After(1 * time.Second):
			f.balance.Increase(1)
		}
	}
}

/*
Stop Остановка/завершение работы предприятия
*/
func (f *Facility) Stop() {
	f.stop()
}

// Получение информации предприятия

/*
GetBalance Получить информацию о балансе предприятия
*/
func (f *Facility) GetBalance() int64 {
	return f.balance.Get()
}

/*
GetEquipments Получить информацию о купленных инструментах
*/
func (f *Facility) GetEquipments() []string {
	return f.equipments
}

/*
GetWorkingMiners Получить список всех работающих щахтёров
*/
func (f *Facility) GetWorkingMiners() []string {
	var items = make([]string, 0)
	for _, item := range f.miners {
		fmt.Println("================================>ITEM", item)
		// TODO OA RACE
		f.mtx.Lock()
		if item.Info().Energy > 0 {
			items = append(items, item.Info().Id)
		}
		f.mtx.Unlock()
	}
	fmt.Println("f.miners", f.miners)
	return items
}

/*
GetMinersInfoByLevel Получить информацию о том, сколько всего шахтёров каждого класса было нанято
*/
func (f *Facility) GetMinersInfoByLevel() {}

// ШАХТЁРЫ

/*
HireMiner Нанять шахтёра
*/
func (f *Facility) HireMiner(level string) {
	// Наняла шахтёра
	mnr := miner.Hire(level)
	f.mtx.Lock()
	f.miners = append(f.miners, mnr)
	f.balance.Decrease(mnr.Info().Salary)
	f.mtx.Unlock()

	go f.doMine(mnr)
}

/*
doMine Отправить шахтёра работать
*/
func (f *Facility) doMine(mnr *miner.Miner) {
	wg := new(sync.WaitGroup)
	mCtx, _ := context.WithCancel(f.ctx)
	wg.Add(1)
	go mnr.Run(mCtx, wg)

	wg.Wait()
	f.balance.Increase(mnr.Coal)
}

// ИНСТРУМЕНТЫ

/*
BuyEquipment Покупка инструмаента по ID
*/
func (f *Facility) BuyEquipment(id string, price int64) error {
	if slices.Contains(f.equipments, id) {
		return helpers.ErrEquipmentWasAlreadyBought
	}

	f.mtx.Lock()
	f.balance.Decrease(price)
	f.equipments = append(f.equipments, id)
	f.mtx.Unlock()

	return nil
}
