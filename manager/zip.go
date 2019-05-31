package manager

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func Zipping(zipChan chan string, content *ContentList) {
	var (
		filePath string
		files    []string
	)
	for value := range zipChan {
		filePath = value
	}

	// TODO add several files zipping
	var (
		targetPath string      = content.Manager.GetDirPath(-1)
		target     os.FileInfo = getFile(targetPath)
	)
	if target.IsDir() {
		files = GetNested(targetPath)
	} else {
		files = []string{targetPath}
	}
	for n, path := range files {
		files[n] = strings.Replace(path, content.Manager.Path+"/", "", 1)
	}
	// TODO print errors in some GUI block
	if err := MakeArchive(filePath, files); err != nil {
		fmt.Println(err)
	}
	content.Reset(true)
}

func MakeArchive(pathToZip string, files []string) error {
	// create zip file
	zipFile, err := os.Create(pathToZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// create zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// iterate files and move it in zip
	for _, file := range files {
		if err := AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	// open file
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// get file stat
	fileStat, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	// initialize file header
	fileHeader, err := zip.FileInfoHeader(fileStat)
	if err != nil {
		return err
	}

	// initialize file stuff to zipped file
	fileHeader.Name = filename
	fileHeader.Method = zip.Deflate

	// create zip file
	writer, err := zipWriter.CreateHeader(fileHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
