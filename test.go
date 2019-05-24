package main

import (
	"fmt"
	// "os"
	// "path/filepath"
	mg "github.com/Cguilliman/terminal-file-browser/manager"
)

func main() {
	fmt.Println(mg.ParentDirPath("/home/guilliman/go/src/github.com/Cguilliman/terminal-file-browser"))
	fmt.Println(mg.ConcatPath("/home/guilliman/go/src/github.com/Cguilliman/terminal-file-browser", "qqqq"))
	// var files []string

	// root := "/some/folder/to/scan"
	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	files = append(files, path)
	// 	return nil
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
}
