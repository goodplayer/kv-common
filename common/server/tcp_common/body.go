package tcp_common

import (
	"errors"
)

const (
	BODY_ERROR_LENGTH_NEGATIVE = "tcp_common: body length is negative"
	BODY_ERROR_COUNT_NOT_MATCH = "tcp_common: body count not match"
)

var (
	EMTPY_DATA = make([]byte, 0)
)

func ReadBody(client ClientChannel, bodyLength int32) (data []byte, err error) {
	if bodyLength > 0 {
		var cnt int
		conn := client.GetConn()
		data = make([]byte, bodyLength)
		read_fully(conn, data)
		return
	} else if bodyLength < 0 {
		//TODO log body length is negative
		err = errors.New(BODY_ERROR_LENGTH_NEGATIVE)
		return
	}
	return EMTPY_DATA, nil
}
