package manager

import (
	"errors"
	"github.com/Cguilliman/terminal-file-browser/utils"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Manager struct {
	Path              string
	Files, Searchable []os.FileInfo
	CurrentFileNumber int
	Highlighting      []int
}

func (manager *Manager) PickOutFile() {
	if manager.CurrentFileNumber > 1 {
		if utils.InIntArray(manager.CurrentFileNumber, manager.Highlighting) {
			manager.Highlighting = utils.RemoveFromIntArray(
				manager.CurrentFileNumber,
				manager.Highlighting,
			)
		} else {
			manager.Highlighting = append(
				manager.Highlighting,
				manager.CurrentFileNumber,
			)
		}
	}
}

func (manager *Manager) PickOutAllFiles() {
	allFiles := []int{}
	for i := 2; i != len(manager.Files); i++ {
		allFiles = append(allFiles, i)
	}
	if len(manager.Highlighting) == 0 && len(allFiles) > 0 {
		manager.Highlighting = allFiles
	} else {
		manager.Highlighting = []int{}
	}
}

func (manager *Manager) DelHighlighting() {
	manager.Highlighting = []int{}
}

// Return array of strings for render
// *get files from manger.Files array
func (manager *Manager) RenderList(fileList []os.FileInfo) []string {
	if len(fileList) == 0 {
		fileList = manager.Files
	}
	var response []string

	for n, file := range fileList {
		var fileName string
		switch n {
		case 0:
			fileName = ".(Current)"
		case 1:
			fileName = "..(Parent)"
		default:
			fileName = file.Name()
		}

		row := fileName + " " + strconv.Itoa(int(file.Size()))
		if utils.InIntArray(n, manager.Highlighting) {
			row = "[" + row + "*](fg:green)"
		} else if file.IsDir() {
			row = "[" + row + "](fg:blue)"
		}
		if n == manager.CurrentFileNumber {
			row = ">> " + row
		}
		response = append(response, row)
	}

	return response
}

// Return default current directory
// and parent directory `os.FileIngo` objects list
func (manager *Manager) defaultFiles() []os.FileInfo {
	return []os.FileInfo{
		utils.GetFile(manager.Path),
		utils.GetFile(utils.ParentDirPath(manager.Path)),
	}
}

// Set files in manage.Files list of files objects
func (manager *Manager) SetFiles() {
	base_files := manager.defaultFiles()
	files, err := ioutil.ReadDir(manager.Path)
	if err != nil {
		// TODO re-factor error logging
		log.Fatal(err)
	}

	files = append(base_files, files...)
	manager.Files = files
}

// Next file. Change only CurrentFileNumber param.
func (manager *Manager) NextFile() []string {
	if len(manager.Files)-1 > manager.CurrentFileNumber {
		manager.CurrentFileNumber++
	}
	return manager.RenderList(nil)
}

// Previous file. Change only CurrentFileNumber param.
func (manager *Manager) PrevFile() []string {
	if 0 < manager.CurrentFileNumber {
		manager.CurrentFileNumber--
	}
	return manager.RenderList(nil)
}

// First file. Change only CurrentFileNumber param.
func (manager *Manager) SetFirstFile() []string {
	manager.CurrentFileNumber = 0
	return manager.RenderList(nil)
}

// Last file. Change only CurrentFileNumber param.
func (manager *Manager) SetLastFile() []string {
	manager.CurrentFileNumber = len(manager.Files) - 1
	return manager.RenderList(nil)
}

// Enter directory
// change `manger.Path` as current inner director
// inner directory files and save as `manager.Files`
func (manager *Manager) EnterDir(isParent bool) ([]string, error) {
	var fileNumber int
	if isParent {
		fileNumber = 1
	} else {
		fileNumber = manager.CurrentFileNumber
	}
	file := manager.Files[fileNumber]

	if !file.IsDir() {
		return nil, errors.New("This is file!")
	}

	switch fileNumber {
	case 0:
		return manager.RenderList(nil), nil
	case 1:
		manager.Path = utils.ParentDirPath(manager.Path)
	default:
		manager.Path = utils.ConcatPath(manager.Path, file.Name())
	}

	manager.CurrentFileNumber = 0
	manager.SetFiles()
	return manager.RenderList(nil), nil
}

func (manager *Manager) GetDirPath(number int) string {
	if number < 0 {
		number = manager.CurrentFileNumber
	}
	return utils.ConcatPath(
		manager.Path,
		manager.Files[number].Name(),
	)
}

func (manager *Manager) Search(searchChan chan string, renderChan chan UpdateData) {
	manager.Searchable = manager.Files
	for searchable := range searchChan {
		manager.CurrentFileNumber = 0
		manager.Files = manager.defaultFiles()

		for _, obj := range manager.Searchable[2:] {
			// TODO: re-factor searching
			if strings.Contains(obj.Name(), searchable) {
				manager.Files = append(manager.Files, obj)
			}
		}

		renderChan <- UpdateData{
			manager.RenderList(manager.Files),
			"GOTOP", "",
		}
	}
	close(renderChan)
}

func initManager(path string) *Manager {
	if path == "" {
		path = utils.GetLocalPath()
	}

	var manager Manager
	manager.Path = path
	manager.Highlighting = []int{}
	manager.SetFiles()
	manager.CurrentFileNumber = 0

	return &manager
}
