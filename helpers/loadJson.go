package helpers

import (
	"os"
)

func LoadJson(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
