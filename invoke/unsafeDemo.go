package main

import (
	"fmt"
	"unsafe"
)

/*
void setInt(int addr,int val){
	int* ptr = (int*)(addr);
	*ptr = val;
}
int getInt(int addr){
	int* ptr = (int*)(addr);
	return *ptr;
}
*/
import "C"

const n = 4

var s [n]int32 = [n]int32{1, 2, 3, 4}
var offset int = (int)(unsafe.Sizeof(int32(0))) // 偏移量
var head uintptr = uintptr(unsafe.Pointer(&s))  // 数组首地址

func init() {
	for i := 0; i < n; i++ {
		C.setInt(C.int(int(head)+offset*i), C.int(i))                       // 用 c 的方法 set
		*(*int32)(unsafe.Pointer(uintptr(int(head) + i*offset))) = int32(i) // 用 go 的方法 set
	}
}
func main() {
	fmt.Println(s)
	fmt.Println(head)
	for i := 0; i < n; i++ {
		var val int32
		val = int32(C.getInt(C.int(int(head) + i*offset)))             // c 的方法 get
		val = *(*int32)(unsafe.Pointer(uintptr(int(head) + i*offset))) // go 的方法 get
		fmt.Printf("%d,%d\n", i, val)
	}
}
