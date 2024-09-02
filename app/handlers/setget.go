package handlers

import (
	"github.com/yagizklc/from-scratch-redis/app/pkg"
	"strconv"
	"strings"
	"time"
)

type KeyValue struct {
	value          string
	expirationTime time.Time
}

var db = make(map[string]KeyValue)

func Set(args []string) ([]byte, error) {
	if len(args) < 2 {
		return []byte(pkg.RespSimpleError("ERR wrong number of arguments for 'set' command")), nil
	}

	expirationTime := time.Time{}
	if len(args) == 4 && strings.ToLower(args[2]) == "px" {
		expiry, err := strconv.Atoi(args[3])
		if err != nil {
			return []byte(pkg.RespSimpleError("ERR value is not an integer or out of range")), nil
		}
		if expiry <= 0 {
			return []byte(pkg.RespSimpleError("ERR invalid expire time in set")), nil
		}
		expirationTime = time.Now().Add(time.Millisecond * time.Duration(expiry))
	}

	db[args[0]] = KeyValue{value: args[1], expirationTime: expirationTime}
	return []byte(pkg.RespSimpleStringEncode("OK")), nil
}

func Get(arguments []string) ([]byte, error) {
	object := db[arguments[0]]

	now := time.Now()
	if object.expirationTime != (time.Time{}) && object.expirationTime.Before(now) {
		delete(db, arguments[0])
		return []byte(pkg.RespSimpleError("ERR key not found")), nil
	}

	return []byte(pkg.RespBulkStringEncode(object.value)), nil
}
