package zipping

import (
	"fmt"
	mg "github.com/Cguilliman/terminal-file-browser/manager"
	"github.com/Cguilliman/terminal-file-browser/utils"
)

func Unzipping(channel chan string, content *mg.ContentList) {
	var filename string
	for _filename := range channel {
		filename = _filename
	}

	Unzip(
		content.Manager.GetDirPath(-1),
		utils.ConcatPath(content.Manager.Path, filename),
	)
	// fmt.Println(s, err)
	content.Reset(true)
}

func Zipping(zipChan chan string, content *mg.ContentList) {
	var (
		filePath string
		files    []string
	)
	for value := range zipChan {
		filePath = value
	}

	for _, fileObjPath := range content.GetSelectedFiles() {
		if utils.GetFile(fileObjPath).IsDir() {
			files = append(files, utils.GetNested(fileObjPath)...)
		} else {
			files = append(files, fileObjPath)
		}
	}
	// TODO print errors in some GUI block
	err := MakeArchive(
		utils.ConcatPath(content.Manager.Path, filePath),
		content.Manager.Path,
		files,
	)
	if err != nil {
		fmt.Println(err)
	}
	content.Reset(true)
}
