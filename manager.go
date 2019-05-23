package main

import (
	"reflect"
	"strconv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	_ = fmt.Println
	_ = reflect.TypeOf
)

type Manager struct {
	Path  string
	Files []os.FileInfo
	CurrentFileNumber int
}

func (manager *Manager) RenderList() []string {
	var response []string
	for n, file := range manager.Files {
		row := file.Name() + " " + strconv.Itoa(int(file.Size()))
		if n == manager.CurrentFileNumber {
			row = ">> " + row
		}
		response = append(response, row)
	}
	return response
}

func (manager *Manager) GetFiles() {
	files, err := ioutil.ReadDir(manager.Path)
	if err != nil {
		log.Fatal(err)
	}
	manager.Files = files
}

func (manager *Manager) NaxtFile() {
	if len(manager.Files)-1 > manager.CurrentFileNumber {
		manager.CurrentFileNumber++
	}
}

func (manager *Manager) PrevFile() {
	if 0 < manager.CurrentFileNumber {
		manager.CurrentFileNumber--
	}
}

func (manager *Manager) SelectDir()  {
	file := manager.Files[manager.CurrentFileNumber]
	if !file.IsDir() {
		fmt.Println("zalupa")  // re-factor to normal error output
		return 
	}

	newPath := manager.Path
	if string("/") != string(newPath[len(newPath)-1]) {
		newPath = newPath + "/"
	}
	newPath = newPath + file.Name() + "/"
	
	manager.Path = newPath
	manager.GetFiles()
}

func getLocalPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func InitManager() *Manager {
	var manager Manager
	manager.Path = getLocalPath()
	manager.GetFiles()
	manager.CurrentFileNumber = 0
	return &manager
}
