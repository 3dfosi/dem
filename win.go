// +build windows

package main

import (
	"os/exec"
	"syscall"
)

// Hide the stupid powershell windows on command execution
func hideWindow(cmd *exec.Cmd) {

   cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

}
