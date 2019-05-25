package manager

import (
	"log"
	"os"
	"strings"
)

func getFile(path string) os.FileInfo {
	file, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err.Error)
	}
	return file
}

func getLocalPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func ConcatPath(path, file string) string {
	if string(path[len(path)-1]) != string("/") {
		path = path + "/"
	}
	return path + file
}

func ParentDirPath(path string) string {
	if string(path[len(path)-1]) == string("/") {
		path = path[:len(path)-2]
	}
	dirs := strings.Split(path, "/")
	return strings.Join(dirs[:len(dirs)-1], "/")
}

func InitManager(path string) *Manager {
	if path == "" {
		path = getLocalPath()
	}

	var manager Manager
	manager.Path = path
	manager.SetFiles()
	manager.CurrentFileNumber = 0
	
	return &manager
}
