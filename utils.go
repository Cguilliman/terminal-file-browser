package main

import (
    "fmt"
    "os/exec"
)

func Cd(manager *Manager) {
    cmd := exec.Command("cd", manager.Path)
    // if runtime.GOOS == "windows" {
    //     cmd = exec.Command("tasklist")
    // }
    err := cmd.Run()
    if err != nil {
        fmt.Println("cmd.Run() failed with %s\n", err)
    }
}
