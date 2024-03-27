---
sidebar_position: 5
---

# Uninstall Rancher Turtles

This gives an overview of Rancher Turtles Operator uninstallation process. 

:::caution
When installing Rancher Turtles in your Rancher environment, by default, Rancher Turtles enables the Cluster API Operator cleanup. This includes cleaning up Cluster API Operator specific webhooks and deployments that otherwise cause issues with Rancher provisioning.

To simplify uninstalling Rancher Turtles (via Rancher Manager or helm command), the official Rancher Turtles Helm chart includes a `post-delete` hook that applies these changes, making it transparent to the end user:
- Delete the `mutating-webhook-configuration` and `validating-webhook-configuration` webhooks that are no longer needed.
- Delete the CAPI `deployments` that are no longer needed.
:::

There are two options to uninstall the Rancher Turtles Operator depending on the installation method.

1. Rancher Turtles Operator installed via Rancher Manager (i.e in local cluster, `Apps->Repositories` to add a turtles repository then `Apps->Charts` to install rancher-turtles extension). To uninstall, simply navigate to local cluster, `Apps->Installed Apps`, find `rancher-turtles` extension and click `Delete`. 

2. Rancher Turtles Operator installed via [helm command](./install-rancher-turtles/using_helm.md). If you would like to uninstall it manually,
it can be simply achieived via `helm`:

```bash
helm uninstall -n rancher-turtles-system rancher-turtles
```

This may take a few minutes to complete.

:::note
Remember that, if you use a different name for the installation or a different namespace, you may need to customize the command for your specific configuration.
:::

Once uninstalled, Rancher's `embedded-cluster-api` feature must be re-enabled:

1. Create a `feature.yaml` file, with `embedded-cluster-api` set to true:
```yaml title="feature.yaml"
apiVersion: management.cattle.io/v3
kind: Feature
metadata:
  name: embedded-cluster-api
spec:
  value: true
```
2. Use `kubectl` to apply the `feature.yaml` file to the cluster:
```bash
kubectl apply -f feature.yaml
```
