# Argo OCI Test

Testing ArgoCD with Github OCI.


## Setup
### Setup Kind
Initial
`kind create cluster`
 
Connect kind with kubectl
`kubectl config use-context kind-kind`
 
(Optional) Check context
`kubectl config current-context`
 

### Install ArgoCD
Install acrgo cli via brew:
`brew install argocd`
 
`kubectl create namespace argocd`
`kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`

`kubectl port-forward svc/argocd-server -n argocd 8080:443`

### Connect
Get password
`argocd admin initial-password -n argocd`
 
`argocd login localhost:8080`

Use Username=admin and PW=<from initial-password>