package zipping

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func MakeArchive(pathToZip, rootDir string, files []string) error {
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
		if err := AddFileToZip(zipWriter, file, rootDir); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename, rootDir string) error {
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
	fileHeader.Name = strings.Replace(filename, rootDir, "", 1)
	fileHeader.Method = zip.Deflate

	// create zip file
	writer, err := zipWriter.CreateHeader(fileHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
