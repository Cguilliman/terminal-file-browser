package manager

import (
    "os/exec"
    "strings"
)

// run command on the system
func runCommand(command string, args ...string) {
    if value != "" {
        // NOTE: is not tracking output, only execute
        exec.Command(command, args...)
    }
}

// convert string in command 
func convert(value, path string) (command string, args ...string) {
    REPLACED := "!!"
    if strings.Index(value, REPLACED) > 0 {
        value := strings.ReplaceAll(value, REPLACED, path)
    } else {

    }
    splited := strings.Split(value, " ")
    return splited[0], splited[1:]
}

// check is command valid
func check(command string, args ...args) bool {
    // NOT IMPLEMENTED
    return true
}

// listen channel, when it close run command
func Run(manager Manager, runChan chan string) {
    value := ""
    defer runCommand(value)

    for getted := range runChan {
        command, args := convert(getted, manager.GetDirPath(-1))
        if check(command, args...) {
            value := convert(command, args...)
        }
    }
}