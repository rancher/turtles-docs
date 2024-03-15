---
sidebar_position: 2
---

# CAPIProvider Resource

The `CAPIProvider` resource allows managing Cluster API Operator manifests in a declarative way. It is used to provision and configure Cluster API providers like AWS, vSphere etc.

`CAPIProvider` follows a GitOps model - the spec fields are declarative user inputs. The controller only updates status.

Every field provided by the upstream CAPI Operator resource for the desired `spec.type` is also available in the spec of the `CAPIProvider` resouce. Feel free to refer to upstream configuration [guides](https://cluster-api-operator.sigs.k8s.io/03_topics/02_configuration/) for advanced scenarios.

[ARD](https://github.com/rancher/turtles/blob/main/docs/adr/0007-rancher-turtles-public-api.md)

## Usage

To use the `CAPIProvider` resource:

1. Create a `CAPIProvider` resource with the desired provider name, type, credentials, configuration, and features.
1. The `CAPIProvider` controller will handle templating the required Cluster API Operator manifests based on the `CAPIProvider` spec.
1. The status field on the `CAPIProvider` resource will reflect the state of the generated manifests.
1. Manage the `CAPIProvider` object declaratively to apply changes to the generated provider manifests.

Here is an example `CAPIProvider` manifest:

```yaml
apiVersion: turtles-capi.cattle.io/v1alpha1
kind: CAPIProvider
metadata:
  name: aws-infra
  namespace: default
spec:
  name: aws
  type: infrastructure
  credentials:
    rancherCloudCredential: aws-creds # Rancher credentials secret for AWS
  configSecret:
    name: aws-config
  features:
    clusterResourceSet: true
```

This will generate an AWS infrastructure provider with the supplied mapping for rancher credential secret and custom enabled features.

The `CAPIProvider` controller will own all the generated provider resources, allowing garbage collection by deleting the `CAPIProvider` object.

## Specification

The key fields in the `CAPIProvider` spec are:

- `name` - Name of the provider (aws, vsphere etc). Inherited from `metadata.name` if is not specified.
- `type` - Kind of provider resource (infrastructure, controlplane etc)
- `credentials` - Source credentials for provider specification
- `configSecret` - Name of the provider config secret, where the variables and synced credential will be stored. By default if not specified, will inherit the name of the `CAPIProvider` resource
- `features` - Enabled provider features
- `variables` - Variables is a map of environment variables to add to the content of the `configSecret`

Full documentation on the CAPIProvider resource - [here](https://doc.crds.dev/github.com/rancher/turtles/turtles-capi.cattle.io/CAPIProvider/v1alpha1@v0.5.0).

## Deletion

When a `CAPIProvider` resource is deleted, the kubernetes garbage collection will clean up all the generated provider resources that it owns. This includes:

- Cluster API Operator resource instance
- Secret referenced by the `configSecret`
