package main

import (
	"context"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

// Workflow GVR for Argo Workflows (argoproj.io/v1alpha1).
// Uses the dynamic client instead of the outdated typed client from
// github.com/argoproj/argo@v2.5.2, which is incompatible with modern
// client-go REST APIs that require context.Context.
var workflowGVR = schema.GroupVersionResource{
	Group:    "argoproj.io",
	Version:  "v1alpha1",
	Resource: "workflows",
}

func main() {
	usr, err := user.Current()
	checkErr(err)

	kubeconfig := flag.String("kubeconfig", filepath.Join(usr.HomeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	checkErr(err)
	namespace := "default"

	dynClient, err := dynamic.NewForConfig(config)
	checkErr(err)
	wfClient := dynClient.Resource(workflowGVR).Namespace(namespace)

	workflow := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "argoproj.io/v1alpha1",
			"kind":       "Workflow",
			"metadata": map[string]interface{}{
				"generateName": "hello-world-",
			},
			"spec": map[string]interface{}{
				"entrypoint": "whalesay",
				"templates": []interface{}{
					map[string]interface{}{
						"name": "whalesay",
						"container": map[string]interface{}{
							"image":   "docker/whalesay:latest",
							"command": []string{"cowsay", "hello world"},
						},
					},
				},
			},
		},
	}

	ctx := context.Background()
	createdWf, err := wfClient.Create(ctx, workflow, metav1.CreateOptions{})
	checkErr(err)
	fmt.Printf("Workflow %s submitted\n", createdWf.GetName())

	timeout := int64(300)
	watcher, err := wfClient.Watch(ctx, metav1.ListOptions{
		FieldSelector:  "metadata.name=" + createdWf.GetName(),
		TimeoutSeconds: &timeout,
	})
	checkErr(err)
	defer watcher.Stop()

	for event := range watcher.ResultChan() {
		if event.Type == watch.Error {
			fmt.Printf("watch error: %v\n", event.Object)
			break
		}
		obj, ok := event.Object.(*unstructured.Unstructured)
		if !ok {
			continue
		}
		finishedAt, found, _ := unstructured.NestedString(obj.Object, "status", "finishedAt")
		phase, _, _ := unstructured.NestedString(obj.Object, "status", "phase")
		if found && finishedAt != "" {
			fmt.Printf("Workflow %s %s at %v\n", obj.GetName(), phase, finishedAt)
			break
		}
		if phase == "Succeeded" || phase == "Failed" || phase == "Error" {
			fmt.Printf("Workflow %s %s\n", obj.GetName(), phase)
			break
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
