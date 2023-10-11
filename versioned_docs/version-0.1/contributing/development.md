---
sidebar_position: 2
---

# Development setup

## Prerequisites:

- [kind](https://kind.sigs.k8s.io/)
- [helm](https://helm.sh/)
- [clusterctl](https://cluster-api.sigs.k8s.io/user/quick-start.html#install-clusterctl)
- [tilt](https://tilt.dev/)


## Create a local development environment

1. Clone the [Rancher Turtles](https://github.com/rancher-sandbox/rancher-turtles) repository locally

2. Create **tilt-settings.yaml**:

```yaml
{
    "k8s_context": "k3d-rancher-test",
    "default_registry": "ghcr.io/rancher-turtles-dev",
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

1. A [kind](https://kind.sigs.k8s.io/) cluster is created with the following [configuration](https://github.com/rancher-sandbox/rancher-turtles/blob/main/scripts/kind-cluster-with-extramounts.yaml).
1. [Cert manager](https://cert-manager.io/) is installed on the cluster, which is a requirement for running `Rancher turtes` extension.
1. `clusterctl` is used to bootstrap CAPI components onto the cluster, and a default configuration includes: core Cluster API controller, Kubeadm bootstrap and control plane providers, Docker infrastructure provider.
1. `Rancher manager` is installed using helm.
1. `tilt up` is run to start the development environment.
