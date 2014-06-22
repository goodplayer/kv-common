package tcp

import (
	"code.google.com/p/goprotobuf/proto"
	prot "github.com/goodplayer/kv-common/common/protocol/protobuf"
)

func unmarshal_body(cmdType byte, data []byte) ProtoProtocol {
	switch cmdType {
	case CMD_GET:
		fallthrough
	case CMD_SET:
		p := &prot.BasicOps{}
		err := proto.Unmarshal(data, p)
		if err != nil {
			//TODO log unmarshal error
			return nil
		}
		r := &GETSET{
			cmd: cmdType,
			ops: p,
		}
		return r
	case CMD_PAGED_LIST:
		p := &prot.PagedListReq{}
		err := proto.Unmarshal(data, p)
		if err != nil {
			//TODO log unmarshal error
			return nil
		}
		r := &PagedListProto{
			ops: p,
		}
		return r
	default:
		return nil
	}
}

type GETSET struct {
	cmd byte
	ops *prot.BasicOps
}

type PagedListProto struct {
	ops *prot.PagedListReq
}
