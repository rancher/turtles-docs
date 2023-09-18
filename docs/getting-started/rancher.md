---
sidebar_position: 2
---

# Rancher Setup

## Installing Rancher

To install `Rancher` in an existing or new Kubernetes cluster, you can use the following steps:

1. First, make sure to follow one of the official [installation guides](https://ranchermanager.docs.rancher.com/pages-for-subheaders/installation-and-upgrade) for `Rancher`.
2. When installing `Rancher` using the `helm` command, use the `--set` option to specify the `features` parameter. For the `embedded-cluster-api` feature, set the value to `false` to disable it.
3. Use the `--version` option to specify the version of `Rancher` you want to install. In this case, use the [recommended](../getting-started/intro.md#prerequisites) `Rancher` version for `Rancher Turtles`.

Here's the complete command to install `Rancher` with the `embedded-cluster-api` feature disabled. Replace `<rancher-hostname>` with the actual hostname of your `Rancher` server:

```bash
helm install rancher rancher-stable/rancher --set features=embedded-cluster-api=false --set hostname=<rancher-hostname> --set version=<rancher-version> --set namespace=cattle-system --create-namespace --wait
```

## Existing Rancher installation

To install `Rancher Turtles` in an existing `Rancher` cluster, follow these steps:

1. Create a `feature.yaml` file, with `embedded-cluster-api` feature disabled:
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