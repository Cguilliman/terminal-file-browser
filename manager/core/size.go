package core

import (
    "os"
    "github.com/Cguilliman/terminal-file-browser/utils"
)

func getSize(basePath string, dir os.FileInfo, n int, resChan chan [2]int) {
    size := 0
    files := utils.GetNested(
        utils.ConcatPath(basePath, dir.Name()),
    )
    for _, file := range files {
        size = size + int(utils.GetFile(file).Size())
    }
    resChan <- [2]int{n, size}
}

func sizeCalc(manager *Manager, updateChan chan bool) {
    resChan := make(chan [2]int)
    if manager.sizes == nil {
        manager.sizes = make(map[int]int)
    }

    for n, obj := range manager.Files {
        if obj.IsDir() {
            go getSize(manager.Path, obj, n, resChan)
        }
    }

    for res := range resChan {
        manager.sizes = map[int]int{res[0]: res[1]}
        updateChan <- true 
    }
}
