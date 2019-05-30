package display

import (
    "os/exec"
    "strings"
    mg "github.com/Cguilliman/terminal-file-browser/manager"
)

const (
    TOUCH string = "touch"
    MKDIR string = "mkdir"
)

func SimpleRun(command string, nameChan chan string, content *mg.ContentList) {
    var path string

    for value := range nameChan {
        sep := ""
        if string(content.Manager.Path[len(content.Manager.Path)-1]) != "/" {
            sep = "/"
        }
        path = strings.Join([]string{
            content.Manager.Path,
            value,
        }, sep)
    }

    cmd := exec.Command(command, path)
    cmd.Run()
    content.Reset()
}

func CreateDir(nameChan chan string, content *mg.ContentList) {
    SimpleRun(MKDIR, nameChan, content)
}

func CreateFile(nameChan chan string, content *mg.ContentList) {
    SimpleRun(TOUCH, nameChan, content)
}
