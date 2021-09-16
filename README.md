# Tekton Demo

Demo of linting, testing, building and running a [Hello World Go app](hello.go)
using [Tekton][tekton].

[Tekton][tekton] is a powerful yet flexible Kubernetes-native open-source framework for creating continuous integration
and delivery (CI/CD) systems. It lets users build, test, and deploy across multiple cloud providers or on-premises
systems by abstracting away the underlying implementation details.

## Setup

    k get pods -n tekton-pipelines


## Usage

      # Install Tasks from the Catalog
      tkn hub install task git-clone
      tkn hub install task golang-test
      tkn hub install task golangci-lint
      tkn hub install task golang-build

      # Install Golang Run Task 
      kubectl apply -f tekton/golang-run.yaml

      # Install the Pipeline
      kubectl apply -f tekton/pipeline.yaml

      # Make the Persistent Volume Claim that we'll be using 
      kubectl apply -f tekton/workspace.yaml
        
      # Execute the Pipeline by creating a PipelineRun
      tkn pipeline start --showlog \
        -f tekton/pipeline.yaml \
        -p package=github.com/jerop/tekton-demo \
        -w name=workarea,volumeClaimTemplateFile=tekton/workspace.yaml

      tkn pipelinerun ls
      tkn pipelinerun describe -L
      tkn taskrun describe <taskRunName>
      kubectl describe tr <taskRunName>
      kubectl get -o yaml pod <podName> | less

[tekton]: https://tekton.dev/