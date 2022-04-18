package main

import (
	"encoding/json"
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
	"log"
	"os"
	"os/user"
	"strconv"
)

// Current User
var uname string   // User Name
var homedir string // Home Directory
var confdir string // Configuration Directory
var dbdir string   // DB Directory

type Info struct {
	LastID int
}

type Image struct {
	ID   int
	Name string
	Repo string
}

type ImageInJSON struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Repo string `json:"Repo"`
}

type Environment struct {
	ID       int
	Image    string
	Name     string
	Hostname string
	IP       string
	Port     string
	Local    string
	Path     string
	Options  string
	Action   string
	Status   string
}

type EnvInJSON struct {
	ID       int    `json:"ID"`
	Image    string `json:"Image"`
	Name     string `json:"Name"`
	Hostname string `json:"Hostname"`
	IP       string `json:"IP"`
	Port     string `json:"Port"`
	Local    string `json:"Local"`
	Path     string `json:"Path"`
	Options  string `json:"Options"`
	Action   string `json:"Action"`
	Status   string `json:"Status"`
}

/*
 *
 * USERLAND
 *
 */

func getUserEnv() {

	usr, err := user.Current() // Get Current User
	if err != nil {
		fmt.Print("Get User Error: ")
		fmt.Println(err)
	}
	uname = usr.Username
	homedir = usr.HomeDir
	confdir = homedir + "/.3df-dem"
	dbdir = confdir + "/db"

	// Check if db dir exists. If not, then create dir
	dbinfo, err := os.Stat(dbdir)
	if os.IsNotExist(err) {
		fmt.Println("DB does not exist.")
		initDB()
	}
	log.Println(dbinfo)

}

func initDB() {

	// Make directory if it doesn't exist
	fmt.Println(dbdir)
	err := os.MkdirAll(dbdir, 0700)
	if err != nil {
		log.Fatal(err)
	}

	// Add counter to track how many records are in DB
	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	ic := Info{}
	ic.LastID = 0
	if err := db.Write("images", "info", &ic); err != nil {
		fmt.Println("Error: ", err)
	}

	ec := Info{}
	ec.LastID = 0
	if err := db.Write("environments", "info", &ec); err != nil {
		fmt.Println("Error: ", err)
	}

}

/*
 *
 * IMAGES
 *
 */

// Add single image to DB
func addImageToDB(image Image) (bool, string) {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		return false, err.Error()
	}

	oinfo := Info{}
	if err := db.Read("images", "info", &oinfo); err != nil {
		return false, err.Error()
	}

	ninfo := Info{}
	ninfo.LastID = oinfo.LastID + 1

	image.ID = ninfo.LastID
	if err := db.Write("images", strconv.Itoa(image.ID), &image); err != nil {
		return false, err.Error()
	}

	if err := db.Write("images", "info", &ninfo); err != nil {
		return false, err.Error()
	}

	return true, ""

}

// Change single image in DB
func changeImageInDB(ids string, name string, repo string) bool {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	status := true
	id, err := strconv.Atoi(ids)
	if err != nil {
		status = false
		fmt.Println("ID type conversion error...")
	}

	img := Image{}
	img.ID = id
	img.Name = name
	img.Repo = repo

	if err := db.Write("images", ids, &img); err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	return status

}

// Remove single image from DB
func removeImageFromDB(id int) {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := db.Delete("images", strconv.Itoa(id)); err != nil {
		fmt.Println("Error: ", err)
	}

}

// Get single image from DB
func getImageFromDB(id string) (Image, bool) {

	status := true

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	img := Image{}
	if err := db.Read("images", id, &img); err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	return img, status

}

// Get all images from DB
func getImagesFromDB() ([]ImageInJSON, bool) {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	info := Info{}
	if err := db.Read("images", "info", &info); err != nil {
		fmt.Println("Error: ", err)
	}

	imgs, err := db.ReadAll("images")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//remove info element
	if len(imgs) > 0 {
		imgs = imgs[:len(imgs)-1]
	}

	images := make([]ImageInJSON, len(imgs))
	for i := 0; i < len(imgs); i++ {
		var element = ImageInJSON{}
		err := json.Unmarshal([]byte(imgs[i]), &element)
		if err != nil {
			log.Fatal(err)
		}
		images[i] = element
	}

	return images, true

}

/*
 * ENVIRONMENTS
 *
 */

// Add single environment to DB
func addEnvToDB(env Environment) (bool, string) {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		return false, err.Error()
	}

	oinfo := Info{}
	if err := db.Read("environments", "info", &oinfo); err != nil {
		return false, err.Error()
	}

	ninfo := Info{}
	ninfo.LastID = oinfo.LastID + 1

	env.ID = ninfo.LastID
	if err := db.Write("environments", strconv.Itoa(env.ID), &env); err != nil {
		return false, err.Error()
	}

	if err := db.Write("environments", "info", &ninfo); err != nil {
		return false, err.Error()
	}

	return true, ""

}

// Change single environment in DB
func changeEnvInDB(env Environment) bool {

	status := true

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	ids := strconv.Itoa(env.ID)

	if err := db.Write("environments", ids, &env); err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	return status

}

// Used to update action after container state has changed
func updateAction(id string, env Environment) bool {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	status := true

	if err := db.Write("environments", id, &env); err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	return status

}

// Remove single environment from DB
func removeEnvFromDB(id int) {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := db.Delete("environments", strconv.Itoa(id)); err != nil {
		fmt.Println("Error: ", err)
	}

}

// Get single environment from DB
func getEnvFromDB(id string) (Environment, bool) {

	status := true

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	env := Environment{}
	if err := db.Read("environments", id, &env); err != nil {
		fmt.Println("Error: ", err)
		status = false
	}

	return env, status

}

// Get all environments from DB
func getEnvsFromDB() ([]EnvInJSON, bool) {

	db, err := scribble.New(dbdir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	info := Info{}
	if err := db.Read("environments", "info", &info); err != nil {
		fmt.Println("Error: ", err)
	}

	envs, err := db.ReadAll("environments")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//remove info element
	if len(envs) > 0 {
		envs = envs[:len(envs)-1]
	}

	environments := make([]EnvInJSON, len(envs))
	for i := 0; i < len(envs); i++ {
		var element = EnvInJSON{}
		err := json.Unmarshal([]byte(envs[i]), &element)
		if err != nil {
			log.Fatal(err)
		}
		environments[i] = element
	}

	return environments, true

}
