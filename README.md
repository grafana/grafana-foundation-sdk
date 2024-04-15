# Grafana Foundation SDK

A set of foundational tools meant to be used to manipulate Grafana resources
as-code.

**_Raw types_** and **_builder libraries_** are provided for a range
of versions of Grafana, in the following languages:

* Go
* Python
* Typescript

> [!NOTE]
> The content of this repository is generated by [`cog`][cog] from
> [schemas exposed by Grafana][kind-registry].

> [!TIP]
> This branch contains **types and builders generated for Grafana `main`.**

## Navigating the SDK

The following table can be used to select a version of the SDK suitable for
your Grafana instance.

| Grafana Version                | `cog` Version | Branch |
| ------------------------------ | ------------- | ------ |
| `next` (Grafana's main branch) | `v0.0.x`      | [next+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/next%2Bcog-v0.0.x) |
| `v10.4.x`                      | `v0.0.x`      | [v10.4.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.4.x%2Bcog-v0.0.x) |
| `v10.3.x`                      | `v0.0.x`      | [v10.3.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.3.x%2Bcog-v0.0.x) |
| `v10.2.x`                      | `v0.0.x`      | [v10.2.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.2.x%2Bcog-v0.0.x) |
| `v10.1.x`                      | `v0.0.x`      | [v10.1.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.1.x%2Bcog-v0.0.x) |

## Maturity

> [!WARNING]
> The code in this repository should be considered experimental. Documentation is only
available alongside the code. It comes with no support, but we are keen to receive
feedback and suggestions on how to improve it, though we cannot commit
to resolution of any particular issue.

Grafana Labs defines experimental features as follows:

> Projects and features in the Experimental stage are supported only by the Engineering
teams; on-call support is not available. Documentation is either limited or not provided
outside of code comments. No SLA is provided.
>
> Experimental projects or features are primarily intended for open source engineers who
want to participate in ensuring systems stability, and to gain consensus and approval
for open source governance projects.
>
> Projects and features in the Experimental phase are not meant to be used in production
environments, and the risks are unknown/high.

## License

[Apache 2.0 License](./LICENSE)

[cog]: <https://github.com/grafana/cog>
[kind-registry]: <https://github.com/grafana/kind-registry>
