package tcp

import (
	"code.google.com/p/goprotobuf/proto"
	prot "github.com/goodplayer/kv-common/common/protocol/protobuf"
	"net"
)

const (
	BASIC_RESP_CODE_SUCCESS   = 1
	BASIC_RESP_CODE_NOT_EXIST = 2
	BASIC_RESP_CODE_SET_ERROR = 3

	DUMMY_DATA = []byte{0}
)

func (e *GETSET) Process(server *KcServer, conn *net.TCPConn) {
	switch e.cmd {
	case CMD_GET:
		key := e.ops.GetKey()
		if key == nil || len(key) == 0 {
			//TODO log key is nil or size = 0
			return
		}
		value, err := server.openedDb.GET(key)
		if err != nil {
			//TODO log get data error from db
			return
		}
		r := &prot.BasicResp{}
		r.Key = key
		if value == nil || len(value) == 0 {
			var notExist int32 = BASIC_RESP_CODE_NOT_EXIST
			r.ResponseCode = &notExist
		} else {
			var sucess int32 = BASIC_RESP_CODE_SUCCESS
			r.ResponseCode = &sucess
			r.Value = value
		}
		data, err := proto.Marshal(r)
		if err != nil {
			//TODO log error when marshalling
			return
		}
		cnt, err := conn.Write(data)
		if err != nil {
			//TODO log error when write data
			return
		}
		if cnt != len(data) {
			//TODO log write data length unmatched
			return
		}
		return
	case CMD_SET:
		//TODO
		key := e.ops.GetKey()
		if key == nil || len(key) == 0 {
			//TODO log key is nil or size = 0
			return
		}
		value := e.ops.GetValue()
		if value == nil || len(value) == 0 {
			//TODO log value is nil or size =0
			return
		}
		err := server.openedDb.Set(key, value)
		if err != nil {
			//TODO log set error
			r := &prot.BasicResp{}
			var setError int32 = BASIC_RESP_CODE_SET_ERROR
			r.ResponseCode = &setError
			r.Key = key
			data, err := proto.Marshal(r)
			if err != nil {
				//TODO log marshal error
				return
			}
			cnt, err := conn.Write(data)
			if err != nil {
				//TODO log send error
				return
			}
			if cnt != len(data) {
				//TODO log data cnt not match
			}
			return
		}
		err = server.sortDb.Set(key, DUMMY_DATA)
		if err != nil {
			//TODO log set error
			r := &prot.BasicResp{}
			var setError int32 = BASIC_RESP_CODE_SET_ERROR
			r.ResponseCode = &setError
			r.Key = key
			data, err := proto.Marshal(r)
			if err != nil {
				//TODO log marshal error
				return
			}
			cnt, err := conn.Write(data)
			if err != nil {
				//TODO log send error
				return
			}
			if cnt != len(data) {
				//TODO log data cnt not match
			}
			return
		}
		r := &prot.BasicResp{}
		var success int32 = BASIC_RESP_CODE_SUCCESS
		r.ResponseCode = &success
		r.Key = key
		data, err := proto.Marshal(r)
		if err != nil {
			//TODO log marshal error
			return
		}
		cnt, err := conn.Write(data)
		if err != nil {
			//TODO log send error
			return
		}
		if cnt != len(data) {
			//TODO log data cnt not match
		}
		return
	default:
		//TODO log cmd unknown
		return
	}
}

func (e *PagedListProto) Process(server *KcServer, conn *net.TCPConn) {
	//TODO
}
