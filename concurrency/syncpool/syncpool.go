package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	/*
		We can create pools in order to avoid creating new instances on heavy processes
	*/

	/*
		Creating pool of the bytes.Buffer instance, sync.Pool, provides a New field that returns a new interface{} type
		which means that we can create a pool of any instance we want

	*/

	var bufPool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Allocating new buffer")
			var b bytes.Buffer
			return &b // We are selecting to return an instance of bytes.buffer
			// return new(bytes.Buffer) this does the same as the line above
		},
	}

	/*
		In the end of the program we see that the line of "Allocating new buffer" just appears once, so the second
		time we call logs , it does not create a new instance, it just uses the instance that is inside the pool bufPool
	*/

	logs(&bufPool, os.Stdout, "debug-string1")
	logs(&bufPool, os.Stdout, "debug-string2")
}

func logs(b *sync.Pool, w io.Writer, debug string) {
	//var b bytes.Buffer

	c := b.Get().(*bytes.Buffer) // Getting from the syncPool and converting it to the proper type (did this on interfaces)

	c.Reset() // We reset the buffer to be empty

	c.WriteString(time.Now().Format("15:04:15"))
	c.WriteString(" : ")
	c.WriteString(debug)
	c.WriteString("\n")
	_, err := w.Write(c.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	b.Put(c) // We put back the buffer, back to the buffer pool, in this case the pool of bytes.Buffer
}
