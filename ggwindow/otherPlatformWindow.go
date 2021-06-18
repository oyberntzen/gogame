// +build !windows

package ggwindow

import (
	"github.com/oyberntzen/gogame/ggcore"
)

func init() {
	ggcore.CoreWarn("GoGame does not support other platforms than windows")
}

func NewWindow(props WindowProps) Window {
	return nil
}
