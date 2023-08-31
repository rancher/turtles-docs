---
sidebar_position: 2
---

# Rancher Setup

## Installing rancher

In order to install rancher in existing/new k8s cluster, aside from following one of the [official installation guides](https://ranchermanager.docs.rancher.com/pages-for-subheaders/installation-and-upgrade), rancher should have `embedded-cluster-api` feature disabled. To do so, when installing rancher in the k8s cluster, the following feature flag should be specified with the `helm` command:

```bash
helm install rancher rancher-stable/rancher
	--set features=embedded-cluster-api=false # Disabling embedded CAPI feature
	--set hostname=<rancher-hostname>
	--version v2.7.5
	--namespace cattle-system
	--create-namespace --wait
```

## Existing rancher installation

To allow Rancher Turtles to be installed in the existing rancher cluster, `embedded-cluster-api` feature should be created from the manifest first:
```yaml title="feature.yaml"
apiVersion: management.cattle.io/v3
kind: Feature
metadata:
  name: embedded-cluster-api
spec:
  value: false
```

To apply it in the cluster:
```bash
kubectl apply -f feature.yaml
```
