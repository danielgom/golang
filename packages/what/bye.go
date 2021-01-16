package main

import (
	"fmt"
	"runtime"
)

func bye() {
	fmt.Println("Bye")
	fmt.Println(noLoSe)
}

func goVersion() string {
	return runtime.Version()
}
