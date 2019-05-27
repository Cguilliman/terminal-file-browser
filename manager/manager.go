package manager

import (
    "os"
    "strconv"
)

type Manager struct {
    Path              string
    Files, Searchable []os.FileInfo
    CurrentFileNumber int
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

// Return default current directory
// and parent directory `os.FileIngo` objects list
func (manager *Manager) defaultFiles() []os.FileInfo {
    return []os.FileInfo{
        getFile(manager.Path),
        getFile(ParentDirPath(manager.Path)),
    }
}

// Set files in manage.Files list of files objects
func (manager *Manager) SetFiles() {
    base_files := manager.defaultFiles()
    files, err := ioutil.ReadDir(manager.Path)
    if err != nil {
        log.Fatal(err)
    }

    files = append(base_files, files...)
    manager.Files = files
}

// Next file. Change only CurrentFileNumber param.
func (manager *Manager) NextFile() {
    if len(manager.Files)-1 > manager.CurrentFileNumber {
        manager.CurrentFileNumber++
    }
}

// Previous file. Change only CurrentFileNumber param.
func (manager *Manager) PrevFile() {
    if 0 < manager.CurrentFileNumber {
        manager.CurrentFileNumber--
    }
}

// First file. Change only CurrentFileNumber param.
func (manager *Manager) FirstFile() {
    manager.CurrentFileNumber = 0
}

// Last file. Change only CurrentFileNumber param.
func (manager *Manager) LastFile() {
    manager.CurrentFileNumber = len(manager.Files) - 1
}

// Enter directory
// change `manger.Path` as current inner director
// inner directory files and save as `manager.Files`
func (manager *Manager) EnterDir() ([]string, error) {
    file := manager.Files[manager.CurrentFileNumber]
    if !file.IsDir() {
        return errors.New("This is file!")
    }

    switch manager.CurrentFileNumber {
    case 0:
        return manager.RenderList(nil), nil
    case 1:
        manager.Path = ParentDirPath(manager.Path)
    default:
        manager.Path = ConcatPath(manager.Path, file.Name())
    }

    manager.CurrentFileNumber = 0
    manager.SetFiles()
    return manager.RenderList(nil), nil
}

func (manager *Manager) Search(searchChan chan string, renderChan chan []string) {
    manager.Searchable = manager.Files
    for searchable := range searchChan {
        manager.CurrentFileNumber = 0
        manager.Files = manager.defaultFiles()

        for _, obj := range manager.Searchable {
            // TODO: re-factor searching
            if strings.Contains(obj.Name(), searchable) {
                manager.Files = append(manager.Files, obj)
            }
        }

        renderChan <- manager.RenderList(manager.Files)
    }
}
