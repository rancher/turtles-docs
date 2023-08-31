---
sidebar_position: 3
---

# Install CAPI Operator

## Installing CAPI and providers

CAPI and desired CAPI providers could be installed using helm based installation for [`Cluster API Operator`](https://github.com/kubernetes-sigs/cluster-api-operator).

### Cluster API operator path

For example, to install latest version of CAPI + docker provider, first, the helm repo should be added with:
```bash
helm repo add capi-operator https://kubernetes-sigs.github.io/cluster-api-operator
helm repo update
```
and installed together with [cert-manager](https://github.com/cert-manager/cert-manager) using:
```bash
helm install capi-operator capi-operator/cluster-api-operator
	--create-namespace -n capi-operator-system
	--set infrastructure=docker
	--set cert-manager.enabled=true
	--timeout 90s --wait # core Cluster API with kubeadm bootstrap and control plane providers will also be installed
```
*Note: `cert-manager` is a hard requirement for CAPI and `Cluster API Operator`*

For any scenarios when there is a need to provide additional environment variables, to choose some feature gates or provide cloud credentials, similar to `clusterctl` common provider [initialization options](https://cluster-api.sigs.k8s.io/user/quick-start#initialization-for-common-providers), in `Cluster API Operator` a variables secret could be used. A `name` and a `namespace` of the secret could be specified for the `Cluster API Operator`:
```bash
helm install capi-operator capi-operator/cluster-api-operator
	--create-namespace -n capi-operator-system
	--set infrastructure=docker:v1.4.4
	--set cert-manager.enabled=true
	--timeout 90s
	--secret-name <secret_name>
	--secret-namespace <secret_namespace>
	--wait
```

Example secret data:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: variables
  namespace: default
type: Opaque
stringData:
  CLUSTER_TOPOLOGY: "true"
  EXP_CLUSTER_RESOURCE_SET: "true"
```

To select more then one desired providers to be installed together with `Cluster API Operator`, `--infrastructure` flag could be specified with multiple provider names separated by `,`. Also a specific version could be specified like so:
```bash
helm install ... infrastructure=docker:v1.4.4,azure:v1.4.4
```

For more fine-grained control on the providers and other components installed with CAPI see the [Add infrastructure provider](../tasks/capi-operator/add_infrastructure_provider.md) section.
