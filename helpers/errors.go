package helpers

import "errors"

// Equipment not found
var ErrEquipmentNotFound = errors.New("Equipment not found")

// Miner not found
var ErrMinerNotFound = errors.New("Miner not found")

// Equipment was already bought
var ErrEquipmentWasAlreadyBought = errors.New("Equipment was already bought")

// Nothing to buy
var ErrNothingToBuy = errors.New("Nothing to buy")

// No store was found
var ErrNoStore = errors.New("no store was found")

// Failed to write HTTP response
var ErrJsonWrite = errors.New("Failed to write HTTP response")

// Wrong query param value
var ErrWrongQueryParamValue = errors.New("Wrong query param value")

// Empty query param
var ErrEmptyQuery = errors.New("Empty query param")

// Empty query param
var ErrJSONUnmarshal = errors.New("Failed to json unmarshal")
