package main

import (
	"fmt"
    "os"
	"ansible-go-modules/common"
)


func main() {
    module, err := common.NewAnsibleModule()
    if err != nil {
        module.FailJSON(fmt.Sprintf("Failed to initialize module: %s", err))
        return
    }

    // Get parameters
    path, pathExists := module.GetStringParam("path")
    if !pathExists {
        module.FailJSON("path parameter is required")
        return
    }

    state, stateExists := module.GetStringParam("state")
    if !stateExists {
        state = "file"
    }

    var changed bool
    var message string

    switch state {
    case "directory":
        if err := os.MkdirAll(path, 0755); err != nil {
            module.FailJSON(fmt.Sprintf("Failed to create directory: %s", err))
            return
        }
        changed = true
        message = fmt.Sprintf("Directory %s created", path)

    case "absent":
        if _, err := os.Stat(path); err == nil {
            if err := os.RemoveAll(path); err != nil {
                module.FailJSON(fmt.Sprintf("Failed to remove path: %s", err))
                return
            }
            changed = true
            message = fmt.Sprintf("Path %s removed", path)
        } else {
            message = fmt.Sprintf("Path %s does not exist", path)
        }

    case "touch":
        file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            module.FailJSON(fmt.Sprintf("Failed to touch file: %s", err))
            return
        }
        file.Close()
        changed = true
        message = fmt.Sprintf("File %s touched", path)

    default:
        // Check if file exists
        if _, err := os.Stat(path); os.IsNotExist(err) {
            module.FailJSON(fmt.Sprintf("File %s does not exist", path))
            return
        }
        message = fmt.Sprintf("File %s exists", path)
    }

    // Get file stats
    var fileInfo map[string]interface{}
    if stat, err := os.Stat(path); err == nil {
        fileInfo = map[string]interface{}{
            "path":    path,
            "size":    stat.Size(),
            "mode":    stat.Mode().String(),
            "is_dir":  stat.IsDir(),
            "mod_time": stat.ModTime().Unix(),
        }
    }

    module.ExitJSON(common.ModuleResponse{
        Changed: changed,
        Msg:     message,
        Results: fileInfo,
    })
}