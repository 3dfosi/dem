// +build !windows

package main

import "os/exec"

// Hide the stupid powershell windows on command execution
func hideWindow(cmd *exec.Cmd) {

	// this is not needed for *nix and mac

}
