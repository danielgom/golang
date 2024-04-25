package embed

import (
	"io"
	"sync"
)

type MyMemExample[T comparable, V any] struct {
	sync.Mutex // embedded, this also implements the Locker interface
	m          map[T]V
}

type MyMemExampleHidden[T comparable, V any] struct {
	mu sync.Mutex // Not embedded, no functionality exported, no interface implemented
	m  map[T]V
}

// Embedded good example

type Logger struct {
	WriteCloser io.WriteCloser // not embedded, must implement methods
}

func (l Logger) Write(p []byte) (n int, err error) {
	return l.WriteCloser.Write(p)
}

func (l Logger) Close() error {
	return l.WriteCloser.Close()
}

type MyLogger struct {
	io.WriteCloser // embedded, no need to implement Write, Close methods
}
