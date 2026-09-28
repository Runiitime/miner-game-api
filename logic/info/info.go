package info

import (
	"encoding/json"
	"log"
	"maps"
	"miner-game-api/models"
	"os"
	"slices"
)

type EquipmentsInfo = map[string]models.Equipment
type MinersInfo = map[string]models.MinerInfo

type Info struct {
	EquipmentsInfo EquipmentsInfo
	MinersInfo     MinersInfo
}

func NewInfo() *Info {
	return &Info{EquipmentsInfo: loadEquipmentsInfo(), MinersInfo: loadMinersInfo()}
}

func loadMinersInfo() MinersInfo {
	data, err := os.ReadFile("storage/miners.json")

	items := make(MinersInfo)

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

func loadEquipmentsInfo() EquipmentsInfo {
	data, err := os.ReadFile("storage/equipments.json")

	items := make(EquipmentsInfo)

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

func (i *Info) MinerLevels() []string {
	levels := i.MinersInfo
	return slices.Collect(maps.Keys(levels))
}

func (i *Info) EquipmentTitles() []string {
	titles := i.EquipmentsInfo
	return slices.Collect(maps.Keys(titles))
}
