---
sidebar_position: 1
---

# Introduction

This section contains information on how you can leverage the existing E2E suite to integrate any CAPI providers with Turtles and verify that the provisioning and importing of clusters works as expected. The validation performs the following actions:

- Create a management cluster in the desired environment.
- Install Rancher and Turtles with all prerequisites.
- Install Gitea.
- Run the suite that will create a git repo, apply cluster template using Fleet and verify the cluster is created and successfully imported in Rancher.

The test suite can be used for certification of providers not listed in the [Certification table](../../reference-guides/providers/certified.md), as detailed in [Provider Certification](../../tasks/provider-certification/intro.md).
