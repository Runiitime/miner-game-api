package info

import (
	"encoding/json"
	"log"
	"miner-game-api/logic/miner"
	"miner-game-api/models"
	"os"
)

type EquipmentsInfo = map[string]models.Equipment
type MinersInfo = map[string]miner.WorkerInfo

type Info struct {
	EquipmentsInfo *EquipmentsInfo
	MinersInfo     *MinersInfo
}

func NewInfo(eq *EquipmentsInfo, m *MinersInfo) *Info {
	return &Info{EquipmentsInfo: loadEquipmentsInfo(), MinersInfo: loadMinersInfo()}
}

func loadMinersInfo() *MinersInfo {
	data, err := os.ReadFile("storage/miners.json")

	items := new(make(MinersInfo))

	if err != nil {
		log.Println(err)
		return items
	}

	if err := json.Unmarshal(data, items); err != nil {
		log.Println(err)
		return items
	}

	return items
}

func loadEquipmentsInfo() *EquipmentsInfo {
	data, err := os.ReadFile("storage/equipments.json")

	items := new(make(EquipmentsInfo))

	if err != nil {
		log.Println(err)
		return items
	}

	if err := json.Unmarshal(data, items); err != nil {
		log.Println(err)
		return items
	}

	return items
}
