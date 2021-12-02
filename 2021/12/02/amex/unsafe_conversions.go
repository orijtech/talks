package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func unsafeByteSliceToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func unsafeStrToByteSlice(s string) (b []byte) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Cap = len(s)
	hdr.Len = len(s)
	hdr.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	return b
}

func main() {
	fmt.Printf("%q\n", unsafeByteSliceToStr([]byte("string")))
	fmt.Printf("% x\n", unsafeStrToByteSlice("string"))
}
