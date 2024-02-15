---
sidebar_position: 3
---

# Add Infrastructure Provider

This section describes how an infrastructure provider such as `Azure` could be added using `Cluster API Operator`.

### Installing Azure Infrastructure Provider

Next, install [Azure Infrastructure Provider](https://capz.sigs.k8s.io/). Before that ensure that `capz-system` namespace exists.

Since the provider requires variables to be set, create a secret containing them in the same namespace as the provider. It is also recommended to include a `github-token` in the secret. This token is used to fetch the provider repository, and it is required for the provider to be installed. The operator may exceed the rate limit of the GitHub API without the token. Like [clusterctl](https://cluster-api.sigs.k8s.io/clusterctl/overview.html?highlight=github_token#avoiding-github-rate-limiting), the token needs only the `repo` scope.

## Option 1: CAPIProvider resource

This section describes how to install the Azure `InfrastructureProvider` via `CAPIProvider`, which is responsible for managing Cluster API Azure CRDs and the Cluster API Azure controller.

*Example:*

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
apiVersion: turtles-capi.cattle.io/v1alpha1
kind: CAPIProvider
metadata:
  name: azure
  namespace: capz-system
spec:
  version: v1.9.3
  type: infrastructure # required
  configSecret:
    name: azure-variables # This will additionally poulate the default set of feature gates for the provider
```

## Option 2: Raw InfrastructureProvider resource

This section describes how to install the Azure `InfrastructureProvider`, which is responsible for managing the Cluster API Azure CRDs and the Cluster API Azure controller.

*Example:*

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
apiVersion: operator.cluster.x-k8s.io/v1alpha2
kind: InfrastructureProvider
metadata:
 name: azure
 namespace: capz-system
spec:
 version: v1.9.3
 configSecret:
   name: azure-variables

```
:::tip
**There are known issues when provisioning clusters using CAPZ v1.12.0.**
We recommend using version v1.11.5.
:::

### Deleting providers

To remove the installed providers and all related Kubernetes objectsm just delete the following CRs:

```bash
kubectl delete coreprovider cluster-api
kubectl delete infrastructureprovider azure
```
