---
sidebar_position: 2
---

# Rancher Setup

## Installing Rancher

*If you're already running Rancher, you can skip this section and jump to [Install Rancher Turtles](./install-rancher-turtles/using_rancher_dashboard.md).*

Helm is the recommended way to install `Rancher` in an existing or new Kubernetes cluster.

:::tip
Make sure to follow one of the official [installation guides](https://ranchermanager.docs.rancher.com/pages-for-subheaders/installation-and-upgrade) for Rancher.
:::

Here's a minimal configuration example of a command to install `Rancher`:

```bash
helm install rancher rancher-stable/rancher
    --namespace cattle-system
    --create-namespace
    --set hostname=<rancher-hostname>
    --version <rancher-version>
    --wait
```

Replace `<rancher-hostname>` with the actual hostname of your `Rancher` server and use the `--version` option to specify the version of `Rancher` you want to install. In this case, use the [recommended](../getting-started/intro.md#prerequisites) `Rancher` version for `Rancher Turtles`.

You are now ready to install and use Rancher Turtles! 🎉
