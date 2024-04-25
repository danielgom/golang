package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/100_go_mistakes/embedding/embed"
	"os"
)

func main() {

	m := &embed.MyMemExample[int, string]{}
	fmt.Println(m)
	m.Lock() // ?? We shouldn't make this functionality visible

	m2 := &embed.MyMemExampleHidden[int, string]{}
	// m2.Lock() We successfully hidden the mutex implementation
	fmt.Println(m2)

	logger := embed.Logger{WriteCloser: os.Stdout}
	_, _ = logger.Write([]byte{1, 2, 3}) // Using the created methods

	myLogger := embed.MyLogger{WriteCloser: os.Stdout}
	_, _ = myLogger.Write([]byte{1, 2, 3}) // Using the standard library methods
}
