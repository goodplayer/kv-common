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
)

type KcServer struct {
	openedDb *kyotocabinet.KCDB
	sortDb   *kyotocabinet.KCDB
	listener *net.TCPListener
	closeCh  chan int
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
//          1 : get
//          2 : set
//          3 : paged_list
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
