---
sidebar_position: 2
---

# Upgrade Instructions for Early Adopters

If you were part of the early adopter programme and using Rancher Turtles v0.5.x or older, then you need to follow these instructions to use the GA version.

:::note
The early adopter programme was for non-production use. If you were using Turtles in a production, please contact Rancher support before upgrading.
:::

## Overview

To upgrade to the GA version of Rancher Turtles, you will need to do the following:

1. Delete any CAPI child cluster definitions
2. Helm uninstall Rancher Turtles
3. Clean-up orphaned resources
4. Install the GA version

## Steps

:::danger
These steps are destructive to child clusters created using CAPI. If you have any concerns, contact Rancher support before proceeding.
:::

1. Delete any CAPI child cluster definitions and wait for CAPI to fully delete the child clusters
2. Delete the CAPI providers installed:

```bash
kubectl delete capiproviders.turtles-capi.cattle.io -n capi-kubeadm-control-plane-system kubeadm-control-plane
kubectl delete capiproviders.turtles-capi.cattle.io -n capi-kubeadm-bootstrap-system kubeadm-bootstrap
kubectl delete capiproviders.turtles-capi.cattle.io -n capi-system cluster-api
kubectl delete ns capi-kubeadm-control-plane-system
kubectl delete ns capi-kubeadm-bootstrap-system
kubectl delete ns capi-system
```

3. Run the following to uninstall the Turtles extension:

```bash
helm uninstall -n rancher-turtles-system rancher-turtles --cascade foreground --wait
kubectl delete ns rancher-turtles-system
```

3. Run the following to delete any orphaned resources:

```bash
kubectl delete deployments.apps/capi-controller-manager -n capi-system --ignore-not-found=true
kubectl delete deployments.apps/capi-kubeadm-bootstrap-controller-manager -n capi-kubeadm-bootstrap-system --ignore-not-found=true
kubectl delete deployments.apps/capi-kubeadm-control-plane-controller-manager -n capi-kubeadm-control-plane-system --ignore-not-found=true
kubectl delete validatingwebhookconfigurations.admissionregistration.k8s.io capi-validating-webhook-configuration capi-kubeadm-bootstrap-validating-webhook-configuration capi-kubeadm-control-plane-validating-webhook-configuration --ignore-not-found=true
kubectl delete mutatingwebhookconfigurations.admissionregistration.k8s.io capi-mutating-webhook-configuration capi-kubeadm-bootstrap-mutating-webhook-configuration capi-kubeadm-control-plane-mutating-webhook-configuration --ignore-not-found=true
```

:::tip
If you are not going to continue using the extension then do the next step, instead go to the [re-enable embedded CAPI](#re-enable-embedded-capi) section.
:::

4. Follow the instructions to install the new version of the extension [here](../getting-started/install-rancher-turtles/using_rancher_dashboard.md).

## Re-enable Embedded CAPI

:::note
This step is only required if you are not going to use the Rancher Turtles extension any further.
:::

1. Create a feature.yaml file, with embedded-cluster-api set to true:

```yaml
apiVersion: management.cattle.io/v3
kind: Feature
metadata:
  name: embedded-cluster-api
spec:
  value: true
```

2. Use kubectl to apply the feature.yaml file to the cluster:

```bash
kubectl apply -f feature.yaml
```
