# TL;DR

## locally 
open VSCode
```
brew install kubebuilder
```
(or equivalent linux)

```
kubebuilder init --domain operator.caas-0002.beta.austrianopencloudcommunity.org --repo demooperator
```
```
kubebuilder create api --group demooperator --version v1 --kind LvaVolume
```

Now, go to api/v1 and modify the lvavooume_types.go with  some very basic stuff like name and type.
```
make generate
make manifests
```

go to internal/controller/lvavolume_controller.go and write custom logic

make sure you are connected to a kubernetes cluster and type `make install` this will install this amazing operator
(which curretnly prints something to the logs and thats it)

```
/Users/croedig/gitrepos/demooperator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
test -s /Users/croedig/gitrepos/demooperator/bin/kustomize || { curl -Ss "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" --output install_kustomize.sh && bash install_kustomize.sh 5.0.0 /Users/croedig/gitrepos/demooperator/bin; rm install_kustomize.sh; }
v5.0.0
kustomize installed to /Users/croedig/gitrepos/demooperator/bin/kustomize
/Users/croedig/gitrepos/demooperator/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/lvavolumes.demooperator.operator.caas-0002.beta.austrianopencloudcommunity.org created
```

check that they got created
```
kubectl get crd
```

edit the spec in the sample yaml config/samples/demooperator_v1_lvavolume.yaml and apply it
