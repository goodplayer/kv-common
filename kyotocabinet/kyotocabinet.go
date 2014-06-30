package kyotocabinet

// #cgo LDFLAGS: -lkyotocabinet
// #include <kclangc.h>
import "C"

import (
	"errors"
	"log"
	"runtime"
	"unsafe"
)

const KCOREADER int = C.KCOREADER
const KCOWRITER int = C.KCOWRITER
const KCOCREATE int = C.KCOCREATE
const KCOTRUNCATE int = C.KCOTRUNCATE
const KCOAUTOTRAN int = C.KCOAUTOTRAN
const KCOAUTOSYNC int = C.KCOAUTOSYNC
const KCONOLOCK int = C.KCONOLOCK
const KCOTRYLOCK int = C.KCOTRYLOCK
const KCONOREPAIR = C.KCONOREPAIR

const KCESUCCESS int = C.KCESUCCESS
const KCENOIMPL int = C.KCENOIMPL
const KCEINVALID int = C.KCEINVALID
const KCENOREPOS int = C.KCENOREPOS
const KCEBROKEN int = C.KCEBROKEN
const KCENOPERM int = C.KCENOPERM
const KCEDUPREC int = C.KCEDUPREC
const KCENOREC int = C.KCENOREC
const KCELOGIC int = C.KCELOGIC
const KCESYSTEM = C.KCESYSTEM
const KCEMISC = C.KCEMISC

const KCMADD int = C.KCMADD
const KCMSET int = C.KCMSET
const KCMAPPEND int = C.KCMAPPEND
const KCMREPLACE int = C.KCMREPLACE

type KCDB struct {
	db *C.KCDB
}

func New() *KCDB {
	db := C.kcdbnew()
	return &KCDB{db}
}

func (kc *KCDB) Ecode() int {
	return int(C.kcdbecode(kc.db))
}

func (kc *KCDB) error() error {
	return errors.New(EcodeName(kc.Ecode()))
}

func (kc *KCDB) OpenHashDb(filename string, mode int) (err error) {
	return kc.Open(filename+".kch", mode)
}

func (kc *KCDB) OpenTreeDb(filename string, mode int) (err error) {
	return kc.Open(filename+".kct", mode)
}

