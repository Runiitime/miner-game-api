package worker

import (
	"miner-game-api/helpers"
	"miner-game-api/logic/miner"
)

type Junior struct {
	worker miner.WorkerInfo
}

func NewJuniorWorker() *Junior {
	return &Junior{
		worker: miner.WorkerInfo{
			Id:        helpers.GenerateId(),
			Salary:    5,
			Energy:    30,
			Coal:      1,
			BreakTime: 3,
			LevelUp:   false,
		},
	}
}

func (s *Junior) Run() {

}

func (s *Junior) Info() miner.WorkerInfo {
	return s.worker
}
