// +build !profile

package ggdebug

func BeginSession(name, path string) {}
func EndSession()                    {}

type Timer struct{}

func Start() *Timer     { return &Timer{} }
func Stop(timer *Timer) {}