func (kc *KCDB) Open(filename string, mode int) (err error) {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	if C.kcdbopen(kc.db, name, C.uint32_t(mode)) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Close() (err error) {
	if C.kcdbclose(kc.db) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Del() {
	C.kcdbdel(kc.db)
}

func (kc *KCDB) Path() (path string, err error) {
	cpath := C.kcdbpath(kc.db)
	defer C.kcfree(unsafe.Pointer(cpath))
	path = C.GoString(cpath)
	if path == "" {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Status() (status string, err error) {
	cstatus := C.kcdbstatus(kc.db)
	if cstatus == nil {
		err = kc.error()
	}
	status = C.GoString(cstatus)
	C.kcfree(unsafe.Pointer(cstatus))
	return
}

func (kc *KCDB) Count() (count uint64, err error) {
	ccount := C.kcdbcount(kc.db)
	if ccount == -1 {
		err = kc.error()
	} else {
		count = uint64(ccount)
	}
	return
}

func (kc *KCDB) Size() (size uint64, err error) {
	csize := C.kcdbsize(kc.db)
	if csize == -1 {
		err = kc.error()
	} else {
		size = uint64(csize)
	}
	return
}

func (kc *KCDB) Clear() (err error) {
	if C.kcdbclear(kc.db) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Sync(hard bool) (err error) {
	var chard C.int32_t
	if hard {
		chard = 1
	} else {
		chard = 0
	}
	if C.kcdbsync(kc.db, chard, nil, nil) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Copy(filename string) (err error) {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	if C.kcdbcopy(kc.db, name) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Merge(sdbs []*KCDB, mode int) (err error) {
	count := len(sdbs)
	csdbs := make([]*C.KCDB, count)
	for i, db := range sdbs {
		csdbs[i] = db.db
	}
	if C.kcdbmerge(kc.db, &csdbs[0], C.size_t(count), C.uint32_t(mode)) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Dump(filename string) (err error) {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	if C.kcdbdumpsnap(kc.db, name) == 0 {
		err = kc.error()
	}
	return
}
func (kc *KCDB) Load(filename string) (err error) {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	if C.kcdbcopy(kc.db, name) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) BeginTran(hard bool) (err error) {
	var chard C.int32_t
	if hard {
		chard = 1
	} else {
		chard = 0
	}
	if C.kcdbbegintran(kc.db, chard) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) BeginTranTry(hard bool) (err error) {
	var chard C.int32_t
	if hard {
		chard = 1
	} else {
		chard = 0
	}
	if C.kcdbbegintrantry(kc.db, chard) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) EndTran(commit bool) (err error) {
	var ccommit C.int32_t
	if commit {
		ccommit = 1
	} else {
		ccommit = 0
	}
	if C.kcdbendtran(kc.db, ccommit) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Set(key, value []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	cvalue := (*C.char)(unsafe.Pointer(&value[0]))
	if C.kcdbset(kc.db, ckey, C.size_t(len(key)), cvalue, C.size_t(len(value))) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Add(key, value []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	cvalue := (*C.char)(unsafe.Pointer(&value[0]))
	if C.kcdbadd(kc.db, ckey, C.size_t(len(key)), cvalue, C.size_t(len(value))) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Replace(key, value []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	cvalue := (*C.char)(unsafe.Pointer(&value[0]))
	if C.kcdbreplace(kc.db, ckey, C.size_t(len(key)), cvalue, C.size_t(len(value))) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Append(key, value []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	cvalue := (*C.char)(unsafe.Pointer(&value[0]))
	if C.kcdbappend(kc.db, ckey, C.size_t(len(key)), cvalue, C.size_t(len(value))) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Remove(key []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	if C.kcdbremove(kc.db, ckey, C.size_t(len(key))) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Cas(key, oval, nval []byte) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	var coval, cnval *C.char
	loval, lnval := len(oval), len(nval)
	if loval > 0 {
		coval = (*C.char)(unsafe.Pointer(&oval[0]))
	}
	if lnval > 0 {
		cnval = (*C.char)(unsafe.Pointer(&nval[0]))
	}
	if C.kcdbcas(kc.db, ckey, C.size_t(len(key)), coval, C.size_t(loval), cnval, C.size_t(lnval)) == 0 {
		err = kc.error()
	}
	return
}

func (kc *KCDB) IncrInt(key []byte, amount int64) (result int64, err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	cresult := C.kcdbincrint(kc.db, ckey, C.size_t(len(key)), C.int64_t(amount), 0)

	if cresult == C.INT64_MIN {
		err = kc.error()
	} else {
		result = int64(cresult)
	}
	return
}

func (kc *KCDB) IncrDouble(key []byte, amount float64) (err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	if C.kcdbincrdouble(kc.db, ckey, C.size_t(len(key)), C.double(amount), C.double(0.0)) == C.kcnan() {
		err = kc.error()
	}
	return
}

func (kc *KCDB) Get(key []byte) (value []byte, err error) {
	ckey := (*C.char)(unsafe.Pointer(&key[0]))
	var vlen C.size_t
	cval := C.kcdbget(kc.db, ckey, C.size_t(len(key)), &vlen)
	if cval == nil {
		err = kc.error()
	} else {
		value = make([]byte, int(vlen))
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(cval), vlen)
		C.kcfree(unsafe.Pointer(cval))
	}
	return
}

func (kc *KCDB) Keys() (out chan []byte) {
	out = make(chan []byte)
	go func() {
		cur := kc.Cursor()
		for {
			k, err := cur.GetKey(true)
			if err != nil {
				if cur.Ecode() != KCENOREC {
					log.Printf("GetKey(true): %s", err)
				}
				break
			}
			out <- k
		}
		cur.Del()
		close(out)
	}()
	return
}

func (kc *KCDB) Values() (out chan []byte) {
	out = make(chan []byte)
	go func() {
		cur := kc.Cursor()
		for {
			v, err := cur.GetValue(true)
			if err != nil {
				if cur.Ecode() != KCENOREC {
					log.Printf("GetValue(true): %s", err)
				}
				break
			}
			out <- v
		}
		cur.Del()
		close(out)
	}()
	return
}

type Item struct {
	Key   []byte
	Value []byte
}

func (kc *KCDB) Items() (out chan *Item) {
	out = make(chan *Item)
	go func() {
		cur := kc.Cursor()
		for {
			k, v, err := cur.Get(true)
			if err != nil {
				if cur.Ecode() != KCENOREC {
					log.Printf("Get(true): %s", err)
				}
				break
			}
			out <- &Item{k, v}
		}
		cur.Del()
		close(out)
	}()
	return
}

func (kc *KCDB) MatchPrefix(prefix string, max int) (matches [][]byte, err error) {
	cprefix := C.CString(prefix)
	strarray := make([]*C.char, max)
	count := C.kcdbmatchprefix(kc.db, cprefix, &strarray[0], C.size_t(max))
	if count == -1 {
		matches = nil
		err = kc.error()
	} else {
		matches = make([][]byte, int(count))
		for i := int64(0); i < int64(count); i++ {
			matches[i] = []byte(C.GoString(strarray[i]))
			C.kcfree(unsafe.Pointer(strarray[i]))
		}
	}
	C.free(unsafe.Pointer(cprefix))
	return
}

func (kc *KCDB) MatchRegex(regex string, max int) (matches [][]byte, err error) {
	cregex := C.CString(regex)
	strarray := make([]*C.char, max)
	count := C.kcdbmatchregex(kc.db, cregex, &strarray[0], C.size_t(max))
	if count == -1 {
		matches = nil
		err = kc.error()
	} else {
		matches = make([][]byte, count)
		for i := int64(0); i < int64(count); i++ {
			matches[i] = []byte(C.GoString(strarray[i]))
			C.kcfree(unsafe.Pointer(strarray[i]))
		}
	}
	C.free(unsafe.Pointer(cregex))
	return
}
