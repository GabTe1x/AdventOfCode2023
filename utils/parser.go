package utils

import (
	"os"
	"strings"
)

func ReadFromFile(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data_string := string(data[:])
	data_split := strings.Split(data_string, "\n")
	return data_split
}
