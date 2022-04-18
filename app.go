package main

import (
	"context"
	"fmt"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
}

type Table struct {
	A string
	B string
	C string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	getUserEnv()
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test() *Table {

	tData := &Table{}
	tData.A = "Test A"
	tData.B = "Test B"
	tData.C = "Test C"

	return tData
}

func (a *App) AddImage(name string, repo string) {

	fmt.Println("Name: ", name)
	fmt.Println("Repo: ", repo)

	img := Image{}
	img.Name = name
	img.Repo = repo
	status, err := addImageToDB(img)
	if status == false {
		fmt.Println("Error", err)
	}

}

func (a *App) RemoveImage(iid string) {

	id, err := strconv.Atoi(iid)
	if err != nil {
		fmt.Println(err)
	}

	removeImageFromDB(id)

}

func (a *App) GetImage(id string) Image {

	image, status := getImageFromDB(id)
	if status == false {
		fmt.Println("Can't find image...")
	}

	return image

}

func (a *App) ChangeImage(id string, name string, repo string) bool {

	status := true

	result := changeImageInDB(id, name, repo)
	if result == false {
		status = false
		fmt.Println("Can't find image...")
	}

	return status

}

func (a *App) GetImages() ([]ImageInJSON, bool) {

	images, status := getImagesFromDB()
	if status == false {
		fmt.Println("false")
		fmt.Println(images)
		fmt.Println(status)
		images = make([]ImageInJSON, 0)
		return images, status
	}

	return images, status

}

func (a *App) AddEnvironment(image string, name string, hostname string, port string, local string, path string, options string) (bool, string) {

	status := true

	env := Environment{}
	env.Image = image
	env.Name = name
	env.Hostname = hostname
	env.Port = port
	env.Local = local
	env.Path = path
	env.Options = options

	// check to make sure there's no container with the same name
	envs, stat := getEnvsFromDB()
	if stat == false {
		fmt.Println("Getting environments from DB failed...")
		return false, "Getting environments from DB failed..."
	}

	for i := 0; i < len(envs); i++ {
		if env.Name == envs[i].Name {
			fmt.Println("Environment name must be unique...")
			return false, "Environment name must be unique..."
		}
	}

	// Start container
	lc := launchContainer(env)
	if lc != false {

		ae, err := addEnvToDB(env)
		if ae == false {
			fmt.Println("Error", err)
			status = false
		}

	}

	return status, ""

}

func (a *App) RemoveEnvironment(iid string) bool {

	env, status := getEnvFromDB(iid)
	if status == false {
		fmt.Println("Can't find environment...")
	}

	id, err := strconv.Atoi(iid)
	if err != nil {
		fmt.Println(err)
	}

	// Stop container first just in case
	sc := stopContainer(env.Name)
	if sc == false {
		fmt.Println("Container was not stopped...")
	}
	// Remove container
	rc := removeContainer(env.Name)
	if rc == false {
		fmt.Println("Counld't remove container")
		return false
	}

	removeEnvFromDB(id)
	return true

}

func (a *App) AccessEnvironment(name string) bool {

	ac := accessContainer(name)
	if ac == false {
		fmt.Println("Container access error...")
		return false
	}

	return true

}

func (a *App) AccessLogs(name string) bool {

	ac := accessLogs(name)
	if ac == false {
		fmt.Println("Container access error...")
		return false
	}

	return true

}

func (a *App) GetEnvironment(id string) Environment {

	env, status := getEnvFromDB(id)
	if status == false {
		fmt.Println("Can't find environment...")
	}

	return env

}

func (a *App) ChangeEnvironment(
	ids string,
	image string,
	name string,
	hostname string,
	port string,
	local string,
	path string,
	options string) bool {

	status := true

	id, err := strconv.Atoi(ids)
	if err != nil {
		status = false
		fmt.Println("ID type conversion error...")
	}

	env := Environment{}
	env.ID = id
	env.Image = image
	env.Name = name
	env.Hostname = hostname
	env.Port = port
	env.Local = local
	env.Path = path
	env.Options = options

	cc := changeContainer(env)
	if cc != false {
		result := changeEnvInDB(env)
		if result == false {
			status = false
			fmt.Println("Can't find environment...")
		}
	} else {
		status = false
	}

	return status

}

func (a *App) GetEnvironments() ([]EnvInJSON, bool) {

	envs, status := getEnvsFromDB()
	if status == false {
		fmt.Println("false")
		envs = make([]EnvInJSON, 0)
		return envs, status
	}

	for i := 0; i < len(envs); i++ {
		ip := getContainerIP(envs[i].Name)
		envs[i].IP = ip
		if envs[i].Port == "" {
			envs[i].Port = getContainerPorts(envs[i].Name)
		}
		action := checkContainer(envs[i].Name)
		envs[i].Action = action
	}

	return envs, status

}

func (a *App) StartEnv(id string) bool {

	status := true

	env, gstatus := getEnvFromDB(id)
	if gstatus == false {
		status = false
		fmt.Println("getEnv: failed")
	}

	ostatus := startContainer(env.Name)
	ustatus := true

	if ostatus == true {
		env.Action = "stop"
		ustatus = updateAction(id, env)
		if ustatus == false {
			status = false
			fmt.Println("ustatus: failed")
		}
	} else {
		status = false
		fmt.Println("status: failed")
	}

	return status

}

func (a *App) StopEnv(id string) bool {

	status := true

	env, gstatus := getEnvFromDB(id)
	if gstatus == false {
		status = false
		fmt.Println("getEnv: failed")
	}

	ostatus := stopContainer(env.Name)
	ustatus := true

	if ostatus == true {
		env.Action = "start"
		ustatus = updateAction(id, env)
		if ustatus == false {
			status = false
		}
	} else {
		status = false
	}

	return status

}
