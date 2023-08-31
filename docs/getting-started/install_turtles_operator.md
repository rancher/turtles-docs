---
sidebar_position: 4
---

# Install Rancher Turtles Operator

Rancher turtles requires to have a direct connection to the Rancher Manager cluster. This could be achieved by:
1. Installing in in the same cluster as the Rancher Manager.

Recommended path of installation for the operator is by using helm. To install it in the cluster a repository should be added first:
```bash
helm repo add turtles https://rancher-sandbox.github.io/rancher-turtles/
helm repo update
```
and installed into `rancher-turtles-system` namespace with:
```bash
helm install rancher-turtles turtles/rancher-turtles
    -n rancher-turtles-system
    --create-namespace --wait
```
