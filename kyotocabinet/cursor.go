package kyotocabinet

// #cgo LDFLAGS: -lkyotocabinet
// #include <kclangc.h>
import "C"

import (
	"errors"
	"unsafe"
)

type KCCUR struct {
	cur *C.KCCUR
}

func (kc *KCDB) Cursor() (kcc *KCCUR) {
	cur := C.kcdbcursor(kc.db)
	if cur == nil {
		return nil
	}
	kcc = &KCCUR{cur}
	kcc.Jump()
	return
}

func (kcc *KCCUR) Ecode() int {
	return int(C.kccurecode(kcc.cur))
}

func (kcc *KCCUR) Emsg() string {
	err := C.kccuremsg(kcc.cur)
	return C.GoString(err)
}

func (kcc *KCCUR) error() error {
	return errors.New(kcc.Emsg())
}

func (kcc *KCCUR) Del() {
	C.kccurdel(kcc.cur)
}

func (kcc *KCCUR) Db() (kc *KCDB) {
	kc = new(KCDB)
	kc.db = C.kccurdb(kcc.cur)
	return
}

func (kcc *KCCUR) Jump() (err error) {
	if C.kccurjump(kcc.cur) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) JumpKey(key []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	if C.kccurjumpkey(kcc.cur, ckey, C.size_t(len(key))) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) JumpBack() (err error) {
	if C.kccurjumpback(kcc.cur) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) JumpBackKey(key []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	if C.kccurjumpbackkey(kcc.cur, ckey, C.size_t(len(key))) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) Step() (err error) {
	if C.kccurstep(kcc.cur) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) StepBack() (err error) {
	if C.kccurstepback(kcc.cur) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) Remove() (err error) {
	if C.kccurremove(kcc.cur) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) SetValue(value []byte, advance bool) (err error) {
	cvalue := (*C.char)(unsafe.Pointer(&value[0]))
	var cadvance C.int32_t
	if advance {
		cadvance = 1
	} else {
		cadvance = 0
	}
	if C.kccursetvalue(kcc.cur, cvalue, C.size_t(len(value)), cadvance) == 0 {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) GetKey(advance bool) (k []byte, err error) {
	var ksiz C.size_t
	var cadvance C.int32_t
	if advance {
		cadvance = 1
	} else {
		cadvance = 0
	}
	kp := C.kccurgetkey(kcc.cur, &ksiz, cadvance)
	if kp != nil {
		k = make([]byte, ksiz)
		C.memcpy(unsafe.Pointer(&k[0]), unsafe.Pointer(kp), ksiz)
		C.kcfree(unsafe.Pointer(kp))
	} else {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) GetValue(advance bool) (v []byte, err error) {
	var vsiz C.size_t
	var cadvance C.int32_t
	if advance {
		cadvance = 1
	} else {
		cadvance = 0
	}
	vp := C.kccurgetvalue(kcc.cur, &vsiz, cadvance)
	if vp != nil {
		v = make([]byte, vsiz)
		C.memcpy(unsafe.Pointer(&v[0]), unsafe.Pointer(vp), vsiz)
		C.kcfree(unsafe.Pointer(vp))
	} else {
		err = kcc.error()
	}
	return
}

func (kcc *KCCUR) Get(advance bool) (k, v []byte, err error) {
	var ksiz, vsiz C.size_t
	var vp *C.char
	var cadvance C.int32_t
	if advance {
		cadvance = 1
	} else {
		cadvance = 0
	}
	kp := C.kccurget(kcc.cur, &ksiz, &vp, &vsiz, cadvance)
	if kp != nil {
		k = make([]byte, ksiz)
		C.memcpy(unsafe.Pointer(&k[0]), unsafe.Pointer(kp), ksiz)
		v = make([]byte, vsiz)
		C.memcpy(unsafe.Pointer(&v[0]), unsafe.Pointer(vp), vsiz)
		C.kcfree(unsafe.Pointer(kp))
	} else {
		err = kcc.error()
	}
	return
}
