package core

import "strings"

func search(manager *Manager) func(string) []string {
    manager.Searchable = manager.Files
    
    return func(searchable string) []string {
        manager.CurrentFileNumber = 0
        manager.Files = manager.defaultFiles()

        for _, obj := range manager.Searchable[2:] {
            // TODO: re-factor searching
            if strings.Contains(obj.Name(), searchable) {
                manager.Files = append(manager.Files, obj)
            }
        }

        return manager.RenderList(manager.Files)
    }
}
