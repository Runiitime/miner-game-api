package helpers

import (
	"math/rand/v2"
	"strings"
)

var letters = "abcdifghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateId() string {
	var id string
	lettersArr := strings.Split(letters, "")
	for i := 1; i <= 20; i++ {
		var r = rand.IntN(len(lettersArr) - 1)
		id += lettersArr[r]
	}

	return id
}
