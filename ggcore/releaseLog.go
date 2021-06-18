// +build release

package ggcore

func Trace(format string, a ...interface{}) {}
func Info(format string, a ...interface{})  {}
func Warn(format string, a ...interface{})  {}
func Error(format string, a ...interface{}) {}
func Fatal(format string, a ...interface{}) {}
func CheckError(err error)                  {}

func CoreTrace(format string, a ...interface{}) {}
func CoreInfo(format string, a ...interface{})  {}
func CoreWarn(format string, a ...interface{})  {}
func CoreError(format string, a ...interface{}) {}
func CoreFatal(format string, a ...interface{}) {}
func CoreCheckError(err error)                  {}
