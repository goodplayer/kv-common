package tcp_common

import (
	"errors"
	"github.com/goodplayer/kv-common/common/util"
)

const (
	PROTOCOL_HEADER_LENGTH = 8
	PROTOCOL_LENGTH_OFFSET = 4
)

const (
	HEADER_ERROR_COUNT_NOT_MATCH = "tcp_common: header count not match"
)

func ReadHeader(client ClientChannel) (version, cmdType, b3, b4 byte, bodySize int32, err error) {
	var cnt int
	conn := client.GetConn()
	data := make([]byte, PROTOCOL_HEADER_LENGTH)
	cnt, err = conn.Read(data)
	if err != nil {
		//TODO log read header error
		return
	}
	if cnt != PROTOCOL_HEADER_LENGTH {
		//TODO log read cnt not match
		err = errors.New(HEADER_ERROR_COUNT_NOT_MATCH)
		return
	}
	version = data[0]
	cmdType = data[1]
	b3 = data[2]
	b4 = data[3]
	bodySize, err = util.ToInt32_BigEndian(data[PROTOCOL_LENGTH_OFFSET:])
	return
}
