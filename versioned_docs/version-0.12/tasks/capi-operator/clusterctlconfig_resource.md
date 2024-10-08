---
sidebar_position: 3
---

# ClusterctlConfig Resource

The `ClusterctlConfig` resource allows managing overrides for clusterctl (CAPI Operator) configurations in a declarative way. It is used to configure clusterctl providers and their urls, as well as version restrictions.

`ClusterctlConfig` follows a GitOps model - the spec fields are declarative user inputs. Turtles does not create or update the resource, it is up to user to specify provider url overrides and maintain its state. It takes precedence over embedded defaults or `clusterctl` default set of provider definitions.

[ADR](https://github.com/rancher/turtles/blob/main/docs/adr/0012-clusterctl-provider.md)

## Usage

To use the `ClusterctlConfig` resource:

1. Create a `ClusterctlConfig` resource with the `clusterctl-config` name in the `turtles` namespace.
1. The `ClusterctlConfig` controller will handle updates for the `ConfigMap` mounted to the `cluster-api-operator` with the required clusterctl configuration based on the `ClusterctlConfig` spec.
1. Manage the `ClusterctlConfig` object declaratively to apply changes to the generated provider configurations. It may require some time for changes to take effect, as `kubelet` takes care of updating mounted point based on `ConfigMap` state.

Here is an example `ClusterctlConfig` manifest:

```yaml
apiVersion: turtles-capi.cattle.io/v1alpha1
kind: ClusterctlConfig
metadata:
  name: clusterctl-config
  namespace: rancher-turtles-system
spec:
  providers:
  - name: metal3
    url: https://github.com/metal3-io/cluster-api-provider-metal3/releases/v1.7.1/infrastructure-components.yaml
    type: InfrastructureProvider
```

This example will generate a clusterctl configuration for the `metal3` provider with the specified URL and type.

## Specification

The key fields in the `ClusterctlConfig` spec are:

- `providers[].name` - Name of the provider (e.g. metal3)
- `providers[].url` - URL of the provider configuration (e.g. 
https://github.com/metal3-io/cluster-api-provider-metal3/releases/v1.7.1/infrastructure-components.yaml). This can use `latest` release, if supported, or pin the maximum version to `v1.7.1` for example.
- `providers[].type` - Type of the provider (e.g. InfrastructureProvider)

## Deletion

When a `ClusterctlConfig` resource is deleted, the config map is reverted to its original state, managed by turtles.