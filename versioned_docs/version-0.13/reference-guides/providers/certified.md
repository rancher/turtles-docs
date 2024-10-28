---
sidebar_position: 1
---
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Certified CAPI Providers

Remember that most Cluster API Providers are upstream projects maintained by the Kubernetes open-source community.

## List of certified providers

:::note
This list is constantly evolving to reflect the ongoing development of the project.
:::

This is a list of the officially certified CAPI Providers by Turtles. These providers are covered by our test suite and we actively ensure that they work properly with the CAPI extension for Rancher.

| Platform        | Code Name                      | Provider Type            | Docs                     |
|-----------------|--------------------------------|--------------------------|--------------------------|
| **RKE2**        | CAPRKE2                    | Bootstrap/Control Plane  | https://rancher.github.io/cluster-api-provider-rke2 |
| **Kubeadm**     | Kubeadm                    | Bootstrap/Control Plane  | https://cluster-api.sigs.k8s.io/tasks/bootstrap/kubeadm-bootstrap |
| **AWS**         | CAPA                           | Infrastructure           | https://cluster-api-aws.sigs.k8s.io |
| **Docker**\*    | CAPD                           | Infrastructure           | https://cluster-api.sigs.k8s.io |
| **vSphere**         | CAPV                           | Infrastructure           | https://github.com/kubernetes-sigs/cluster-api-provider-vsphere |
| **Azure** (Only AKS managed clusters)         | CAPZ                           | Infrastructure           | https://capz.sigs.k8s.io/ |
| **Addon Provider Fleet**    | CAAPF                           | Addon           | http://github.com/rancher-sandbox/cluster-api-addon-provider-fleet |

*Recommended only for development purposes.

## List of providers in experimental mode

This is a list of providers that are in an advanced state of development and will soon become certified.

| Platform        | Code Name                      | Provider Type            | Docs                     |
|-----------------|--------------------------------|--------------------------|--------------------------|
| **Elemental**   | elemental                      | Infrastructure           | https://github.com/rancher-sandbox/cluster-api-provider-elemental |

## ClusterClass support for certified providers

The following is a support matrix for each certified provider and their support of the cluster topology feature:

<Tabs>
  <TabItem value="azure" label="CAPZ">
        - **Full support** of `ClusterClass`: both managed (AKS) and unmanaged (virtual machines) clusters can be provisioned via topology.
  </TabItem>
  <TabItem value="aws" label="CAPA" default>
        - **Supports** `ClusterClass` when provisioning unmanaged (EC2-based) clusters.
        - **Does not support** `ClusterClass` when provisioning managed (EKS) clusters: this is a work-in-progress.
    </TabItem>
  <TabItem value="rke2" label="CAPRKE2" default>
        - **Full support** of `ClusterClass`.
    </TabItem>
  <TabItem value="kubeadm" label="CABPK" default>
        - **Full support** of `ClusterClass`.
    </TabItem>
  <TabItem value="vsphere" label="CAPV" default>
        - **Full support** of `ClusterClass`.
    </TabItem>
  <TabItem value="docker" label="CAPD" default>
        - **Full support** of `ClusterClass`.
    </TabItem>
</Tabs>
