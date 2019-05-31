package manager

import (
    "strings"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path"
)

const (
    COPY string = "copy"
    CUT  string = "cut"
    DEL  string = "delete"
)

type TemporaryFiles struct {
    TemporaryFiles []string
    NewFiles       []string
    baseDirectory  string
    action         string
}

func (self *TemporaryFiles) Paste(path string) {
    if self.action == COPY || self.action == CUT {
        self.Copy(path)
        if self.action == CUT {
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
        oldObj := getFile(filePath)
        newPath := strings.Replace(
            filePath,
            self.baseDirectory,
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

func DirCopping(oldPath, newPath string) error {
    var (
        err error
        fileObjs []os.FileInfo
        info os.FileInfo
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
        err error
        file *os.File
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
