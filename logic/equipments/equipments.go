package equipments

import (
	"encoding/json"
	"log"
	"maps"
	"miner-game-api/helpers"
	"miner-game-api/models"
	"os"
	"slices"
	"strconv"
)

type Equipments struct {
	Equipments models.EquipmentInfoMap
	Titles     []string
}

func NewEquipmentsStore() *Equipments {
	items := loadEquipments()
	titles := slices.Collect(maps.Keys(items))

	return &Equipments{
		Equipments: items,
		Titles:     titles,
	}
}

func loadEquipments() models.EquipmentInfoMap {
	data, err := os.ReadFile("storage/equipments.json")

	items := make(models.EquipmentInfoMap)

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

func (e *Equipments) BuyCheck(id string, balance int64) error {
	eq := e.Equipments[id]
	if eq.Price > balance {
		return helpers.ErrNotEnoughBalance
	}

	if eq.IsSold {
		return helpers.ErrEquipmentWasSold
	}

	eq.IsSold = true
	e.Equipments[id] = eq

	return nil
}

func (e *Equipments) GetEquipmentsByStatus(status string) (models.EquipmentInfoMap, error) {
	value, err := strconv.ParseBool(status)
	if err != nil {
		return nil, helpers.ErrWrongQueryParamValue
	}

	var items = make(models.EquipmentInfoMap)

	for _, item := range e.Equipments {
		if item.IsSold == value {
			items[item.Id] = item
		}
	}

	return items, nil
}
