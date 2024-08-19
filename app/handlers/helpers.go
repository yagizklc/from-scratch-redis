package handlers

import (
	"fmt"
)

var crlf_delim = []byte("\r\n")

const dbPath = "/Users/ykc/Documents/Worksplace/learn/codecrafters-redis-go/app/db.txt"

func RespSimpleStringEncode(message string) string {
	return fmt.Sprintf("+%s%s", message, string(crlf_delim))
}

func RespSimpleError(message string) string {
	return fmt.Sprintf("-%s%s", message, string(crlf_delim))
}

func RespBulkStringEncode(message string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(message), message)
}
