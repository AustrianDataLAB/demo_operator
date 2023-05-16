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