package ref

import (
	"unsafe"

	"github.com/lemon-mint/libuseful"
)

var heap Array = Array{
	pointer:   unsafe.Pointer(uintptr(0)),
	length:    0,
	blockSize: 0,
}

var pageSize = 64 * MB
var segmentSize = 32 * Byte

func SetPageSize(b uintptr) {
	pageSize = b
}

func SetSegmentSize(b uintptr) {
	segmentSize = b
}

func newPage() {
	if heap.length == 0 {
		heap = Array{
			pointer:   libuseful.Alloc(pageSize),
			length:    1,
			blockSize: pageSize,
		}
		return
	}
	heap.Extend(1)
}
