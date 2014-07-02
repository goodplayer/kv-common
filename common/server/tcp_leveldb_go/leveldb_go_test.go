package tcp_leveldb_go

import (
	"code.google.com/p/leveldb-go/leveldb"
	"os"
	// "code.google.com/p/leveldb-go/leveldb/db"
	"fmt"
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

func Test_leveldb_go_use_2(t *testing.T) {
	defer os.RemoveAll("./temp.db")
	openDb, err := leveldb.Open("./temp.db", nil)
	if err != nil {
		t.Error("error occurs", err)
		return
	}
	openDb.Set([]byte("key1"), []byte("value1"), nil)
	openDb.Set([]byte("bbbb"), []byte("bbbbvalue"), nil)
	openDb.Set([]byte("mul"), []byte("mulvalue"), nil)
	iter := openDb.Find([]byte("key1"), nil)
	if iter.Next() {
		fmt.Println(iter.Key())
	}
}
