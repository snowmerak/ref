package ref

import (
	"reflect"
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

func New(value interface{}) unsafe.Pointer {
	return new(value)
}

func Next(r unsafe.Pointer, n uintptr, s uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(r) + n*s)
}

func new(value interface{}) unsafe.Pointer {
	v := reflect.ValueOf(value)
	s := v.Elem().Type().Size()
	p := libuseful.Alloc(s)
	libuseful.MemMove(p, unsafe.Pointer(v.Pointer()), s)
	return p
}

func Release(r unsafe.Pointer) {
	release(r)
}

func release(r unsafe.Pointer) {
	//libuseful.Free(r.pointer, r.size)
}
