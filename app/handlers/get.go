package handlers

import (
	"bufio"
	"os"
	"strings"
)

func Get(arguments []string) ([]byte, error) {
	file, err := os.OpenFile(dbPath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	values := make([]string, 2, 10)
	for scanner.Scan() {
		line = scanner.Text()
		lineValues := strings.Split(line, " ")
		values[0] = lineValues[0]
		if values[0] == arguments[0] {
			values[1] = lineValues[1]
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return []byte(RespBulkStringEncode(values[1])), nil
}
