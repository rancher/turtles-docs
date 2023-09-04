---
sidebar_position: 4
---

# Install Rancher Turtles Operator

A `Rancher Turtles` requires a connection to the `Rancher Manager` cluster. This can be achieved by:

1. Installing it in the same cluster as the `Rancher Manager`.

*Note: In future we will support different deployment topologies*

Recommended path of installation for the operator is by using `Helm`. To install it in the cluster, a repository should be added first:

```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```
and then it can be installed into the `rancher-turtles-system` namespace with:

```bash
helm install rancher-turtles turtles/rancher-turtles
    -n rancher-turtles-system
    --create-namespace --wait
```

### Install `Cluster API Operator` as a Helm dependency

*Note: this section will be extended with additional details later*

To install `Cluster API Operator` as a dependency to the `Rancher Turtles`, a set of additional helm flags should be specified:

```bash
helm install rancher-turtles turtles/rancher-turtles
    -n rancher-turtles-system
    --set cluster-api-operator.cert-manager.enabled=true # Allows to install cert manager dependency with Cluster API Operator
	--set secret-name=<secret_name> # Passing secret name and namespace for additional environment variables to be used when deploying CAPI provider
	--set secret-namespace=<secret_namespace>
    --create-namespace --wait
    --timeout 90s
```

Any values passed to `helm` with the `cluster-api-operator` label will be passed along to the `Cluster API Operator` project. Full set of avaliable values for the `Cluster API Operator` could be found in the operator [values.yaml](https://github.com/kubernetes-sigs/cluster-api-operator/blob/main/hack/charts/cluster-api-operator/values.yaml).
