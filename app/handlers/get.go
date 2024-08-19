package handlers

func Get(arguments []string) ([]byte, error) {
	value := db[arguments[0]]
	encoded := RespBulkStringEncode(value)
	return []byte(encoded), nil
}
