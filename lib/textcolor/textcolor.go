package textcolor

import "github.com/fatih/color"

var Green func(format string, a ...interface{}) = color.New(color.FgGreen).PrintfFunc()

var Magenta func(format string, a ...interface{}) = color.New(color.FgMagenta).PrintfFunc()

var Cyan func(format string, a ...interface{}) = color.New(color.FgCyan).PrintfFunc()

var Blue func(format string, a ...interface{}) = color.New(color.FgBlue).PrintfFunc()

var Yellow func(format string, a ...interface{}) = color.New(color.FgYellow).PrintfFunc()

var Red func(format string, a ...interface{}) = color.New(color.FgRed).PrintfFunc()
