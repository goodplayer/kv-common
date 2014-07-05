package tcp_kyotocabinet

import (
	"github.com/goodplayer/kv-common/kyotocabinet"
)

func Version() string {
	return kyotocabinet.Version()
}
