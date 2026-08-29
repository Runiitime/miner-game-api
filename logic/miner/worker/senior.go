package worker

import (
	"context"
	"miner-game-api/helpers"
)

type Senior struct {
	id        string
	salary    int
	energy    int
	coal      int64
	breakTime int
	levelUp   bool
}

func NewSeniorWorker() *Senior {
	return &Senior{
		id:        helpers.GenerateId(),
		salary:    450,
		energy:    60,
		coal:      10,
		breakTime: 1,
		levelUp:   true,
	}
}

func (s *Senior) Run(ctx context.Context, ch chan int64) {

}

func (s *Senior) LevelUp() {
	s.coal += 3
}

func (s *Senior) Info() {

}
