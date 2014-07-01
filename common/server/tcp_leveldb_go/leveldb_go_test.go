package tcp_leveldb_go

import (
	"code.google.com/p/leveldb-go/leveldb"
	"os"
	// "code.google.com/p/leveldb-go/leveldb/db"
	"testing"
)

func Test_leveldb_go_basic_use(t *testing.T) {
	defer os.RemoveAll("./temp.db")
	openDb, err := leveldb.Open("./temp.db", nil)
	if err != nil {
		t.Error("error occurs", err)
		return
	}
	err = openDb.Set([]byte("aa"), []byte("bb"), nil)
	if err != nil {
		t.Error("error occurs", err)
		return
	}
	data, err := openDb.Get([]byte("aa"), nil)
	if err != nil {
		t.Error("error occurs", err)
		return
	}
	if string(data) != "bb" {
		t.Error("not equal")
	}
}
