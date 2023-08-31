---
sidebar_position: 2
---

# Add Infrastructure Provider

This section describes how an infrastructure provider such as Azure could be added using Cluster API Operator.

## Basic Cluster API Provider Installation

This section describes the basic process of installing Cluster API providers using the operator. The Cluster API operator manages five types of objects:

- CoreProvider
- BootstrapProvider
- ControlPlaneProvider
- InfrastructureProvider
- AddonProvider

Please note that this example provides a basic configuration of Azure Infrastructure provider for getting started. More detailed examples and CRD descriptions will be provided in subsequent sections of this document.

### Installing the CoreProvider

The first step is to install the CoreProvider, which is responsible for managing the Cluster API CRDs and the Cluster API controller.

Any existing namespace could be utilized for providers in the Kubernetes cluster. However, before creating a provider object, make sure the specified namespace has been created. In the example below, we use the `capi-system` namespace. To create this namespace through either the Command Line Interface (CLI) by running `kubectl create namespace capi-system`, or the declarative approach described in the [official Kubernetes documentation](https://kubernetes.io/docs/tasks/administer-cluster/namespaces-walkthrough/#create-new-namespaces) could be used.

*Example:*

```yaml
apiVersion: operator.cluster.x-k8s.io/v1alpha2
kind: CoreProvider
metadata:
  name: cluster-api
  namespace: capi-system
spec:
  version: v1.4.4
```

**Note:** Only one CoreProvider can be installed at the same time on a single cluster.

### Installing Azure Infrastructure Provider

Next, install [Azure Infrastructure Provider](https://capz.sigs.k8s.io/). Before that ensure that `capz-system` namespace exists.

Since the provider requires variables to be set, create a secret containing them in the same namespace as the provider. It is also recommended to include a `github-token` in the secret. This token is used to fetch the provider repository, and it is required for the provider to be installed. The operator may exceed the rate limit of the GitHub API without the token. Like [clusterctl](https://cluster-api.sigs.k8s.io/clusterctl/overview.html?highlight=github_token#avoiding-github-rate-limiting), the token needs only the `repo` scope.

```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: azure-variables
  namespace: capz-system
type: Opaque
stringData:
  AZURE_CLIENT_ID_B64: Zm9vCg==
  AZURE_CLIENT_SECRET_B64: Zm9vCg==
  AZURE_SUBSCRIPTION_ID_B64: Zm9vCg==
  AZURE_TENANT_ID_B64: Zm9vCg==
  github-token: ghp_fff
---
apiVersion: operator.cluster.x-k8s.io/v1alpha1
kind: InfrastructureProvider
metadata:
 name: azure
 namespace: capz-system
spec:
 version: v1.9.3
 configSecret:
   name: azure-variables
```

### Deleting providers

To remove the installed providers and all related kubernetes objects just delete the following CRs:

```bash
kubectl delete coreprovider cluster-api
kubectl delete infrastructureprovider azure
```
