package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	s := "hello"
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // prints 5

	// read the bytes in the string using "pointer arithmetic"
	sHdrData := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(uintptr(sHdrData) + uintptr(i)))
		fmt.Print(string(bp))
	}
	fmt.Println()

	sHdr.Len = sHdr.Len + 10
	fmt.Println(s)
	// don't garbage collect s until after the call to KeepAlive
	runtime.KeepAlive(s)

	// bp := (*byte)(unsafe.Pointer(uintptr(sHdrData) + 2))
	// *bp = *bp + 1
	// fmt.Println(s)

	// slices
	sl := []int{10, 20, 30}
	slHdr := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	fmt.Println(slHdr.Len) // prints 3
	fmt.Println(slHdr.Cap) // prints 3
	intByteSize := unsafe.Sizeof(sl[0])
	fmt.Println(intByteSize)
	for i := 0; i < slHdr.Len; i++ {
		intVal := *(*int)(unsafe.Pointer(slHdr.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
	runtime.KeepAlive(sl)
}
