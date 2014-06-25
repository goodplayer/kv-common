package tcp_common

import (
	"github.com/goodplayer/kv-common/common/util"
	"io"
	"net"
)

const (
	ERROR_SEND_ERROR      = "1.errno"
	ERROR_MARSHAL_ERROR   = "2.errno"
	ERROR_UNMARSHAL_ERROR = "3.errno"
	ERROR_READ_ERROR      = "4.errno"
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
	write_fully(conn, header)
	write_fully(conn, data)
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
		case ERROR_READ_ERROR:
			//TODO log read error
		default:
			panic(err)
		}
	}
}

func write_fully(conn *net.TCPConn, data []byte) {
	n, err := conn.Write(data)
	if err != nil {
		//TODO log write error
		panic(ERROR_SEND_ERROR)
	}
	lenght := len(data)
	if n < lenght {
		idx := n
		for idx < lenght {
			scnt, err := conn.Write(data[idx:])
			if err != nil {
				//TODO log write error
				panic(ERROR_SEND_ERROR)
			}
			idx = idx + scnt
		}
	}
}

func read_fully(conn *net.TCPConn, buf []byte) {
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		//TODO log read error
		panic(ERROR_READ_ERROR)
	}
}
