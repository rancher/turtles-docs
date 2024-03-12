---
sidebar_position: 3
---

# Installing Azure Infrastructure Provider using CAPIProvider resource

This section describes how to install the Azure `InfrastructureProvider` via `CAPIProvider`, which is responsible for managing Cluster API Azure CRDs and the Cluster API Azure controller.

:::note
This section describes how to install the raw Azure `InfrastructureProvider`, which is responsible for managing the Cluster API Azure CRDs and the Cluster API Azure controller. The detailed configuration steps are described in the [official](https://cluster-api-operator.sigs.k8s.io/03_topics/03_basic-cluster-api-provider-installation/02_installing-capz#installing-azure-infrastructure-provider) CAPI Operator documentation.
:::

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
    name: azure-variables # This will additionally populate the default set of feature gates for the provider inside the secret
```

### Deleting providers

To remove the installed providers and all related Kubernetes objectsm just delete the following CRs:

```bash
kubectl delete coreprovider cluster-api
kubectl delete infrastructureprovider azure
```
