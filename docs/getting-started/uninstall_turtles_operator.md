---
sidebar_position: 5
---

# Uninstall Rancher Turtles Operator

This gives an overview of Rancher Turtles Operator uninstallation process.

:::caution
When installing Rancher Turtles in your Rancher environment, by default, Rancher Turtles enables the Cluster API Operator cleanup. This includes cleaning up Cluster API Operator specific webhooks and deployments that otherwise would block the uninstallation.

To simplify uninstalling Rancher Turtles, the official Rancher Turtles Helm chart includes a `post-delete` hook that applies these changes, making it transparent to the end user:
- Delete the `mutating-webhook-configuration` and `validating-webhook-configuration` webhooks that are no longer needed.
- Delete the CAPI `deployments` that are no longer needed.
:::

Once uninstalled, Rancher's `embedded-cluster-api` functionality can be enabled:

1. Create a `feature.yaml` file, with `embedded-cluster-api` set to false:
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