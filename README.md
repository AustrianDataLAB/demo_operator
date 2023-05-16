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