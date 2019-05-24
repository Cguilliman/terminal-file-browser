package manager

import (
	"io/ioutil"
	"errors"
	"log"
	"os"
	"strconv"
)

type Manager struct {
	Path              string
	Files             []os.FileInfo
	CurrentFileNumber int
}

func (manager *Manager) RenderList() []string {
	var response []string

	for n, file := range manager.Files {
		var fileName string
		switch n {
		case 0:
			fileName = "[.(Current)](fg:blue)"
		case 1:
			fileName = "[..(Parent)](fg:blue)"
		default:
			fileName = file.Name()
		}

		row := fileName + " " + strconv.Itoa(int(file.Size()))
		if n == manager.CurrentFileNumber {
			row = ">> " + row
		}
		if file.IsDir() {
			row = "[" + row + "](fg:blue)"
		}
		response = append(response, row)
	}

	return response
}

func (manager *Manager) SetFiles() {
	base_files := []os.FileInfo{
		getFile(manager.Path),
		getFile(ParentDirPath(manager.Path)),
	}
	files, err := ioutil.ReadDir(manager.Path)
	if err != nil {
		log.Fatal(err)
	}

	files = append(base_files, files...)
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

func (manager *Manager) EnterDir() error {
	file := manager.Files[manager.CurrentFileNumber]
	if !file.IsDir() {
		return errors.New("This is file!")
	}
	switch manager.CurrentFileNumber{
		case 0:
			return nil
		case 1:
			manager.Path = ParentDirPath(manager.Path)
		default:
			manager.Path = ConcatPath(manager.Path, file.Name())
	}
	manager.CurrentFileNumber = 0 // reset current file number
	manager.SetFiles()
	return nil
}
