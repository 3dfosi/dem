package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var wcmd = "powershell"
var wflag = "-Command"
var wwf = "-windowstyle"
var wwfv = "hidden"
var ucmd = "sh"
var uflag = "-c"
var mcmd = [2]string{"/usr/bin/osascript -e 'tell application \"TERMINAL\" to do script \"", "\"'"}

// get command based on runtime
func getCmd() (string, string) {

	c := ucmd
	f := uflag

	if runtime.GOOS == "windows" {
		c = wcmd
		f = wflag
	}

	return c, f

}

// get default terminal emulator in Linux
func getLinuxTerm() string {

	term := ""

	deb := exec.Command("x-terminal-emulator", "--help")
	if derr := deb.Run(); derr == nil {
		term = "x-terminal-emulator"
	}

	xdg := exec.Command("xdg-terminal", "--help")
	if xerr := xdg.Run(); xerr == nil {
		term = "xdg-terminal"
	}

	return term
}

// access contianer
func accessContainer(name string) bool {

	state := checkContainer(name)
	if state == "start" {
		fmt.Println("Please start container first")
		return false
	}

	c := "docker exec -ti " + name + " bash"

	if runtime.GOOS == "linux" {

		t := getLinuxTerm()
		la := exec.Command(t, "-e", c)
		if lerr := la.Run(); lerr != nil {
			fmt.Println("Access Error: ", lerr)
			return false
		}

		return true

	}

	if runtime.GOOS == "darwin" {

		mc := mcmd[0] + c + mcmd[1]
		ma := exec.Command("sh", "-c", mc)
		if merr := ma.Run(); merr != nil {
			fmt.Print("Access Error: ", merr)
			return false
		}

		return true

	}

	if runtime.GOOS == "windows" {

		wa := exec.Command("cmd", "/k", "start", "powershell", "-Command", c)
		if werr := wa.Run(); werr != nil {
			fmt.Print("Access Error: ", werr)
			return false
		}
		return true

	}

	return false

}

// access contianer logs
func accessLogs(name string) bool {

	state := checkContainer(name)
	if state == "start" {
		fmt.Println("Please start container first")
		return false
	}

	c := "docker logs --tail 100 -f " + name

	if runtime.GOOS == "linux" {

		t := getLinuxTerm()
		la := exec.Command(t, "-e", c)
		if lerr := la.Run(); lerr != nil {
			fmt.Println("Access Error: ", lerr)
			return true
		}

		return true

	}

	if runtime.GOOS == "darwin" {

		mc := mcmd[0] + c + mcmd[1]
		ma := exec.Command("sh", "-c", mc)
		if merr := ma.Run(); merr != nil {
			fmt.Print("Access Error: ", merr)
			return false
		}

		return true

	}

	if runtime.GOOS == "windows" {

		wa := exec.Command("cmd", "/k", "start", "powershell", "-Command", c)
		hideWindow(wa)
		if werr := wa.Run(); werr != nil {
			fmt.Print("Access Error: ", werr)
			return false
		}

		return true

	}

	return false

}

// check if the contianer is running or not
func checkContainer(name string) string {

	// label for action button in UI
	action := "start"

	cs, fs := getCmd()

	// check to make sure the container exists
	exist := "docker ps -a --format \"table {{.Names}}\\t{{.Status}}\" |grep " + name
	if runtime.GOOS == "darwin" {
		exist = "/usr/local/bin/" + exist
	}
	ecmd := exec.Command(cs, fs, exist)
	if runtime.GOOS == "windows" {
		hideWindow(ecmd)
	}
	e := ecmd.Run()
	if e != nil {
		action = "error"
		fmt.Println("check existance: ", e)
		return action
	}

	// check if the contianer is running
	check := "docker ps --format \"table {{.Names}}\\t{{.Status}}\" |grep " + name
	if runtime.GOOS == "darwin" {
		check = "/usr/local/bin/" + check
	}
	ccmd := exec.Command(cs, fs, check)
	if runtime.GOOS == "windows" {
		hideWindow(ccmd)
	}
	c := ccmd.Run()
	if c == nil {
		action = "stop"
		fmt.Println("check state: ", e)
		return action
	}

	return action

}

