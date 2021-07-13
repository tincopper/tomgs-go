package client

import (
    "flag"
    "io"
    "io/ioutil"
    v1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/kubernetes/scheme"
    corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/remotecommand"
    cmdutil "k8s.io/kubectl/pkg/cmd/util"
    "log"
    "os"
    "path/filepath"
    "time"
)

/**
* 测试k8s exec命令
*
* @Author: tangzy
* @Date: 2020/4/9 9:42
 */
func K8sClient() {
    config := GetK8sConfig()
    namespace := "ns-retail-dev"
    podName := "sentinel-dashboard-d6d4dbd6-xqrjz"
    // exec
    coreClient, err := corev1client.NewForConfig(config)
    if err != nil {
        log.Fatalf("NewForConfig err: %v", err)
    }
    
    req := coreClient.RESTClient().Post().
        Resource("pods").
        Name(podName).
        Namespace(namespace).
        SubResource("exec")
    
    req.VersionedParams(&v1.PodExecOptions{
        //Container: pod.,
        Command: []string{"ps", "-ef", "|", "grep", "sentinel-dashboard"},
        Stdin:   true,
        Stdout:  true,
        Stderr:  true,
        TTY:     true,
    }, scheme.ParameterCodec)
    executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
    if err != nil {
        log.Printf("NewSPDYExecutor err: %v", err)
    }
    //reader, writer := io.Pipe()
    // Stream
    err = executor.Stream(remotecommand.StreamOptions{
        Stdin:  os.Stdin,
        Stdout: os.Stdout,
        Stderr: os.Stderr,
        Tty:    false,
    })
    if err != nil {
        log.Fatalf("error %s", err)
    }
    
    time.Sleep(10 * time.Second)
}

func K8sClient2() {
    config := GetK8sConfig()
    namespace := "ns-retail-dev"
    podName := "sentinel-dashboard-555b558fcb-5pm2j"
    
    // create the clientSet
    clientSet, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    req := clientSet.CoreV1().RESTClient().Post().
        Resource("pods").
        Name(podName).
        Namespace(namespace).
        SubResource("exec")
    req.VersionedParams(&v1.PodExecOptions{
        //Container: pod.ContainerName,
        Command:   []string{"ps", "-ef", "|", "grep", "java"},
        Stdin:     true,
        Stdout:    true,
        Stderr:    true,
        TTY:       true,
    }, scheme.ParameterCodec)
    executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
    if err != nil {
        log.Printf("NewSPDYExecutor err: %v", err)
    }
    //reader, writer := io.Pipe()
    // Stream
    err = executor.Stream(remotecommand.StreamOptions{
        Stdin:  os.Stdin,
        Stdout: os.Stdout,
        Stderr: os.Stderr,
        Tty:    false,
    })
    if err != nil {
        log.Fatalf("error %s", err)
    }
    
    time.Sleep(10 * time.Second)
}

func GetK8sClient(restConfig *rest.Config) *kubernetes.Clientset {
    // create the clientSet
    clientSet, err := kubernetes.NewForConfig(restConfig)
    if err != nil {
        panic(err.Error())
    }
    return clientSet
}

func GetK8sConfig() *rest.Config {
    var kubeconfig *string
    if home := homeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, "kube", "configs"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()
    
    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    return config
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}

func Exec(command []string) string {
    config := GetK8sConfig()
    namespace := "ide"
    podName := "theia-bos-maven-release-v1-b699c999d-grlqp"
    // exec
    coreClient, err := corev1client.NewForConfig(config)
    if err != nil {
        log.Panic("NewForConfig err: ", err)
    }
    
    req := coreClient.RESTClient().Post().
        Resource("pods").
        Namespace(namespace).
        Name(podName).
        SubResource("exec")
    
    req.VersionedParams(&v1.PodExecOptions{
        //Container: pod.,
        Command: command,
        Stdin:   true,
        Stdout:  true,
        Stderr:  true,
        TTY: false,
    }, scheme.ParameterCodec)
    executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
    if err != nil {
        log.Panic("NewSPDYExecutor err: ", err)
    }
    //reader, writer := io.Pipe()
    reader, outStream := io.Pipe()
    go func() {
        defer outStream.Close()
        err = executor.Stream(remotecommand.StreamOptions{
            Stdin:  os.Stdin,
            Stdout: outStream,
            Stderr: os.Stderr,
            Tty: false,
        })
        cmdutil.CheckErr(err)
    }()
    
    body, err := ioutil.ReadAll(reader)
    if err != nil {
        log.Panic("ReadAll err: ", err)
    }
    return string(body)
}