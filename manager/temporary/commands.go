package temporary

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func DirCopping(oldPath, newPath string) error {
	var (
		err      error
		fileObjs []os.FileInfo
		info     os.FileInfo
	)

	if info, err = os.Stat(oldPath); err != nil {
		return err
	}
	if err = os.MkdirAll(newPath, info.Mode()); err != nil {
		return err
	}

	if fileObjs, err = ioutil.ReadDir(oldPath); err != nil {
		return err
	}
	for _, obj := range fileObjs {
		oldObjPath := path.Join(oldPath, obj.Name())
		newObjPath := path.Join(newPath, obj.Name())

		if obj.IsDir() {
			if err := DirCopping(oldObjPath, newObjPath); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := FileCopping(oldObjPath, newObjPath); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func FileCopping(oldPath, newPath string) error {
	var (
		err     error
		file    *os.File
		newFile *os.File
		oldInfo os.FileInfo
	)
	if file, err = os.Open(oldPath); err != nil {
		return err
	}
	defer file.Close()

	if newFile, err = os.Create(newPath); err != nil {
		return err
	}
	defer newFile.Close()

	if _, err = io.Copy(newFile, file); err != nil {
		return err
	}
	if oldInfo, err = os.Stat(oldPath); err != nil {
		return err
	}
	return os.Chmod(newPath, oldInfo.Mode())
}
