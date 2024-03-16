---
sidebar_position: 2
---

# Create a cluster using Fleet

This section will guide you through creating a cluster that utilizes ClusterClass using a GitOps workflow with Fleet.

:::note
This guide uses the [examples repository](https://github.com/rancher-sandbox/rancher-turtles-fleet-example/tree/clusterclass).
:::

## Prerequisites

- Rancher Manager cluster with Rancher Turtles installed
- Cluster API providers installed for your scenario - we'll be using the Docker infrastructure and Kubeadm bootstrap/control plane providers in these instructions - see [Initialization for common providers](https://cluster-api.sigs.k8s.io/user/quick-start.html#initialization-for-common-providers)
- The **ClusterClass** feature enabled - see [the introduction](./intro.md)

## Configure Rancher Manager

The clusterclass and cluster definitions will be imported into the Rancher Manager cluster (which is also acting as a Cluster API management cluster) using the **Continuous Delivery** feature (which uses Fleet).

The guide will apply the manifests using a 2-step process. However, this isn't required and they could be combined into one step.

There are 2 options to provide the configuration. The first is using the Rancher Manager UI and the second is by applying some YAML to your cluster. Both are covered below.

### Import ClusterClass Definitions

#### Using the Rancher Manager UI

1. Go to Rancher Manager
2. Select **Continuos Delivery** from the menu:
3. Select **fleet-local** as the namespace from the top right
4. Select **Git Repos** from the sidebar
5. Click **Add Repository**
6. Enter **classes** as the name
7. Get the **HTTPS** clone URL from your git repo
8. Add the URL into the **Repository URL** field
9. Change the branch name to **clusterclass**
10. Click **Add Path**
11. Enter `/classes`
12. Click **Next**
13. Click **Create**
14. Click on the **clusters** name
15. Watch the resources become ready

### Using kubectl

1. Get the **HTTPS** clone URL from your git repo
2. Create a new file called **repo.yaml**
3. Add the following contents to the new file:

```yaml
apiVersion: fleet.cattle.io/v1alpha1
kind: GitRepo
metadata:
  name: classes
  namespace: fleet-local
spec:
  branch: clusterclass
  repo: https://github.com/rancher-sandbox/rancher-turtles-fleet-example.git
  paths:
    - /classes
  targets: []
```

4. Apply the file to the Rancher Manager cluster using **kubectl**:

```bash
kubectl apply -f repo.yaml
```

5. Go to Rancher Manager
6. Select **Continuos Delivery** from the side bar
7. Select **fleet-local** as the namespace from the top right
8. Select **Git Repos** from the sidebar
9. Click on the **clusters** name
10. Watch the resources become ready
11. Select **Cluster Management** from the menu
12. Check your cluster has been imported

### Import Cluster Definitions

Now the classes have been imported its possible to use them with cluster definitions.

#### Using the Rancher Manager UI

1. Go to Rancher Manager
2. Select **Continuos Delivery** from the menu:
3. Select **fleet-local** as the namespace from the top right
4. Select **Git Repos** from the sidebar
5. Click **Add Repository**
6. Enter **clusters** as the name
7. Get the **HTTPS** clone URL from your git repo
8. Add the URL into the **Repository URL** field
9. Change the branch name to **clusterclass**
10. Click **Add Path**
11. Enter `/clusters`
12. Click **Next**
13. Click **Create**
14. Click on the **clusters** name
15. Watch the resources become ready
16. Select **Cluster Management** from the menu
17. Check your cluster has been imported

### Using kubectl

1. Get the **HTTPS** clone URL from your git repo
2. Create a new file called **repo.yaml**
3. Add the following contents to the new file:

```yaml
apiVersion: fleet.cattle.io/v1alpha1
kind: GitRepo
metadata:
  name: clusters
  namespace: fleet-local
spec:
  branch: clusterclass
  repo: https://github.com/rancher-sandbox/rancher-turtles-fleet-example.git
  paths:
    - /clusters
  targets: []
```

4. Apply the file to the Rancher Manager cluster using **kubectl**:

```bash
kubectl apply -f repo.yaml
```

5. Go to Rancher Manager
6. Select **Continuos Delivery** from the side bar
7. Select **fleet-local** as the namespace from the top right
8. Select **Git Repos** from the sidebar
9. Click on the **classes** name
10. Watch the resources become ready
11. Select **Cluster Management** from the menu
12. Check your cluster has been imported
