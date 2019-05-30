package manager

import (
	"os/exec"
	"strings"
)

// run command on the system
func runCommand(command string, args ...string) {
	// NOTE: is not tracking output, only execute
	cmd := exec.Command(command, args...)
	cmd.Run()
}

// convert string in command
func convert(value, path string) (command string, args []string) {
	REPLACED := "!!"
	if strings.Index(value, REPLACED) > 0 {
		value = strings.Replace(value, REPLACED, path, -1)
	} else {
		// remove space char at the end of the string, if exists
		value = value[:len(value)-1] + strings.Replace(
			value[len(value)-1:], " ", "", 1,
		)
		value = value + " " + path
	}
	splited := strings.Split(value, " ")
	return splited[0], splited[1:]
}

// check is command valid
func check(command string, args ...string) bool {
	// NOT IMPLEMENTED
	return true
}

// listen channel, when it close run command
func Run(runChan chan string, manager *Manager) {
	var (
		command string
		args    []string
	)

	for getted := range runChan {
		_command, _args := convert(getted, manager.GetDirPath(-1))
		if check(command, args...) {
			command = _command
			args = _args
		}
	}
	runCommand(command, args...)
}
