---
sidebar_position: 2
---

# Rancher Setup

## Installing Rancher

*If you're already running Rancher, you can skip this section and jump to [Setting up Rancher for Rancher Turtles](#setting-up-rancher-for-rancher-turtles).*

Helm is the recommended way to install `Rancher` in an existing or new Kubernetes cluster.

:::tip
Make sure to follow one of the official [installation guides](https://ranchermanager.docs.rancher.com/pages-for-subheaders/installation-and-upgrade) for Rancher.
:::

Here's a minimal configuration example of a command to install `Rancher`:

```bash
helm install rancher rancher-stable/rancher
    --namespace cattle-system
    --create-namespace
    --set hostname=<rancher-hostname>
    --version <rancher-version>
    --wait
```

Replace `<rancher-hostname>` with the actual hostname of your `Rancher` server and use the `--version` option to specify the version of `Rancher` you want to install. In this case, use the [recommended](../getting-started/intro.md#prerequisites) `Rancher` version for `Rancher Turtles`.

## Setting up Rancher for Rancher Turtles

Before installing Rancher Turtles in your Rancher environment, the `embedded-cluster-api` functionality must be disabled:

1. Create a `feature.yaml` file, with `embedded-cluster-api` set to false:
```yaml title="feature.yaml"
apiVersion: management.cattle.io/v3
kind: Feature
metadata:
  name: embedded-cluster-api
spec:
  value: false
```
2. Use `kubectl` to apply the `feature.yaml` file to the cluster:
```bash
kubectl apply -f feature.yaml
```

Your Rancher installation is now ready to install and use Rancher Turtles! 🎉
