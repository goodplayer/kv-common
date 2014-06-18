package kyotocabinet

// #cgo LDFLAGS: -lkyotocabinet
// #include <kclangc.h>
import "C"

import ()

func Version() string {
	return C.GoString(C.KCVERSION)
}

func EcodeName(ecode int) string {
	return C.GoString(C.kcecodename(C.int32_t(ecode)))
}
