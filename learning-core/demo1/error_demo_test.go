package demo1

import "testing"

func TestDiv(t *testing.T) {
    _, err := Div(1, 0)
    if err != nil {
        t.Error(err)
    }
}

func TestPanic(t *testing.T) {
    Panic()
}

func TestRecoverMock(t *testing.T) {
    RecoverMock()
}