---
sidebar_position: 5
---

# Uninstall Rancher Turtles

This gives an overview of Rancher Turtles uninstallation process. 

:::caution
When installing Rancher Turtles in your Rancher environment, by default, Rancher Turtles enables the Cluster API Operator cleanup. This includes cleaning up Cluster API Operator specific webhooks and deployments that otherwise cause issues with Rancher provisioning.

To simplify uninstalling Rancher Turtles (via Rancher Manager or helm command), the official Rancher Turtles Helm chart includes a `post-delete` hook that applies these changes, making it transparent to the end user:
- Delete the `mutating-webhook-configuration` and `validating-webhook-configuration` webhooks that are no longer needed.
- Delete the CAPI `deployments` that are no longer needed.
:::

To uninstall the Rancher Turtles Extension use the following helm command:

```bash
helm uninstall -n rancher-turtles-system rancher-turtles --cascade foreground --wait
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
