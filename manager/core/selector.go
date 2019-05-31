package core

import "github.com/Cguilliman/terminal-file-browser/utils"

func pickOutFile(manager *Manager) {
	if manager.CurrentFileNumber > 1 {
		if utils.InIntArray(manager.CurrentFileNumber, manager.Highlighting) {
			manager.Highlighting = utils.RemoveFromIntArray(
				manager.CurrentFileNumber,
				manager.Highlighting,
			)
		} else {
			manager.Highlighting = append(
				manager.Highlighting,
				manager.CurrentFileNumber,
			)
		}
	}
}

func pickOutAllFiles(manager *Manager) {
	allFiles := []int{}
	for i := 2; i != len(manager.Files); i++ {
		allFiles = append(allFiles, i)
	}
	if len(manager.Highlighting) == 0 && len(allFiles) > 0 {
		manager.Highlighting = allFiles
	} else {
		manager.Highlighting = []int{}
	}
}
