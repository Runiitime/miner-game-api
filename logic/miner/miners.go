package miner

import (
	"encoding/json"
	"log"
	"maps"
	"miner-game-api/helpers"
	"miner-game-api/models"
	"os"
	"slices"
)

type Miners struct {
	Miners models.MinerInfoMap
	Levels []string
}

func NewMinersHR() *Miners {
	items := loadMiners()
	levels := slices.Collect(maps.Keys(items))

	return &Miners{
		Miners: items,
		Levels: levels,
	}

}

func loadMiners() models.MinerInfoMap {
	data, err := os.ReadFile("storage/miners.json")

	items := make(models.MinerInfoMap)

	if err != nil {
		log.Println(err)
		return items
	}

	if err := json.Unmarshal(data, &items); err != nil {
		log.Println(err)
		return items
	}

	return items
}

func (m *Miners) HireCheck(level string, balance int64) error {
	mnr := m.Miners[level]

	if mnr.Salary > balance {
		return helpers.ErrNotEnoughBalance
	}

	return nil
}

func (m *Miners) SalaryByLevel(level string) int64 {
	return m.Miners[level].Salary
}
