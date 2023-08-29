---
sidebar_position: 3
---

# Install CAPI Operator

## Installing CAPI and providers

There are two options when it comes to installing [CAPI](https://github.com/kubernetes-sigs/cluster-api). You can install CAPI and desired providers with [clusterctl](https://cluster-api.sigs.k8s.io/user/quick-start#install-clusterctl), following the official [quick-start](https://cluster-api.sigs.k8s.io/user/quick-start#installation) guide, or you can use helm based installation with [Cluster API Operator](https://github.com/kubernetes-sigs/cluster-api-operator).

### Cluster API operator path

For example, to install latest version of CAPI + docker provider, you can setup helm repo:
```bash
helm repo add capi-operator https://kubernetes-sigs.github.io/cluster-api-operator
helm repo update
```
and install it together with [cert-manager](https://github.com/cert-manager/cert-manager) with:
```bash
helm install capi-operator capi-operator/cluster-api-operator
	--create-namespace -n capi-operator-system
	--set infrastructure=docker
	—-set cert-manager.enabled=true
	—-timeout 90s --wait # core Cluster API with kubeadm bootstrap and control plane providers will also be installed 
```
*Note: `cert-manager` is a hard requirement for CAPI and Cluster API Operator*

For any scenarios when you need to provide additional environment variables, to choose some feature gates or provide cloud credentials, similar to clusterctl common provider [initialization options](https://cluster-api.sigs.k8s.io/user/quick-start#initialization-for-common-providers), in Cluster API Operator you could use variables secret. You need to specify the `name` and the `namespace` of the secret for the Cluster API Operator in helm:
```bash
helm install capi-operator capi-operator/cluster-api-operator
	--create-namespace -n capi-operator-system 
	--set infrastructure=docker
	—-set cert-manager.enabled=true
	—-timeout 90s
	—-secret-name <secret_name>
	—-secret-namespace <secret_namespace>
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

You can select providers to install together with Cluster API Operator. To do so, specify one or more providers in `—-infrastructure` flag when starting with helm. You can pass a specific version, if needed:
```bash
helm install … infrastructure=docker:v1.5.0,azure # For azure assumed latest version
```

### Clusterctl path

Follow the official [quick-start](https://cluster-api.sigs.k8s.io/user/quick-start#installation) guide.