package tcp_common

import (
	"github.com/goodplayer/kv-common/common/util"
	"net"
)

const (
	ERROR_SEND_ERROR      = "1.errno"
	ERROR_MARSHAL_ERROR   = "2.error"
	ERROR_UNMARSHAL_ERROR = "3.error"
	ERROR_UNKNOWN_ERROR   = "255.errno"
)

// client - version, cmdType, b3, b4, bodyLength, error
type ReadHeaderType func(ClientChannel) (byte, byte, byte, byte, int32, error)

// client, bodyLength - data, error
type ReadBodyType func(ClientChannel, int32) ([]byte, error)

// client, version, cmdType, b3, b4, data
type SendResponseType func(ClientChannel, byte, byte, byte, byte, []byte)

type ClientChannel interface {
	GetConn() *net.TCPConn
	StartProcess()
}

type Processor interface {
	Process(client ClientChannel, version, cmdType, b3, b4 byte, body []byte, sendResp SendResponseType) error
}

type ClientChannelImpl struct {
	Conn             *net.TCPConn
	HeaderReaderFunc ReadHeaderType
	BodyReaderFunc   ReadBodyType
	processor        Processor
}

func CreateClientChannel(conn *net.TCPConn, processor Processor) ClientChannel {
	clientChannel := &ClientChannelImpl{
		Conn:             conn,
		HeaderReaderFunc: ReadHeader,
		BodyReaderFunc:   ReadBody,
		processor:        processor,
	}
	return clientChannel
}

func (client *ClientChannelImpl) GetConn() *net.TCPConn {
	return client.Conn
}

func (client *ClientChannelImpl) StartProcess() {
	defer client.Conn.Close()
	defer handle_panic()
	readHeaderFunc := client.HeaderReaderFunc
	readBodyFunc := client.BodyReaderFunc
	processor := client.processor
	for true {
		version, cmdType, b3, b4, bodyLength, err := readHeaderFunc(client)
		if err != nil {
			//TODO log read header error
			break
		}
		data, err := readBodyFunc(client, bodyLength)
		if err != nil {
			//TODO log read body error
			break
		}
		err = processor.Process(client, version, cmdType, b3, b4, data, send_data_default)
		if err != nil {
			//TODO log process error
			break
		}
	}
}

func send_data_default(client ClientChannel, version, cmdType, b3, b4 byte, data []byte) {
	send_data(client, 1, cmdType, 0, 0, data)
}

func send_data(client ClientChannel, version, cmdType, b3, b4 byte, data []byte) {
	conn := client.GetConn()
	header := make([]byte, 8)
	header[0] = version
	header[1] = cmdType
	header[2] = b3
	header[3] = b4
	sizeArr := util.ToBytesFromInt32_BigEndian(int32(len(data)))
	header[4] = sizeArr[0]
	header[5] = sizeArr[1]
	header[6] = sizeArr[2]
	header[7] = sizeArr[3]
	cnt, err := conn.Write(header)
	if err != nil {
		//TODO log write error
		panic(ERROR_SEND_ERROR)
	}
	if cnt != 8 {
		//TODO log cnt not match
		panic(ERROR_SEND_ERROR)
	}
	cnt, err = conn.Write(data)
	if err != nil {
		//TODO log write error
		panic(ERROR_SEND_ERROR)
	}
	if cnt != len(data) {
		//TODO log cnt not match
		panic(ERROR_SEND_ERROR)
	}
}

func handle_panic() {
	err := recover()
	if err != nil {
		switch err {
		case ERROR_SEND_ERROR:
			//TODO log send data error
		case ERROR_UNKNOWN_ERROR:
			//TODO log unknown error
		case ERROR_MARSHAL_ERROR:
			//TODO log marshal error
		case ERROR_UNMARSHAL_ERROR:
			//TODO log unmarshal error
		default:
			panic(err)
		}
	}
}
