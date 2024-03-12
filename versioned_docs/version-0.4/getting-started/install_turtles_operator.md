---
sidebar_position: 4
---

# Install Rancher Turtles Operator

This section walks through different installation options for the Rancher Turtles Operator.

:::info
Before [installing Rancher Turtles](#install-rancher-turtles-operator-with-cluster-api-operator-as-a-helm-dependency) in your Rancher environment, Rancher's `embedded-cluster-api` functionality must be disabled. This includes also cleaning up Rancher-specific webhooks that otherwise would conflict with CAPI ones.

To simplify setting up Rancher for installing Rancher Turtles, the official Rancher Turtles Helm chart includes a `pre-install` hook that applies these changes, making it transparent to the end user:
- Disable the `embedded-cluster-api` feature in Rancher.
- Delete the `mutating-webhook-configuration` and `validating-webhook-configuration` webhooks that are no longer needed.
:::

### Install Rancher Turtles Operator with `Cluster API Operator` as a Helm dependency

A `rancher-turtles` chart repository should be added first:

```bash
helm repo add turtles https://charts.rancher-turtles.com/
helm repo update
```

To install `Cluster API Operator` as a dependency to the `Rancher Turtles`, a minimum set of additional helm flags should be specified:

```bash
helm install rancher-turtles turtles/rancher-turtles --version v0.4.0 \
    -n rancher-turtles-system \
    --dependency-update \
    --create-namespace --wait \
    --timeout 180s
```

:::note
- If `cert-manager` is already available in the cluster, you can disable its installation as a Rancher Turtles dependency to avoid conflicts:
`--set cluster-api-operator.cert-manager.enabled=false`
- For a list of Rancher Turtles versions, refer to [Releases page](https://github.com/rancher-sandbox/rancher-turtles/releases).
:::

This is the basic, recommended configuration, which manages the creation of a secret containing the required feature flags (`CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` enabled) in the core provider namespace.

If you need to override the default behavior and use an existing secret (or add custom environment variables), you can pass the secret name helm flag. In this case, as a user, you are in charge of managing the secret creation and its content, including the minimum required features: `CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` enabled.

```bash
helm install ...
    # Passing secret name and namespace for additional environment variables
    --set cluster-api-operator.cluster-api.configSecret.name=<secret_name>
```

The following is an example of a user-managed secret `cluster-api-operator.cluster-api.configSecret.name=variables` with `CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` feature flags set and an extra custom variable:

```yaml title="secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: variables
  namespace: rancher-turtles-system
type: Opaque
stringData:
  CLUSTER_TOPOLOGY: "true"
  EXP_CLUSTER_RESOURCE_SET: "true"
  EXP_MACHINE_POOL: "true"
  CUSTOM_ENV_VAR: "false"
```

:::info
For detailed information on the values supported by the chart and their usage, refer to [Helm chart options](../reference-guides/rancher-turtles-chart/values)
:::

### Install Rancher Turtles Operator without `Cluster API Operator` as a Helm dependency

A `Rancher Turtles` requires a connection to the `Rancher Manager` cluster. This can be achieved by:

1. Installing it in the same cluster as the `Rancher Manager`.

:::tip
For information on deployment options, refer to [Deployment Scenarios](../reference-guides/architecture/deployment)
:::

The recommended path of installation for the operator is by using `Helm`. To install it in the cluster, a chart repository should be added first:

```bash
helm repo add turtles https://charts.rancher-turtles.com/
helm repo update
```
and then it can be installed into the `rancher-turtles-system` namespace with:

```bash
helm install rancher-turtles turtles/rancher-turtles --version v0.4.0
    -n rancher-turtles-system
    --set cluster-api-operator.enabled=false
    --set cluster-api-operator.cluster-api.enabled=false
    --create-namespace --wait
    --dependency-update
```
