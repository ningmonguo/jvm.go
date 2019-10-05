package vmutils

import (
	"unsafe"
)

func CastInt8sToBytes(s []int8) []byte {
	ptr := unsafe.Pointer(&s)
	return *((*[]byte)(ptr))
}

func CastBytesToInt8s(s []byte) []int8 {
	ptr := unsafe.Pointer(&s)
	return *((*[]int8)(ptr))
}

func CastBytesToUint32s(s []byte) []uint32 {
	ptr := unsafe.Pointer(&s)
	return *((*[]uint32)(ptr))
}

func CastBytesToInt32s(s []byte) []int32 {
	ptr := unsafe.Pointer(&s)
	return *((*[]int32)(ptr))
}
