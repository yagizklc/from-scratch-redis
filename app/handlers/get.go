package handlers

import (
	"time"
)

func Get(arguments []string) ([]byte, error) {
	object := db[arguments[0]]

	now := time.Now()
	if object.expirationTime != (time.Time{}) && object.expirationTime.Before(now) {
		delete(db, arguments[0])
		return []byte(RespSimpleError("ERR key not found")), nil
	}

	encoded := RespBulkStringEncode(object.value)
	return []byte(encoded), nil
}
