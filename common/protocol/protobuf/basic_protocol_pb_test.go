package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
	"testing"
)
import "fmt"

func Test_BasicOps_Marshal_Unmarshal(t *testing.T) {
	key := []byte{5, 7, 2, 9}
	value := []byte{3, 4, 5, 6, 7, 8, 9}

	msg := &BasicOps{}
	msg.Key = key
	msg.Value = value

	data, _ := proto.Marshal(msg)

	msgUnmarshal := &BasicOps{}
	proto.Unmarshal(data, msgUnmarshal)

	if !slice_equal(key, msgUnmarshal.GetKey(), len(key)) {
		t.Error("BasicOps key not match", key, msgUnmarshal.GetKey())
		return
	}
	if !slice_equal(value, msgUnmarshal.GetValue(), len(key)) {
		t.Error("BasicOps value not match", value, msgUnmarshal.GetValue())
		return
	}
}

func Test_BasicResp_Marshal_Unmarshal(t *testing.T) {
	key := []byte{5, 7, 2, 9}
	value := []byte{3, 4, 5, 6, 7, 8, 9}
	var code int32 = 1

	msg := &BasicResp{}
	msg.ResponseCode = &code
	msg.Key = key
	msg.Value = value

	data, _ := proto.Marshal(msg)

	msgUn := &BasicResp{}
	proto.Unmarshal(data, msgUn)

	if code != *msgUn.ResponseCode {
		t.Error("BasicResp responseCode not match", code, *msgUn.ResponseCode)
		return
	}
	if !slice_equal(key, msgUn.GetKey(), len(key)) {
		t.Error("BasicResp key not match", key, msgUn.GetKey())
		return
	}
	if !slice_equal(value, msgUn.GetValue(), len(value)) {
		t.Error("BasicResp value not match", value, msgUn.GetValue())
		return
	}
}

func slice_equal(expect, actual []byte, length int) bool {
	if len(actual) < length || len(expect) < length {
		return false
	}
	for i := 0; i < length; i++ {
		if expect[i] != actual[i] {
			return false
		}
	}
	return true
}
