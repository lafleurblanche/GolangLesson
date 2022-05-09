package main

/*
#cgo LDFLAGS: -L. -lhello
#include <hello.h>
*/
import "C"

//export cgo_hello
func cgo_hello() {
    C.hello()
}

func main() {}
