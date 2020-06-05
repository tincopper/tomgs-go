package client

import "testing"

func TestK8sClient(t *testing.T) {
    K8sClient()
}

func TestK8sClient2(t *testing.T) {
    K8sClient2()
}

func TestExec(t *testing.T) {
    // 命令不能这样传，要分割开
    //Exec("ps -ef | grep sentinel-dashboard")
    //command := []string{"ps", "-ef", "|", "grep", "-v", "'grep'", "|" , "grep", "sentinel-dashboard"}
    // go里面不能直接这样写，管道符会失效
    //https://www.linuxidc.com/Linux/2013-09/90749.htm
    //command := []string{"ps", "-ef", "|", "grep", "sentinel-dashboard"}
    // 应该这样处理，或者使用上面连接的处理方式
    command := []string{"sh", "-c", "ps -ef | grep -v grep | grep sentinel-dashboard"}
    Exec(command)
}