---
sidebar_position: 3
---

# Development setup

## Prerequisites:

- [kind](https://kind.sigs.k8s.io/)
- [helm](https://helm.sh/)
- [tilt](https://tilt.dev/)

## Create a local development environment

1. Clone the [Rancher Turtles](https://github.com/rancher/turtles) repository locally

2. Create **tilt-settings.yaml**:

```yaml
{
    "k8s_context": "k3d-rancher-test",
    "default_registry": "ghcr.io/turtles-dev",
    "debug": {
        "turtles": {
            "continue": true,
            "port": 40000
        }
    }
}
```

3. Open a terminal in the root of the Rancher Turtles repository
4. Run the following:

```bash
make dev-env

# Or if you want to use a custom hostname for Rancher
RANCHER_HOSTNAME=my.customhost.dev make dev-env
```

5. When tilt has started, open a new terminal and start ngrok or inlets

```bash
kubectl port-forward --namespace cattle-system svc/rancher 10000:443
ngrok http https://localhost:10000
```

## What happens when you run `make dev-env`?

1. A [kind](https://kind.sigs.k8s.io/) cluster is created with the following [configuration](https://github.com/rancher/turtles/blob/main/scripts/kind-cluster-with-extramounts.yaml).
1. [Cluster API Operator](../contributing/install_capi_operator.md) is installed using helm, which includes: 
    - Core Cluster API controller
    - Kubeadm Bootstrap and Control Plane Providers
    - Docker Infrastructure Provider
    - Cert manager
1. `Rancher manager` is installed using helm.
1. `tilt up` is run to start the development environment.
