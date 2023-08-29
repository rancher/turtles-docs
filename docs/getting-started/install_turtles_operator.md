---
sidebar_position: 4
---

# Install Rancher Turtles Operator

Rancher turtles requires to have a direct connection to the rancher manager cluster. This could be achieved by:
1. Installing in in the same cluster as the rancher manager.

Recommended path of installation for the operator is by using helm. To install it in your cluster you need to add repository first:
```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```
and install it into `rancher-turtles-system` namespace with:
```bash
helm install rancher-turtles turtles/rancher-turtles
    -n rancher-turtles-system
    --create-namespace --wait
```
