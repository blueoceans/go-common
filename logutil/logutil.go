// +build !appengine

package logutil

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"unicode/utf8"
)

const ()

func sprintf(format string, args ...interface{}) string {
	_, file, line, _ := runtime.Caller(2)
	var fileLine = fmt.Sprintf("%s:%d]", file, line)
	return fmt.Sprintf("%s %s", fileLine, fmt.Sprintf(format, args...))
}

func Debugf(c appengine.Context, format string, args ...interface{}) {
	c.Debugf(sprintf(format, args...))
}

func Infof(c appengine.Context, format string, args ...interface{}) {
	c.Infof(sprintf(format, args...))
}

func StackTracef(c appengine.Context, format string, args ...interface{}) {
	c.Infof(sprintf("%s", StackTrace(2, 2048)))
	c.Infof(sprintf(format, args...))
}

//https://github.com/knightso/base/blob/master/errors/errors.go
func StackTrace(skip, maxBytes int) []byte {
	// this func is debug purpose and ignores errors

	buf := make([]byte, maxBytes)
	n := runtime.Stack(buf, false)
	var gotall bool
	if n < len(buf) {
		buf = buf[:n]
		gotall = true
	} else {
		for !utf8.Valid(buf) || len(buf) == 0 {
			buf = buf[:len(buf)-1]
		}
	}

	var w bytes.Buffer

	writeOrSkip := func(buf []byte, w io.Writer, line int) {
		if line == 1 || line > 1+skip*2 {
			w.Write(buf)
		}
	}

	line := 1
	for { //EOFまでloopする
		lf := bytes.IndexByte(buf, '\n')
		if lf < 0 {
			writeOrSkip(buf, &w, line)
			break
		}
		writeOrSkip(buf[:lf+1], &w, line)
		buf = buf[lf+1:]
		line++
	}

	if !gotall {
		w.WriteString("\n        ... (omitted)")
	}
	w.WriteString("\n")

	return w.Bytes()
}
