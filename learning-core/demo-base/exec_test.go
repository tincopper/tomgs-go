package demo_base

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

func TestSliceSimple(t *testing.T) {
    SliceSimple()
}