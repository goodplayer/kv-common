package kyotocabinet

import (
	"fmt"
	"testing"
)

func Test_newDB(t *testing.T) {
	kc := New()
	defer kc.Del()
	err := kc.Open("test_db.kch", KCOWRITER|KCOCREATE)
	if err != nil {
		t.Error(err)
	}
	err = kc.Set([]byte("hello"), []byte("world"))
	if err != nil {
		t.Error(err)
	}
	world, err := kc.Get([]byte("hello"))
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(world))
	}
	err = kc.Close()
	if err != nil {
		t.Error(t)
	}
}

func Test_order_hash(t *testing.T) {
	order_template("test_db.kch", "hash", t)
}
func Test_order_tree(t *testing.T) {
	order_template("test_db.kct", "tree", t)
}

func order_template(filename string, testcase string, t *testing.T) {
	fmt.Println("=========", testcase, "=========")
	kc := New()
	defer kc.Del()
	err := kc.Open(filename, KCOWRITER|KCOCREATE|KCOREADER)
	if err != nil {
		t.Error(err)
		return
	}
	err = kc.Set([]byte("fgsddsbdfsb"), []byte("va"))
	err = kc.Set([]byte("mnsbkk"), []byte("vc"))
	err = kc.Set([]byte("gihoklkl"), []byte("vb"))
	err = kc.Set([]byte("gfheh"), []byte("vb"))
	err = kc.Set([]byte("javkjsavjehfdg"), []byte("vb"))
	err = kc.Set([]byte("sfdjewn"), []byte("vb"))
	err = kc.Set([]byte("ewkfjnsdnf"), []byte("vb"))
	if err != nil {
		t.Error(err)
		return
	}
	cursor := kc.Cursor()
	cursor.Jump()
	k, v, e := cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	k, v, e = cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	k, v, e = cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	k, v, e = cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	k, v, e = cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	k, v, e = cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	k, v, e = cursor.Get(true)
	fmt.Println(string(k), string(v), e)
	cursor.Del()
	fmt.Println(kc.Status())
	kc.Close()
}
