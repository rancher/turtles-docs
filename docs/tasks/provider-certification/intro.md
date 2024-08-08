---
sidebar_position: 1
---

# What is a Certified Provider?

As most CAPI providers are upstream projects maintained by the open community, there is no safe way to guarantee that any clusters provisioned with a given provider can be imported into Rancher via Turtles. However, we do implement a certification status for those providers that:
- Are actively tested as part of our E2E suites.
    - Theses tests are kept up-to-date to validate recent versions of the provider.
- Satisfy the prerequisites in the certification process. 

## Certify your custom provider

Additionally, if you are a provider developer, you can [reuse](../../reference-guides/test-suite/intro.md) the existing Turtles E2E suite to get started with the certification status request or simply verify that Turtles is a viable solution for you to use Rancher and CAPI together. You can read about the certification process and requirements in the [Provider Certification Guide](./process.md)
