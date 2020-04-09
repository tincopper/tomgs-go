package client

import (
    "log"
    "testing"
)

func TestPod_CopyFromPod(t *testing.T) {
    namespace := "ns-retail-dev"
    podName := "sentinel-dashboard-555b558fcb-5pm2j"
    p := &Pod{
        Namespace: namespace,
        Name:      podName,
    }
    err := p.CopyFromPod("home/dashboard.hprof", "F:/")
    if err != nil {
        log.Fatal(err)
    }
}