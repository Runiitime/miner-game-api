package worker

import "miner-game-api/helpers"

type Middle struct {
	id        string
	salary    int
	energy    int
	coal      int
	breakTime int
	levelUp   bool
}

func NewMiddleWorker() *Middle {
	return &Middle{
		id:        helpers.GenerateId(),
		salary:    50,
		energy:    45,
		coal:      3,
		breakTime: 2,
		levelUp:   false,
	}
}

func (s *Middle) Run() {

}
