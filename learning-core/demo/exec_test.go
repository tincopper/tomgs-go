package demo

import (
    "fmt"
    "log"
    "os/exec"
    "testing"
)

func TestExecCmd(t *testing.T) {
    command := exec.Command("ipconfig")
    output, err := command.Output()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(output))
}