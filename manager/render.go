package manager

import (
	"os"
	"strconv"
)

func getFileName(n int, file os.FileInfo) string {
	var fileName string
	switch n {
	case 0:
		fileName = ".(Current)"
	case 1:
		fileName = "..(Parent)"
	default:
		fileName = file.Name()
	}
	return fileName
}

func colorRender(name string, isDir, isCurrent bool) string {
	switch {
	case isDir && isCurrent:
		return ">> [ " + name + "](fg:green)"
	case isCurrent:
		return ">> " + name
	case isDir:
		return "[" + name + "](fg:blue)"
	default:
		return name
	}
}

func renderList(manager Manager) []string {
	var response []string

	for n, file := range manager.Files {
		fileName := getFileName(n, file) + " " + strconv.Itoa(int(file.Size()))
		response = append(
			response,
			colorRender(
				fileName,
				file.IsDir(),
				n == manager.CurrentFileNumber,
			),
		)
	}

	return response
}
