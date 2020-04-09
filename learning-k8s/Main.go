package main

import (
    "context"
    "flag"
    "fmt"
    v1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes/scheme"
    "k8s.io/client-go/tools/remotecommand"
    "log"
    "os"
    "path/filepath"
    "time"
    
    "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
    "k8s.io/client-go/tools/clientcmd"
)

/**
* @Author: tangzy
* @Date: 2020/4/8 15:50
 */
func main() {
    var kubeconfig *string
    if home := homeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, "kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()
    
    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    
    // create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    namespace := "ns-retail-dev"
    podName := "sentinel-dashboard-555b558fcb-5pm2j"
    //pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
    // exec
    coreclient, err := corev1client.NewForConfig(config)
    req := coreclient.RESTClient().Post().
        Resource("pods").
        Name(podName).
        Namespace(namespace).
        SubResource("exec")
    
    req.VersionedParams(&v1.PodExecOptions{
        //Container: pod.,
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
    
    //
    /*req := clientset.CoreV1().RESTClient().Post().
        Resource("pods").
        Name(podName).
        Namespace(namespace).
        SubResource("exec")
    req.VersionedParams(&v1.PodExecOptions{
        //Container: pod.,
        Command:   []string{"/bin/sh echo `date`"},
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
    }*/
    // base
    for {
        pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
        
        // Examples for error handling:
        // - Use helper functions like e.g. errors.IsNotFound()
        // - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
        namespace := "ns-retail-dev"
        pod := "sentinel-dashboard-555b558fcb-5pm2j"
        _, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
        if errors.IsNotFound(err) {
            fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
        } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
            fmt.Printf("Error getting pod %s in namespace %s: %v\n",
                pod, namespace, statusError.ErrStatus.Message)
        } else if err != nil {
            panic(err.Error())
        } else {
            fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
        }
        
        time.Sleep(10 * time.Second)
    }
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows
}
