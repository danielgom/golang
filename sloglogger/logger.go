package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"time"
)

type ctxKey string

const (
	slogFields ctxKey = "slog_fields"
)

type something struct {
	foo int
}

type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying
// handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Println("over here")
	/*
		traceAppendedCtx := AppendCtx(ctx, slog.String("first key",
			"yay"), slog.String("second_key", "yay"))

	*/

	/*
		if attrs, ok := traceAppendedCtx.Value(slogFields).([]slog.Attr); ok {
			for _, v := range attrs {
				r.AddAttrs(v)
			}
		}

	*/

	r.AddAttrs(slog.String("first key", "yay"), slog.String("second_key", "yay"))

	fmt.Println("over here 2")
	err := thisf()
	if err != nil {
		return err
	}

	defer func() {
		fmt.Println("over here 3")
	}()

	return h.Handler.Handle(ctx, r)
}

func thisf() error {
	return errors.New("this is an error")
}

func main() {
	initLog()
	/*
		h := &ContextHandler{slog.NewJSONHandler(os.Stdout, nil)}

		logger := slog.New(h)

	*/

	//ctx := AppendCtx(context.Background(), slog.String("request_id", "req-123"))

	ctx := context.Background()

	slog.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))
	slog.InfoContext(ctx, "something else")
	slog.Info("nocontext")

	unix := time.Unix(1577858400, 0).UTC()
	fmt.Println(unix)
	fmt.Println(unix.Format("2006-01-02T15:04:05Z"))

	ints := make([]int, 0, 10)
	ints = append(ints, 1, 2, 3, 4, 5)
	fmt.Println(len(ints))
	fmt.Println(cap(ints))

	fmt.Println(cap(ints))
	fmt.Println(cap(slices.Clip(ints)))

	fmt.Println()
	ints2 := make([]*something, 10)
	fmt.Println(ints2)
}

func initLog() {
	output := os.Stderr
	options := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := &ContextHandler{slog.NewTextHandler(output, options)}

	slog.SetDefault(slog.New(handler))
}
