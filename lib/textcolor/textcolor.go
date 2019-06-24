package textcolor

import (
	"io"

	"github.com/fatih/color"
)

var Green func(io.Writer, string, ...interface{}) = color.New(color.FgGreen).FprintfFunc()
var HiGreen func(io.Writer, string, ...interface{}) = color.New(color.FgHiGreen).FprintfFunc()

var Magenta func(io.Writer, string, ...interface{}) = color.New(color.FgMagenta).FprintfFunc()
var HiMagenta func(io.Writer, string, ...interface{}) = color.New(color.FgHiMagenta).FprintfFunc()

var Cyan func(io.Writer, string, ...interface{}) = color.New(color.FgCyan).FprintfFunc()
var HiCyan func(io.Writer, string, ...interface{}) = color.New(color.FgHiCyan).FprintfFunc()

var Blue func(io.Writer, string, ...interface{}) = color.New(color.FgBlue).FprintfFunc()
var HiBlue func(io.Writer, string, ...interface{}) = color.New(color.FgHiBlue).FprintfFunc()

var Yellow func(io.Writer, string, ...interface{}) = color.New(color.FgYellow).FprintfFunc()
var HiYellow func(io.Writer, string, ...interface{}) = color.New(color.FgHiYellow).FprintfFunc()

var Red func(io.Writer, string, ...interface{}) = color.New(color.FgRed).FprintfFunc()
var HiRed func(io.Writer, string, ...interface{}) = color.New(color.FgHiRed).FprintfFunc()
