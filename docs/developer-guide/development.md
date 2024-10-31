---
sidebar_position: 3
---

# Development setup

## Prerequisites:

- [kind](https://kind.sigs.k8s.io/)
- [helm](https://helm.sh/)
- [tilt](https://tilt.dev/)

## Environment description

The Turtles development environment relies on Tilt for hot-reloading and a Kind cluster with Rancher and Turtles installed. We also use ngrok to expose the Rancher server to the internet.
The environment can be provisioned using the `make dev-env` command. However, there are some prerequisites that need to be met before running this command. Executing the command will perform the following actions:

1. Create a [kind](https://kind.sigs.k8s.io/) cluster with the following [configuration](https://github.com/rancher/turtles/blob/main/scripts/kind-cluster-with-extramounts.yaml). This cluster will be used to run 
Rancher and Turtles and will also be prepared to work with CAPD (CAPI Provider Docker).
2. In this setup, we will install the [Cluster API Operator](../developer-guide/install_capi_operator.md) separately from Turtles for a decentralized development environment. 
The CAPI Operator will be installed using a Helm chart and will include the following providers:
    - Core Cluster API
    - Kubeadm Bootstrap and Control Plane Providers
    - Docker Infrastructure Provider
    - Cert Manager
3. Install the ngrok ingress controller to expose the Rancher server to the internet.
4. Install the `Rancher manager` using Helm.
5. Run `tilt up` to install Turtles and enable hot-reloading of containers. 

## Create a local development environment

1. Clone the [Rancher Turtles](https://github.com/rancher/turtles) repository locally.

2. Create **tilt-settings.yaml**:

```yaml
{
    "default_registry": "ghcr.io/turtles-dev",
    "debug": {
        "turtles": {
            "continue": true,
            "port": 40000
        }
    }
}
```

3. Open a terminal in the root of the Rancher Turtles repository.
4. Run the following:

```bash
# Rancher hostname can be configured using ngrok.
RANCHER_HOSTNAME=my.customhost.dev NGROK_API_KEY="<api-key>" NGROK_AUTHTOKEN="<api-authtoken>" make dev-env
```
