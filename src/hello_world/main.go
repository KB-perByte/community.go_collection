package main

import (
    "fmt"
    "strings"
    "ansible-go-modules/common"
)


func main() {
    module, err := common.NewAnsibleModule()
    if err != nil {
        fmt.Printf(`{"failed": true, "msg": "Failed to initialize module: %s"}`, err)
        return
    }

    // Get parameters
    name, nameExists := module.GetStringParam("name")
    if !nameExists {
        name = "World"
    }

    greeting, greetingExists := module.GetStringParam("greeting")
    if !greetingExists {
        greeting = "Hello"
    }

    uppercase, _ := module.GetBoolParam("uppercase")

    // Generate message
    message := fmt.Sprintf("%s, %s!", greeting, name)
    if uppercase {
        message = strings.ToUpper(message)
    }

    // Return response
    module.ExitJSON(common.ModuleResponse{
        Changed: false,
        Msg:     message,
        Results: map[string]interface{}{
            "message":   message,
            "name":      name,
            "greeting":  greeting,
            "uppercase": uppercase,
        },
    })
}