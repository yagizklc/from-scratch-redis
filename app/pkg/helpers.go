package pkg

import (
	"fmt"
)

func RespSimpleStringEncode(message string) string {
	return fmt.Sprintf("+%s\r\n", message)
}

func RespSimpleError(message string) string {
	return fmt.Sprintf("-%s\r\n", message)
}

func RespBulkStringEncode(message string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(message), message)
}
