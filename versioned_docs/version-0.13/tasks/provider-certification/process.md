---
sidebar_position: 1
---

# Provider Certification Guide

The process of certification is based on verifying Rancher Turtles integration with CAPI providers. To simplify this task, we prepared a generic test that validates the provisioning and importing of a downstream CAPI cluster.

:::tip
We recommend you refer to this [example](https://github.com/rancher-sandbox/turtles-integration-suite-example) on how to use Turtles' test suite.
:::

## Test & Certify your provider

The first step in validating that your provider is compatible with Turtles and that you can provision CAPI clusters and import them into Rancher via Turtles is to integrate with our test suite. We provide a repository with an [integration example](https://github.com/rancher-sandbox/turtles-integration-suite-example) that you can use as a reference for your integration.

Turtles as a project contains a [number of suites](https://github.com/rancher/turtles/tree/main/test/e2e/suites) to verify different features and processes but, for provider certification, we require you to run only one test that uses a GitOps flow. Turtles is a project that integrates well with a GitOps approach for cluster provisioning and that is why this is our primary way of validating provider integration with Rancher. Running the full suite for a given CAPI provider will:

- Create a management cluster in the desired environment.
- Install Rancher and Turtles with all prerequisites.
- Install Gitea.
- Run the suite that will create a git repo, apply cluster template using Fleet and verify the cluster is created and successfully imported in Rancher.

### Test configuration

To successfully run the test suite, you will have to provide a number of environment variables. Some of these are agnostic, which means they are required for any provider you want to test, but others will be specific for the provider you are validating. Please, be aware of the particular specifications of the provider being tested, such as credentials, endpoints, etc.

:::tip
Next, we recommend you read the [Test suite guide](../../reference-guides/test-suite/usage.md)
:::

## Request for certification

Integrating with Turtles' test suite and running checks on your provider of interest is enough to validate that it is compatible with Rancher Turtles. As it is not feasible for us to continuously test every CAPI provider, this certification workflow will allow you as a user to verify the expected functionality. However, as we are not actively testing newer iterations of the provider (with newer versions of Turtles), the support and guarantee for the given provider is limited, and you will be responsible for validating future releases.

If, after successfully running checks on your provider, you would like to request it be added to the table of [Certified Providers](../../reference-guides/providers/certified.md), hence being added to the project's periodic E2E suite, you can request so via GitHub issue, using the [Request for Certification template](https://github.com/rancher/turtles/issues/new/choose). The proposal will be studied by the community and we will decide on the feasibility of adding the provider to the certification matrix.
