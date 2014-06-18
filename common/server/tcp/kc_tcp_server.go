package tcp

import (
	"github.com/goodplayer/kv-common/kyotocabinet"
	"net"
)

type KcServer struct {
	openedDb *kytocabinet.KCDB
	listener *net.TCPListener
	closeCh  chan int
}

type Connection struct {
	conn *net.TCPConn
}

// an opened db must be provided
// db close should be controlled by yourself
// protocol is :
// 1 byte - version, current is 1
// 1 byte - type, current is 1
// 2 bytes - not used, current is 0
// 4 bytes - body size, total size of the body following, big-endian
// n bytes - body
func StartServer(openedDb *kytocabinet.KCDB, addr *net.TCPAddr) (*KcServer, error) {
	listener, err := net.ListenTCP(addr.Network(), addr)
	if err != nil {
		return nil, err
	}
	kcServer := &KcServer{
		openedDb: openedDb,
		listener: listener,
		closeCh:  make(chan int, 1),
	}
	go kcServer.listen_dispatch()
	return kcServer, nil
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
			go process_request(connection)
			//TODO add conn to idle detector
		} else {
			break
		}
	}
}

func process_request(conn *Connection) {
	//TODO
}
