package handlers

import ()

var db = make(map[string]string)

func Set(args []string) ([]byte, error) {
	if len(args) < 2 {
		return []byte(RespSimpleError("ERR wrong number of arguments for 'set' command")), nil
	}
	db[args[0]] = args[1]
	return []byte(RespSimpleStringEncode("OK")), nil
}
