package client

import (
    "archive/tar"
    "fmt"
    "io"
    corev1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes/scheme"
    "k8s.io/client-go/tools/remotecommand"
    "path/filepath"
    "strings"
    
    //cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
    cmdutil "k8s.io/kubectl/pkg/cmd/util"
    "log"
    "os"
    "path"
)

type Pod struct {
    Namespace string
    Name string
    ContainerName string
}

func (p *Pod) CopyFromPod(srcPath string, destPath string) error {
    restConfig := GetK8sConfig()
    client := GetK8sClient(restConfig)
    
    reader, outStream := io.Pipe()
    //todo some containers failed : tar: Refusing to write archive contents to terminal (missing -f option?) when execute `tar cf -` in container
    cmdArr := []string{"tar", "cf", "-", srcPath}
    req := client.CoreV1().RESTClient().
        Get().
        Namespace(p.Namespace).
        Resource("pods").
        Name(p.Name).
        SubResource("exec").
        VersionedParams(&corev1.PodExecOptions{
            Container: p.ContainerName,
            Command:   cmdArr,
            Stdin:     true,
            Stdout:    true,
            Stderr:    true,
            TTY:       false,
        }, scheme.ParameterCodec)
    
    exec, err := remotecommand.NewSPDYExecutor(restConfig, "GET", req.URL())
    if err != nil {
        log.Fatalf("error %s\n", err)
        return err
    }
    go func() {
        defer outStream.Close()
        err = exec.Stream(remotecommand.StreamOptions{
            Stdin:  os.Stdin,
            Stdout: outStream,
            Stderr: os.Stderr,
            Tty:    false,
        })
        cmdutil.CheckErr(err)
    }()
    prefix := getPrefix(srcPath)
    prefix = path.Clean(prefix)
    prefix = stripPathShortcuts(prefix)
    destPath = path.Join(destPath, path.Base(prefix))
    err = untarAll(reader, destPath, prefix)
    return err
}

func untarAll(reader io.Reader, destDir, prefix string) error {
    tarReader := tar.NewReader(reader)
    for {
        header, err := tarReader.Next()
        if err != nil {
            if err != io.EOF {
                return err
            }
            break
        }
        
        if !strings.HasPrefix(header.Name, prefix) {
            return fmt.Errorf("tar contents corrupted")
        }
        
        mode := header.FileInfo().Mode()
        destFileName := filepath.Join(destDir, header.Name[len(prefix):])
        
        baseName := filepath.Dir(destFileName)
        if err := os.MkdirAll(baseName, 0755); err != nil {
            return err
        }
        if header.FileInfo().IsDir() {
            if err := os.MkdirAll(destFileName, 0755); err != nil {
                return err
            }
            continue
        }
        
        evaledPath, err := filepath.EvalSymlinks(baseName)
        if err != nil {
            return err
        }
        
        if mode&os.ModeSymlink != 0 {
            linkname := header.Linkname
            
            if !filepath.IsAbs(linkname) {
                _ = filepath.Join(evaledPath, linkname)
            }
            
            if err := os.Symlink(linkname, destFileName); err != nil {
                return err
            }
        } else {
            outFile, err := os.Create(destFileName)
            if err != nil {
                return err
            }
            defer outFile.Close()
            if _, err := io.Copy(outFile, tarReader); err != nil {
                return err
            }
            if err := outFile.Close(); err != nil {
                return err
            }
        }
    }
    
    return nil
}

func getPrefix(file string) string {
    return strings.TrimLeft(file, "/")
}

func stripPathShortcuts(p string) string {
    newPath := path.Clean(p)
    trimmed := strings.TrimPrefix(newPath, "../")
    
    for trimmed != newPath {
        newPath = trimmed
        trimmed = strings.TrimPrefix(newPath, "../")
    }
    
    // trim leftover {".", ".."}
    if newPath == "." || newPath == ".." {
        newPath = ""
    }
    
    if len(newPath) > 0 && string(newPath[0]) == "/" {
        return newPath[1:]
    }
    
    return newPath
}