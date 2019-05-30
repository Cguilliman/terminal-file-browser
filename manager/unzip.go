package manager

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func Unzip(outputs string) {
    var filenames []string

    reader, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer reader.Close()

    for _, file := range reader.File {

        // Store filename/path for returning and using later on
        filePath := filepath.Join(dest, file.Name)

        // Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
        if !strings.HasPrefix(filePath, filepath.Clean(outputs)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", filePath)
        }

        filenames = append(filenames, filePath)

        if file.FileInfo().IsDir() {
            // Make Folder
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        // Make File
        if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
            return filenames, err
        }

        outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return filenames, err
        }

        rc, err := file.Open()
        if err != nil {
            return filenames, err
        }

        _, err = io.Copy(outFile, rc)

        // Close the file without defer to close before next iteration of loop
        outFile.Close()
        rc.Close()

        if err != nil {
            return filenames, err
        }
    }
    return filenames, nil
}
