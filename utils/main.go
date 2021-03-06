package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetFile(path string) os.FileInfo {
	file, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err.Error)
	}
	return file
}

func RemoveFromIntArray(val int, arr []int) []int {
	for n, item := range arr {
		if item == val {
			lastArr := []int{}
			if len(arr) > n+1 {
				lastArr = arr[n+1:]
			}
			arr = append(arr[:n], lastArr...)
			return arr
		}
	}
	return arr
}

func InIntArray(val int, arr []int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func GetLocalPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func ConcatPath(path, file string) string {
	if string(path[len(path)-1]) != string("/") {
		path = path + "/"
	}
	return path + file
}

func ParentDirPath(path string) string {
	if string(path[len(path)-1]) == string("/") {
		path = path[:len(path)-2]
	}
	dirs := strings.Split(path, "/")
	return strings.Join(dirs[:len(dirs)-1], "/")
}

func GetNested(path string) []string {
	var files []string
	fileObjs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, obj := range fileObjs {
		_path := ConcatPath(path, obj.Name())

		if obj.IsDir() {
			files = append(files, GetNested(_path)...)
		} else {
			files = append(files, _path)
		}
	}

	return files
}
