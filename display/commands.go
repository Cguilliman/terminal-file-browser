package display

import (
    "os/exec"
    "strings"
    mg "github.com/Cguilliman/terminal-file-browser/manager"
)

func CreateDir(nameChan chan string, content *mg.ContentList) {
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

    cmd := exec.Command("mkdir", path)
    cmd.Run()
    content.Reset()
}