package tcp

import (
	"fmt"
	"github.com/goodplayer/kv-common/common/util"
	"github.com/goodplayer/kv-common/kyotocabinet"
	"net"
)

const (
	PROTOCOL_HEADER_LENGTH = 8
	PROTOCOL_LENGTH_OFFSET = 4

	CMD_GET        byte = 1
	CMD_SET        byte = 2
	CMD_PAGED_LIST byte = 3

	ERROR_SEND_ERROR      = "1.errno"
	ERROR_MARSHAL_ERROR   = "2.error"
	ERROR_UNMARSHAL_ERROR = "3.error"
	ERROR_UNKNOWN_ERROR   = "255.errno"
)

type KcServer struct {
	openedDb *kyotocabinet.KCDB
	sortDb   *kyotocabinet.KCDB
	listener *net.TCPListener
	closeCh  chan int

	//for further use
	StoreDbList map[string]*kyotocabinet.KCDB
	SortDbList  map[string]*kyotocabinet.KCDB
}

type Connection struct {
	conn *net.TCPConn
}

type ProtoProtocol interface {
	Process(server *KcServer, conn *net.TCPConn)
}

// provide KcServer by yourself
// protocol is :
// 1 byte - version, current is 1
// 1 byte - type, current is:
//          1 : get - BasicOps -> BasicResp
//          2 : set - BasicOps -> BasicResp
//          3 : paged_list - PagedListReq -> PagedListResp
// 2 bytes - not used, current is 0
// 4 bytes - body size, total size of the body following, big-endian
// n bytes - body
func StartServer(kcServer *KcServer) {
	go kcServer.listen_dispatch()
}

func (kcserver *KcServer) StopServer() {
	close(kcserver.closeCh)
	kcserver.listener.Close()
}

func (kcserver *KcServer) send_data(conn *net.TCPConn, cmdType byte, data []byte) {
	header := make([]byte, 8)
	header[0] = 1
	header[1] = cmdType
	header[2] = 0
	header[3] = 0
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

func (kcserver *KcServer) listen_dispatch() {
	listener := kcserver.listener
	ch := kcserver.closeCh
	for {
		_, ok := <-ch
		if ok {
			conn, err := listener.AcceptTCP()
			connection := &Connection{
				conn: conn,
			}
			if err != nil {
				//TODO log it
				break
			}
			// create goroutine to handle connection
			go kcserver.process_request(connection)
			//TODO add conn to idle detector
		} else {
			break
		}
	}
}

func (server *KcServer) process_request(conn *Connection) {
	tcpConn := conn.conn
	defer tcpConn.Close()
	defer handle_panic()
	for {
		header := make([]byte, PROTOCOL_HEADER_LENGTH)
		cnt, err := tcpConn.Read(header)
		if err != nil {
			//TODO log read error
			break
		}
		if cnt != PROTOCOL_HEADER_LENGTH {
			//TODO log header count not match
			break
		}
		if header[0] != 1 {
			//TODO log version not match
			break
		}
		bodyLength, err := util.ToInt32_BigEndian(header[PROTOCOL_LENGTH_OFFSET:])
		if err != nil {
			//TODO log body length convert err
			break
		}
		if bodyLength > 0 {
			body := make([]byte, bodyLength)
			cnt, err = tcpConn.Read(body)
			if err != nil {
				//TODO log read error
				break
			}
			if cnt != bodyLength {
				//TODO log body count not match
				break
			}
			protoProtocol := unmarshal_body(header[1], body)
			if protoProtocol != nil {
				protoProtocol.Process(server, tcpConn)
			} else {
				//TODO log no mapped protocol
			}
		}
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
