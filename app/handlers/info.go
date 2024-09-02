package handlers

import (
	"fmt"
	"github.com/yagizklc/from-scratch-redis/app/pkg"
	"strings"
)

func Info(args []string) ([]byte, error) {
	lenArgs := len(args)
	if lenArgs > 1 {
		return nil, fmt.Errorf("too many arguments for INFO command: expected 1, got %d", lenArgs)
	}
	if lenArgs == 0 {
		return nil, fmt.Errorf("too few arguments for INFO command: expected 1, got %d", lenArgs)
	}

	var response string
	if strings.ToLower(args[0]) == "replication" {
		response = fmt.Sprintf(
			"# Replication\nrole:%s\nconnected_slaves:%d\nmaster_replid:%s",
			"master",
			0,
			"8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb",
		)

	}

	return []byte(pkg.RespBulkStringEncode(response)), nil

}

/*
$ redis-cli INFO replication
# Replication
role:master
connected_slaves:0
master_replid:8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:
*/
