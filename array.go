package ref

import (
	"reflect"
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

type Array struct {
	pointer           unsafe.Pointer
	blockSize         uintptr
	length            uintptr
	autoReleaseSwitch bool
}

func NewArray(v interface{}, n uintptr) Array {
	s := reflect.TypeOf(v).Size()
	p := libuseful.Alloc(s * n)
	return Array{
		pointer:   p,
		blockSize: s,
		length:    n,
	}
}

func (a Array) autoRelease() {
	if a.autoReleaseSwitch {
		a.Release()
	}
}

func (a Array) AutoRelease() Array {
	a.autoReleaseSwitch = true
	return a
}

func (a Array) Len() uintptr {
	return a.length
}

func (a Array) At(n uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(a.pointer) + a.blockSize*n)
}

func (a Array) Extend(n uintptr) Array {
	na := Array{
		pointer:   libuseful.Alloc(a.blockSize * (a.length + n)),
		blockSize: a.blockSize,
		length:    a.length + n,
	}
	libuseful.MemMove(na.pointer, a.pointer, a.length*a.blockSize)
	a.autoRelease()
	return na
}

func (a Array) Foreach(f func(p unsafe.Pointer)) Array {
	for i := uintptr(0); i < a.length; i++ {
		f(a.At(i))
	}
	return a
}

func (a Array) Map(fn func(f unsafe.Pointer, t unsafe.Pointer), t interface{}) Array {
	s := reflect.TypeOf(t).Size()
	na := Array{
		pointer:   libuseful.Alloc(s * a.length),
		blockSize: s,
		length:    a.length,
	}
	for i := uintptr(0); i < a.length; i++ {
		fn(a.At(i), na.At(i))
	}
	a.autoRelease()
	return na
}

func (a Array) Release() {
	libuseful.Free(a.pointer, a.blockSize*a.length)
}
