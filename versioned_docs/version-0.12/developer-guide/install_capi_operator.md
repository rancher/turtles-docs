---
sidebar_position: 2
---

# Installing Cluster API Operator

:::caution
Installing Cluster API Operator by following this page (without it being a Helm dependency to Rancher Turtles) is not a recommended installation method and intended only for local development purposes. 
:::

This section describes how to install `Cluster API Operator` in the Kubernetes cluster.

## Installing Cluster API (CAPI) and providers

`CAPI` and desired `CAPI` providers could be installed using the helm-based installation for [`Cluster API Operator`](https://github.com/kubernetes-sigs/cluster-api-operator) or as a helm dependency for the `Rancher Turtles`.

### Install manually with Helm (alternative)
To install `Cluster API Operator` with version `1.7.3` of the `CAPI` + `Docker` provider using helm, follow these steps:

1. Add the Helm repository for the `Cluster API Operator`:
```bash
helm repo add capi-operator https://kubernetes-sigs.github.io/cluster-api-operator
helm repo add jetstack https://charts.jetstack.io

```
2. Update the Helm repository:
```bash
helm repo update
```
3. Install the Cert-Manager:
```bash
helm install cert-manager jetstack/cert-manager --namespace cert-manager --create-namespace --set installCRDs=true
```
4. Install the `Cluster API Operator`:
```bash
helm install capi-operator capi-operator/cluster-api-operator
	--create-namespace -n capi-operator-system
	--set infrastructure=docker:v1.7.3
	--set core=cluster-api:v1.7.3
	--timeout 90s --wait # Core Cluster API with kubeadm bootstrap and control plane providers will also be installed
```

:::note
`cert-manager` is a hard requirement for `CAPI` and `Cluster API Operator`*
:::

To provide additional environment variables, enable feature gates, or supply cloud credentials, similar to `clusterctl` [common provider](https://cluster-api.sigs.k8s.io/user/quick-start#initialization-for-common-providers), variables secret with `name` and a `namespace` of the secret could be specified for the `Cluster API Operator` as shown below.

```bash
helm install capi-operator capi-operator/cluster-api-operator
	--create-namespace -n capi-operator-system
	--set infrastructure=docker:v1.7.3
	--set core=cluster-api:v1.7.3
	--timeout 90s
	--secret-name <secret_name>
	--wait
```

Example secret data:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: variables
  namespace: default
type: Opaque
stringData:
  CLUSTER_TOPOLOGY: "true"
  EXP_CLUSTER_RESOURCE_SET: "true"
```

To select more than one desired provider to be installed together with the `Cluster API Operator`, the `--infrastructure` flag can be specified with multiple provider names separated by a comma. For example:

```bash
helm install ... --set infrastructure="docker:v1.7.3;aws:v2.6.1"
```

The `infrastructure` flag is set to `docker:v1.7.3;aws:v2.6.1`, representing the desired provider names. This means that the `Cluster API Operator` will install and manage multiple providers, `Docker` and `AWS`, with versions `v1.7.3` and `v2.6.1` respectively.

The cluster is now ready to install Rancher Turtles. The default behavior when installing the chart is to install Cluster API Operator as a Helm dependency. Since we decided to install it manually before installing Rancher Turtles, the feature `cluster-api-operator.enabled` must be explicitly disabled as otherwise it would conflict with the existing installation. You can refer to [Install Rancher Turtles without Cluster API Operator](../developer-guide/install_capi_operator.md#install-rancher-turtles-without-cluster-api-operator-as-a-helm-dependency) to see next steps.

:::tip
For more fine-grained control of the providers and other components installed with CAPI, see the [Add the infrastructure provider](../tasks/capi-operator/add_infrastructure_provider.md) section.
:::
''

### Install Rancher Turtles without `Cluster API Operator` as a Helm dependency

:::note
This option is only suitable for development purposes and not recommended for production environments.
:::

The `rancher-turtles` chart is available in https://rancher.github.io/turtles and this Helm repository must be added before proceeding with the installation:

```bash
helm repo add turtles https://rancher.github.io/turtles
helm repo update
```

and then it can be installed into the `rancher-turtles-system` namespace with:

```bash
helm install rancher-turtles turtles/rancher-turtles --version v0.12.0
    -n rancher-turtles-system
    --set cluster-api-operator.enabled=false
    --set cluster-api-operator.cluster-api.enabled=false
    --create-namespace --wait
    --dependency-update
```

As you can see, we are telling Helm to ignore installing `cluster-api-operator` as a dependency.

