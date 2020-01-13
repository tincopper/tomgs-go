package main

import (
    "log"
    "runtime/debug"
    "testing"
)

func TestUseChannel(t *testing.T) {
    defer func() {
        if err := recover(); err != nil {
            log.Println(string(debug.Stack()))
            return
        }
    }()
    UseChannel()
}

func TestUseChannel2(t *testing.T) {
    UseChannel2()
}