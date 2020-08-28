package internalutils

import (
	"bufio"
	"os"
	"strings"
)

func SetEnvir(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, each := range lines {
		keyVal := strings.Split(each, ":")
		os.Setenv(keyVal[0], keyVal[1])
	}
	return
}
