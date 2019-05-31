package temporary

import (
	"fmt"
	"github.com/Cguilliman/terminal-file-browser/utils"
	"os"
	"strings"
)

const (
	COPY string = "copy"
	CUT  string = "cut"
	DEL  string = "delete"
)

type TemporaryFiles struct {
	TemporaryFiles []string
	NewFiles       []string
	BaseDirectory  string
	Action         string
}

func (self *TemporaryFiles) Paste(path string) {
	if self.Action == COPY || self.Action == CUT {
		self.Copy(path)
		if self.Action == CUT {
			self.Delete()
		}
	}
}

func (self *TemporaryFiles) Delete() {
	for _, file := range self.TemporaryFiles {
		os.RemoveAll(file)
	}
}

func (self *TemporaryFiles) Copy(path string) {
	var err error

	for _, filePath := range self.TemporaryFiles {
		oldObj := utils.GetFile(filePath)
		newPath := strings.Replace(
			filePath,
			self.BaseDirectory,
			path, 1,
		)

		if oldObj.IsDir() {
			err = DirCopping(filePath, newPath)
		} else {
			err = FileCopping(filePath, newPath)
		}

		// TODO edit to normal error displaying
		if err != nil {
			fmt.Println(err)
		}
		self.NewFiles = append(self.NewFiles, newPath)
	}
}
