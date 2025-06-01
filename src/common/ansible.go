package common

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// ModuleResponse represents the standard Ansible module response
type ModuleResponse struct {
    Changed bool        `json:"changed"`
    Failed  bool        `json:"failed,omitempty"`
    Msg     string      `json:"msg,omitempty"`
    Results interface{} `json:"results,omitempty"`
    Meta    interface{} `json:"ansible_facts,omitempty"`
}

// AnsibleModule provides utilities for Ansible module development
type AnsibleModule struct {
    Params map[string]interface{}
}

// NewAnsibleModule creates a new Ansible module instance
func NewAnsibleModule() (*AnsibleModule, error) {
    if len(os.Args) < 2 {
        return nil, fmt.Errorf("no arguments file provided")
    }

    argsFile := os.Args[1]
    data, err := ioutil.ReadFile(argsFile)
    if err != nil {
        return nil, fmt.Errorf("failed to read arguments file: %v", err)
    }

    var params map[string]interface{}
    if err := json.Unmarshal(data, &params); err != nil {
        return nil, fmt.Errorf("failed to parse arguments: %v", err)
    }

    return &AnsibleModule{
        Params: params,
    }, nil
}

// GetParam returns a parameter value with type assertion
func (m *AnsibleModule) GetParam(key string) (interface{}, bool) {
    val, exists := m.Params[key]
    return val, exists
}

// GetStringParam returns a string parameter
func (m *AnsibleModule) GetStringParam(key string) (string, bool) {
    if val, exists := m.Params[key]; exists {
        if str, ok := val.(string); ok {
            return str, true
        }
    }
    return "", false
}

// GetBoolParam returns a boolean parameter
func (m *AnsibleModule) GetBoolParam(key string) (bool, bool) {
    if val, exists := m.Params[key]; exists {
        if b, ok := val.(bool); ok {
            return b, true
        }
    }
    return false, false
}

// ExitJSON outputs JSON and exits
func (m *AnsibleModule) ExitJSON(response ModuleResponse) {
    output, _ := json.Marshal(response)
    fmt.Println(string(output))
    if response.Failed {
        os.Exit(1)
    }
    os.Exit(0)
}

// FailJSON outputs failure JSON and exits
func (m *AnsibleModule) FailJSON(msg string) {
    m.ExitJSON(ModuleResponse{
        Failed: true,
        Msg:    msg,
    })
}