// +build !release

package ggcore

import (
	"fmt"
	"time"

	"gopkg.in/gookit/color.v1"
)

func Trace(format string, a ...interface{}) {
	t := time.Now()
	color.White.Printf("[%v] APP: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func Info(format string, a ...interface{}) {
	t := time.Now()
	color.Green.Printf("[%v] APP: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
	t := time.Now()
	color.Yellow.Printf("[%v] APP: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	t := time.Now()
	color.Red.Printf("[%v] APP: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
	panic(nil)
}

func Fatal(format string, a ...interface{}) {
	t := time.Now()
	color.BgRed.Printf("[%v] APP: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func CheckError(err error) {
	if err != nil {
		Error(err.Error())
		panic("")
	}
}

func CoreTrace(format string, a ...interface{}) {
	t := time.Now()
	color.White.Printf("[%v] GoGame: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func CoreInfo(format string, a ...interface{}) {
	t := time.Now()
	color.Green.Printf("[%v] GoGame: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func CoreWarn(format string, a ...interface{}) {
	t := time.Now()
	color.Yellow.Printf("[%v] GoGame: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func CoreError(format string, a ...interface{}) {
	t := time.Now()
	color.Red.Printf("[%v] GoGame: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
	panic(nil)
}

func CoreFatal(format string, a ...interface{}) {
	t := time.Now()
	color.BgRed.Printf("[%v] GoGame: %v\n", t.Format("15:04:05"), fmt.Sprintf(format, a...))
}

func CoreCheckError(err error) {
	if err != nil {
		CoreError(err.Error())
		panic("")
	}
}
