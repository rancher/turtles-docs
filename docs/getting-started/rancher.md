---
sidebar_position: 2
---

# Rancher Setup

## Installing rancher

In order to install rancher in existing/new k8s cluster, aside from following one of the [official installation guides](https://ranchermanager.docs.rancher.com/pages-for-subheaders/installation-and-upgrade), rancher should have `embedded-cluster-api` feature disabled. To do so, when installing rancher in your k8s cluster, you need to specify the following feature flag with the `helm` command:

```bash
helm install rancher rancher-stable/rancher
	--set features=embedded-cluster-api=false # Disabling embedded CAPI feature
	--set hostname=<rancher-hostname>
	--version v2.7.5
	--namespace cattle-system
	--create-namespace --wait
```