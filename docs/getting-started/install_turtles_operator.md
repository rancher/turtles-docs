---
sidebar_position: 4
---

# Install Rancher Turtles Operator

This section walks through different installation options for the Rancher Turtles Operator.

### Install Rancher Turtles Operator with `Cluster API Operator` as a Helm dependency

A `rancher-turtles` chart repository should be added first:

```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```

To install `Cluster API Operator` as a dependency to the `Rancher Turtles`, a minimum set of additional helm flags should be specified:

```bash
helm install rancher-turtles turtles/rancher-turtles --version <chart-version>
    -n rancher-turtles-system
    --dependency-update
    --create-namespace --wait
    --timeout 180s
```

:::note
- If `cert-manager` is already available in the cluster, you can disable its installation as a Rancher Turtles dependency to avoid conflicts:
`--set cluster-api-operator.cert-manager.enabled=false`
- For a list of Rancher Turtles versions, refer to [Releases page](https://github.com/rancher-sandbox/rancher-turtles/releases).
:::

This is the basic, recommended configuration, which manages the creation of a secret containing the required feature flags (`CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` enabled) in the core provider namespace.

If you need to override the default behavior and use an existing secret (or add custom environment variables), you can pass the secret name and namespace helm flags. In this case, as a user, you are in charge of managing the secret creation and its content, including the minimum required features: `CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` enabled.

```bash
helm install ...
    # Passing secret name and namespace for additional environment variables
    --set cluster-api-operator.cluster-api.configSecret.name=<secret_name>
    --set cluster-api-operator.cluster-api.configSecret.namespace=<secret_namespace>
```

The following is an example of a user-managed secret `cluster-api-operator.cluster-api.configSecret.name=variables`, `cluster-api-operator.cluster-api.configSecret.namespace=default` with `CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` feature flags set and an extra custom variable:

```yaml title="secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: variables
  namespace: default
type: Opaque
stringData:
  CLUSTER_TOPOLOGY: "true"
  EXP_CLUSTER_RESOURCE_SET: "true"
  EXP_MACHINE_POOL: "true"
  CUSTOM_ENV_VAR: "false"
```

Any values passed to `helm` with the `cluster-api-operator` key will be passed along to the `Cluster API Operator` project. A full set of available values for the `Cluster API Operator` can be found in the operator [values.yaml](https://github.com/kubernetes-sigs/cluster-api-operator/blob/main/hack/charts/cluster-api-operator/values.yaml).

Currently the available set of values for the `cluster-api-operator` setup in the `rancher-turtles`:

```yaml
cluster-api-operator:
  enabled: true # indicates if CAPI operator should be installed (default: true)
  cert-manager:
    enabled: true # indicates if cert-manager should be installed (default: true)
  cluster-api:
    enabled: true # indicates if core CAPI controllers should be installed (default: true)
    version: v1.4.6 # version of CAPI to install (default: v1.4.6)
    configSecret:
      name: "" # (provide only if using a user-managed secret) name of the config secret to use for core CAPI controllers, used by the CAPI operator. See [CAPI operator](https://github.com/kubernetes-sigs/cluster-api-operator/tree/main/docs#installing-azure-infrastructure-provider) docs for more details.
      namespace: "" # (provide only if using a user-managed secret) namespace of the config secret to use for core CAPI controllers, used by the CAPI operator.
      defaultName: "capi-env-variables" # default name for the automatically created secret.
    core:
      namespace: capi-system
      fetchConfig: # (only required for airgapped environments)
        url: ""  # url to fetch config from, used by the CAPI operator. See [CAPI operator](https://github.com/kubernetes-sigs/cluster-api-operator/tree/main/docs#provider-spec) docs for more details.
        selector: ""  # selector to use for fetching config, used by the CAPI operator.
    kubeadmBootstrap:
      namespace: capi-kubeadm-bootstrap-system
      fetchConfig:
        url: ""
        selector: ""
    kubeadmControlPlane:
      namespace: capi-kubeadm-control-plane-system
      fetchConfig:
        url: ""
        selector: ""
```

### Install Rancher Turtles Operator without `Cluster API Operator` as a Helm dependency

A `Rancher Turtles` requires a connection to the `Rancher Manager` cluster. This can be achieved by:

1. Installing it in the same cluster as the `Rancher Manager`.

:::tip
For information on deployment options, refer to [Deployment Scenarios](../reference-guides/architecture/deployment)
:::

The recommended path of installation for the operator is by using `Helm`. To install it in the cluster, a chart repository should be added first:

```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```
and then it can be installed into the `rancher-turtles-system` namespace with:

```bash
helm install rancher-turtles turtles/rancher-turtles --version <chart-version>
    -n rancher-turtles-system
    --set cluster-api-operator.enabled=false
    --set cluster-api-operator.cluster-api.enabled=false
    --create-namespace --wait
    --dependency-update
```
