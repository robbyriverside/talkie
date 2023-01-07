package talkie

import (
	"fmt"
	"os"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type ColorWriter struct {
	out   *os.File
	color string
}

func NewColorWriter(color string) *ColorWriter {
	return &ColorWriter{
		out:   os.Stdout,
		color: color,
	}
}

func NewColorErrorWriter(color string) *ColorWriter {
	return &ColorWriter{
		out:   os.Stderr,
		color: color,
	}
}

func (cw ColorWriter) Write(msg []byte) (n int, err error) {
	out := cw.color + string(msg) + Reset
	return cw.out.Write([]byte(out))
}

func (cw ColorWriter) Print(ar ...any) (n int, err error) {
	return fmt.Fprint(cw, ar...)
}

func (cw ColorWriter) Println(ar ...any) (n int, err error) {
	return fmt.Fprintln(cw, ar...)
}

func (cw ColorWriter) Printf(format string, ar ...any) (n int, err error) {
	return fmt.Fprintf(cw, format, ar...)
}

func (cw ColorWriter) WrapSize() int {
	return len(cw.color) + len(Reset)
}
