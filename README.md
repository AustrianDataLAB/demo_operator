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
make
```

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

then you can run `make run` and observe that the amazing logs are being produced
```
2023-05-16T10:51:18+02:00       INFO    You are executing the reconcile loop right now  {"controller": "lvavolume", "controllerGroup": "demooperator.operator.caas-0002.beta.austrianopencloudcommunity.org", "controllerKind": "LvaVolume", "LvaVolume": {"name":"lvavolume-sample","namespace":"default"}, "namespace": "default", "name": "lvavolume-sample", "reconcileID": "057d3a41-0fee-49d0-aca7-8fe2f31834ff", "req": "default/lvavolume-sample"}
```
