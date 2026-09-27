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
	wg         *sync.WaitGroup
	miners     []*miner.Miner
	mtx        *sync.Mutex
}

func NewFacility(ctx context.Context, wg *sync.WaitGroup) *Facility {
	return &Facility{
		equipments: make([]string, 0),
		balance:    balance.NewBalance(),
		ctx:        ctx,
		wg:         wg,
		miners:     make([]*miner.Miner, 0),
		mtx:        new(sync.Mutex),
	}
}

/*
Start Запуск предприятия. Пассивное увеличение баланса предприятия каждую секунду
*/
func (f *Facility) Start() {
	defer f.wg.Done()
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
func (f *Facility) Stop() {}

// Получение информации предприятия

/*
GetBalance Получить информацию о балансе предприятия
*/
func (f *Facility) GetBalance() int64 {
	return f.balance.Get()
}

/*
GetEquipments Полученть информацию о купленных инструментах
*/
func (f *Facility) GetEquipments() []string {
	return f.equipments
}

/*
GetWorkingMiners Получить список всех работающих щахтёров
*/
func (f *Facility) GetWorkingMiners() {

}

/*
GetMinersInfoByLevel Получить информацию о том, сколько всего шахтёров каждого класса было нанято
*/
func (f *Facility) GetMinersInfoByLevel() {}

// ШАХТЁРЫ

/*
HireMiner Нанять шахтёра
*/
func (f *Facility) HireMiner(miner *miner.Miner) {
	defer f.wg.Done()
	// Наняла шахтёра
	f.mtx.Lock()
	f.miners = append(f.miners, miner)
	f.mtx.Unlock()
	// Нанятый шахтёр начал работать
	wg := new(sync.WaitGroup)
	mCtx, _ := context.WithCancel(f.ctx)
	wg.Add(1)
	go miner.Run(mCtx, wg)

	wg.Wait()

	// После окончания работы, увеличиваю баланс предприятия на суммарное кол-во добытого угля шахтёром
	f.balance.Increase(miner.Coal)
}

/*
RunMiner Отправить шахтёра работать
*/
func (f *Facility) RunMiner(id string) {

}

// ИНСТРУМЕНТЫ

/*
BuyEquipment Покупка инструмаента по ID
*/
func (f *Facility) BuyEquipment(id string) error {
	if slices.Contains(f.equipments, id) {
		return helpers.ErrEquipmentWasAlreadyBought
	}

	f.equipments = append(f.equipments, id)
	return nil
}
