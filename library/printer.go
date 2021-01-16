package library

import (
	"fmt"
	"runtime"
)

var Number int32 = 29

func PrintSomething() {
	fmt.Println("hello from other library")
}

func ReturnGoVersion() string {
	return runtime.Version()
}
