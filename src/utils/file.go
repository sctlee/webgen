package utils

import (
	"fmt"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Open(path string) ([]byte, int) {
	// Usage:
	// data, count := openFile("info.yml")
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
	file, err := os.Open(path) // For read access.
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 1000)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	return data, count
}
