package models

type EquipmentId struct {
	Id string `json:"id"`
}

type Equipment struct {
	Id     string
	Label  string
	Price  int
	IsSold bool
}
