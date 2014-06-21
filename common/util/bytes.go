package util

import (
	"errors"
)

func ToInt32_BigEndian(data []byte) (int32, error) {
	if len(data) < 4 {
		return 0, errors.New("data length is less than 4")
	}
	var result int32 = 0
	result = result | ((int32(data[0]) & 0xff) << 24)
	result = result | ((int32(data[1]) & 0xff) << 16)
	result = result | ((int32(data[2]) & 0xff) << 8)
	result = result | (int32(data[3]) & 0xff)
	return result, nil
}

func ToBytesFromInt32_BigEndian(data int32) (r []byte) {
	r = make([]byte, 4)
	r[0] = byte((data >> 24) & 0xff)
	r[1] = byte((data >> 16) & 0xff)
	r[2] = byte((data >> 8) & 0xff)
	r[3] = byte(data & 0xff)
	return
}
