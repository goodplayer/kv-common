package tcp_kyotocabinet

import (
	"code.google.com/p/goprotobuf/proto"
	prot "github.com/goodplayer/kv-common/common/protocol/protobuf"
	"net"
)

const (
	BASIC_RESP_CODE_SUCCESS            = 1
	BASIC_RESP_CODE_NOT_EXIST          = 2
	BASIC_RESP_CODE_SET_ERROR          = 3
	BASIC_RESP_CODE_NO_KEY_SPECIFIED   = 4
	BASIC_RESP_CODE_NO_VALUE_SPECIFIED = 5
	BASIC_RESP_CODE_ERROR_OCCURS       = 255
)

var (
	DUMMY_DATA = []byte{0}
)

// business error:
func (e *GETSET) Process(server *KcServer, conn *net.TCPConn) {
	switch e.cmd {
	case CMD_GET:
		_process_get(e, server, conn)
	case CMD_SET:
		_process_set(e, server, conn)
	default:
		//TODO log cmd unknown
		panic(ERROR_UNKNOWN_ERROR)
	}
}

func _process_get(e *GETSET, server *KcServer, conn *net.TCPConn) {
	r := &prot.BasicResp{}
	for true {
		key := e.ops.GetKey()
		if key == nil || len(key) == 0 {
			var code int32 = BASIC_RESP_CODE_NO_KEY_SPECIFIED
			r.ResponseCode = &code
			//TODO log key is nil or size = 0
			break
		}
		value, err := server.openedDb.Get(key)
		if err != nil {
			var code int32 = BASIC_RESP_CODE_ERROR_OCCURS
			r.ResponseCode = &code
			//TODO log get data error from db
			break
		}
		r.Key = key
		if value == nil || len(value) == 0 {
			var notExist int32 = BASIC_RESP_CODE_NOT_EXIST
			r.ResponseCode = &notExist
		} else {
			var sucess int32 = BASIC_RESP_CODE_SUCCESS
			r.ResponseCode = &sucess
			r.Value = value
		}
		break
	}
	data, err := proto.Marshal(r)
	if err != nil {
		//TODO log error when marshalling
		panic(ERROR_MARSHAL_ERROR)
	}
	server.send_data(conn, CMD_GET, data)
}

func _process_set(e *GETSET, server *KcServer, conn *net.TCPConn) {
	r := &prot.BasicResp{}
	for true {
		key := e.ops.GetKey()
		if key == nil || len(key) == 0 {
			var code int32 = BASIC_RESP_CODE_NO_KEY_SPECIFIED
			r.ResponseCode = &code
			//TODO log key is nil or size = 0
			break
		}
		value := e.ops.GetValue()
		if value == nil || len(value) == 0 {
			var code int32 = BASIC_RESP_CODE_NO_VALUE_SPECIFIED
			r.ResponseCode = &code
			//TODO log value is nil or size =0
			break
		}
		err := server.openedDb.Set(key, value)
		if err != nil {
			//TODO log set error
			var setError int32 = BASIC_RESP_CODE_SET_ERROR
			r.ResponseCode = &setError
			r.Key = key
			break
		}
		if e.ops.GetSorted() {
			err = server.sortDb.Set(key, DUMMY_DATA)
			if err != nil {
				//TODO log set error
				var setError int32 = BASIC_RESP_CODE_SET_ERROR
				r.ResponseCode = &setError
				r.Key = key
				break
			}
		}
		var success int32 = BASIC_RESP_CODE_SUCCESS
		r.ResponseCode = &success
		r.Key = key
	}
	data, err := proto.Marshal(r)
	if err != nil {
		//TODO log marshal error
		panic(ERROR_MARSHAL_ERROR)
	}
	server.send_data(conn, CMD_SET, data)
}

func (e *PagedListProto) Process(server *KcServer, conn *net.TCPConn) {
	//TODO
	// req := e.ops
}
