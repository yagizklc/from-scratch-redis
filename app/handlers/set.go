package handlers

import (
	"bufio"
	"fmt"
	"os"
)

func Set(arguments []string) ([]byte, error) {
	file, err := os.OpenFile(dbPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	m := fmt.Sprintf("%s %s\r\n", arguments[0], arguments[1])
	_, err = writer.WriteString(m)
	if err != nil {
		return nil, err
	}

	if err := writer.Flush(); err != nil {
		return nil, err
	}

	return []byte(RespSimpleStringEncode("OK")), nil
}
