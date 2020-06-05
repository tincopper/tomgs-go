package main

import (
    "gopkg.in/yaml.v2"
    "log"
)
var data = `
# This example demonstrates the use of a git repo as a hard-wired input artifact.
# The argo repo is cloned to its target destination at '/src' for the main container to consume.
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: v7-auto-test-demo1-
  namespace: v7-auto-test
spec:
  serviceAccountName: sa-v7-auto-test
  imagePullSecrets:
  - name: regsecret
  entrypoint: autotest-dag
  templates:
  - name: autotest-dag
  - name: autotest-from-git
    inputs:
      parameters:
      - name: casename
      - name: reportdir
      artifacts:
      - name: argo-source
        path: /src
        git:
          repo: http://git.kingdee.com/autotest/V7_AutoTest.git
          revision: "master"
          usernameSecret:
            name: autotest
            key: username
          passwordSecret:
            name: autotest
            key: password
          depth: 1
    container:
      workingDir: /src/
      image: reg.jdy.com/auto-test/v7_autotest:alpine
      command: [sh, -xc]
      args:
      - |
        REMOTE_URL=http://selenium-hub:4444/wd/hub &&
        ACCOUNT_NAME=自动化测试-docker &&
        pytest {{inputs.parameters.casename}} \
        -c pytest_utf8.ini \
        -m "level_1" \
        -s \
        --remote_url=$REMOTE_URL \
        --browser_name=chrome \
        --html={{inputs.parameters.reportdir}} \
        --self-contained-html \
        --account_name=$ACCOUNT_NAME \
        --env="PRODUCTION"
`

/**
* @Author: tangzy
* @Date: 2020/5/16 18:05
 */
func main() {
    m := make(map[interface{}]interface{})
    m["a"] = 1
    m["b"] = 2
    
    result, err := yaml.Marshal(&m)
    if err != nil {
        log.Fatalf("fata: %v", err)
    }
    log.Printf("\n%s\n", string(result))
    
    t := make(map[interface{}]interface{})
    err = yaml.Unmarshal([]byte(data), &t)
    if err != nil {
        log.Fatalf("fata: %v", err)
    }
    
    log.Printf("\n%s\n", t)
    
    spec := t["spec"]
    // 将interface类型强制转换为map类型
    //specMap := spec.(map[string]interface{})
    //specMap[""]
    
    log.Println(spec)
    
    result1, err := yaml.Marshal(&t)
    if err != nil {
        log.Fatalf("fata: %v", err)
    }
    log.Printf("\n%s\n", string(result1))
}
