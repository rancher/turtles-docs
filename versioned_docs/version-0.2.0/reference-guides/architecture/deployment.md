---
sidebar_position: 0
---

# Deployment Scenarios

:::note
The early release of Rancher Turtles only supports having Rancher Manager and 
Rancher Turtles running in the same cluster. A topology with a separate Rancher 
Manager cluster and one/multiple CAPI management cluster/s will be supported in 
future releases.
:::

## Rancher Manager & CAPI Management Combined

In this topology, both Rancher Manager and Rancher Turtles are deployed to the 
same Kubernetes cluster, and it acts as a centralized management cluster.

![Rancher Manager & CAPI Management Combined](./in_cluster_topology.png)

This architecture offers a simplified deployment of components and provides a 
single view of all clusters. On the flip side, it's important to consider that 
the number of clusters that can be managed effectively by Cluster API (CAPI) is 
limited by the resources available within the single management cluster. 
