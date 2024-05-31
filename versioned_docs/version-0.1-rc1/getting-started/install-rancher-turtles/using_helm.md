---
sidebar_position: 4
---

# Via Helm install

:::caution
In case you need to review the list of prerequisites (including `cluster-api-operator` and `cert-manager`), you can refer to [this table](../intro.md#prerequisites).
:::

If you want to manually apply the Helm chart and be in full control of the installation.

This section walks through different installation options for the Rancher Turtles.
- [Install Rancher Turtles with Cluster API Operator as a dependency](#install-rancher-turtles-with-cluster-api-operator-as-a-helm-dependency).
- [Install Rancher Turtles without Cluster API Operator](#install-rancher-turtles-without-cluster-api-operator-as-a-helm-dependency).

The Cluster API Operator is required for installing Rancher Turtles and you can choose whether you want to take care of this dependency yourself or let the Rancher Turtles Helm chart manage it for you. We recommend [installing as a dependency](#install-rancher-turtles-with-cluster-api-operator-as-a-helm-dependency) for the sake of simplicity, but the best option may depend on your specific configuration.

CAPI Operator allows handling the lifecycle of Cluster API Providers using a declarative approach, extending the capabilities of `clusterctl`. If you want to learn more about it, you can refer to [Cluster API Operator book](https://cluster-api-operator.sigs.k8s.io/). 

:::info
Before [installing Rancher Turtles](#install-rancher-turtles-with-cluster-api-operator-as-a-helm-dependency) in your Rancher environment, Rancher's `embedded-cluster-api` functionality must be disabled. This includes also cleaning up Rancher-specific webhooks that otherwise would conflict with CAPI ones.

To simplify setting up Rancher for installing Rancher Turtles, the official Rancher Turtles Helm chart includes a `pre-install` hook that applies these changes, making it transparent to the end user:
- Disable the `embedded-cluster-api` feature in Rancher.
- Delete the `mutating-webhook-configuration` and `validating-webhook-configuration` webhooks that are no longer needed.
:::

If you would like to understand how Rancher Turtles works and what the architecture looks like, you can refer to the [Architecture](../../reference-guides/architecture/intro.md) section.

:::note
If uninstalling, you can refer to [Uninstalling Rancher Turtles](../uninstall_turtles.md)
:::

### Install Rancher Turtles with `Cluster API Operator` as a Helm dependency

The `rancher-turtles` chart is available in https://rancher.github.io/turtles and this Helm repository must be added before proceeding with the installation:

```bash
helm repo add turtles https://rancher.github.io/turtles
helm repo update
```

As mentioned before, installing Rancher Turtles requires the [Cluster API Operator](https://github.com/kubernetes-sigs/cluster-api-operator) and the Helm chart can handle its installation automatically with a minimum set of flags:

```bash
helm install rancher-turtles turtles/rancher-turtles --version v0.8.0 \
    -n rancher-turtles-system \
    --dependency-update \
    --create-namespace --wait \
    --timeout 180s
```

This operation could take a few minutes and, after installing, you can take some time to study the installed controllers, including:
- `rancher-turtles-controller`.
- `capi-operator`.

:::note
- For a list of Rancher Turtles versions, refer to [Releases page](https://github.com/rancher/turtles/releases).
:::

This is the basic, recommended configuration, which manages the creation of a secret containing the required CAPI feature flags (`CLUSTER_TOPOLOGY`, `EXP_CLUSTER_RESOURCE_SET` and `EXP_MACHINE_POOL` enabled) in the core provider namespace. These feature flags are required to enable additional Cluster API functionality.

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
For detailed information on the values supported by the chart and their usage, refer to [Helm chart options](../../reference-guides/rancher-turtles-chart/values)
:::

### Install Rancher Turtles without `Cluster API Operator` as a Helm dependency

:::note
Remember that if you opt for this installation option, you will need to manage the Cluster API Operator installation yourself. You can follow the [CAPI Operator guide](../../contributing/install_capi_operator.md)
:::

The `rancher-turtles` chart is available in https://rancher.github.io/turtles and this Helm repository must be added before proceeding with the installation:

```bash
helm repo add turtles https://rancher.github.io/turtles
helm repo update
```

and then it can be installed into the `rancher-turtles-system` namespace with:

```bash
helm install rancher-turtles turtles/rancher-turtles --version v0.8.0
    -n rancher-turtles-system
    --set cluster-api-operator.enabled=false
    --set cluster-api-operator.cluster-api.enabled=false
    --create-namespace --wait
    --dependency-update
```

As you can see, we are telling Helm to ignore installing `cluster-api-operator` as a dependency.

This operation could take a few minutes and, after installing, you can take some time to study the installed controller:
- `rancher-turtles-controller`.