// get IP of single container
func getContainerIP(name string) string {

	cs, fs := getCmd()

	command := "docker inspect -f \"{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}\" " + name
	if runtime.GOOS == "darwin" {
		command = "/usr/local/bin/" + command
	}
	cmd := exec.Command(cs, fs, command)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}

	o, err := cmd.Output()
	if err != nil {
		return err.Error()
		// return "output error"
	}

	cmd.Run()

	output := string(o)

	if len(output) == 1 {
		return "none"
	}

	return output

}

// get port of single container that's not published
func getContainerPorts(name string) string {

	cs, fs := getCmd()

	command := "docker ps --format \"table {{.Ports}} , {{.Names}}\" |grep " + name
	if runtime.GOOS == "darwin" {
		command = "/usr/local/bin/" + command
	}
	cmd := exec.Command(cs, fs, command)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}

	o, err := cmd.Output()
	if err != nil {
		return err.Error()
		// return "output error"
	}

	cmd.Run()

	rawout := string(o)

	if len(rawout) == 1 {
		return "none"
	}

	roa := strings.Split(rawout, " , ")
	if len(roa) < 1 {
		fmt.Println("No ports...")
		return "none"
	}

	output := roa[0]

	return output

}

// launch container
func launchContainer(env Environment) bool {

	status := true

	img := strings.Split(env.Image, " | ")
	if len(img) < 1 {
		fmt.Println("Invalid repo...")
		return false
	}

	fmt.Println("PORT: ", env.Port)
	port := " -p " + env.Port
	if env.Port == "" {
		port = ""
	}

	cs, fs := getCmd()
	launch := "docker run -d --name " + env.Name + " --hostname " + env.Hostname + port + " -v " + env.Local + ":" + env.Path + " " + env.Options + " " + img[1]
	if runtime.GOOS == "darwin" {
		launch = "/usr/local/bin/" + launch
	}
	cmd := exec.Command(cs, fs, launch)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}

	c := cmd.Run()

	if c != nil {
		status = false
	}

	return status

}

// change container
func changeContainer(env Environment) bool {

	status := true

	stopContainer(env.Name)
	removeContainer(env.Name)

	img := strings.Split(env.Image, " | ")
	if len(img) < 1 {
		fmt.Println("Invalid repo...")
		return false
	}

	port := " -p " + env.Port
	if env.Port == "" {
		port = ""
	}

	cs, fs := getCmd()
	launch := "docker run -d --name " + env.Name + " --hostname " + env.Hostname + port + " -v " + env.Local + ":" + env.Path + " " + env.Options + " " + img[1]
	if runtime.GOOS == "darwin" {
		launch = "/usr/local/bin/" + launch
	}
	cmd := exec.Command(cs, fs, launch)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}

	c := cmd.Run()

	fmt.Println(c)

	if c != nil {
		status = false
	}

	return status

}

// start container
func startContainer(name string) bool {

	status := true

	cs, fs := getCmd()
	start := "docker start " + name
	if runtime.GOOS == "darwin" {
		start = "/usr/local/bin/" + start
	}
	cmd := exec.Command(cs, fs, start)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}
	c := cmd.Run()

	if c != nil {
		fmt.Println("c != nil")
		status = false
	}

	return status

}

// stop container
func stopContainer(name string) bool {

	status := true

	cs, fs := getCmd()
	start := "docker stop " + name
	if runtime.GOOS == "darwin" {
		start = "/usr/local/bin/" + start
	}
	cmd := exec.Command(cs, fs, start)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}

	c := cmd.Run()

	if c != nil {
		status = false
	}

	return status

}

// remove container
func removeContainer(name string) bool {

	status := true

	cs, fs := getCmd()
	start := "docker rm " + name
	if runtime.GOOS == "darwin" {
		start = "/usr/local/bin/" + start
	}
	cmd := exec.Command(cs, fs, start)
	if runtime.GOOS == "windows" {
		hideWindow(cmd)
	}

	c := cmd.Run()

	if c != nil {
		status = false
	}

	return status

}
