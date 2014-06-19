package util

import (
	"testing"
)

func Test_bytes_to_int32(t *testing.T) {
	data := []byte{0x4a, 0x64, 0x62, 0x53}
	var i int32 = 1248092755
	r, err := ToInt32_BigEndian(data)
	if err != nil {
		t.Error("err occurs")
		return
	}
	if r != i {
		t.Error("not match")
	}
}
