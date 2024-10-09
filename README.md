# Grafana Foundation SDK

A set of foundational tools meant to be used to manipulate Grafana resources
as-code.

**_Raw types_** and **_builder libraries_** are provided for a range
of versions of Grafana, in the following languages:

* Go
* PHP
* Python
* Typescript

> [!NOTE]
> The content of this repository is generated by [`cog`][cog] from
> [schemas exposed by Grafana][kind-registry].

> [!TIP]
> This branch contains **types and builders generated for Grafana `next`.**

## Navigating the SDK

The following table can be used to select a version of the SDK suitable for
your Grafana instance.

| Grafana Version                | `cog` Version | Branch |
| ------------------------------ | ------------- | ------ |
| `next` (Grafana's main branch) | `v0.0.x`      | [next+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/next%2Bcog-v0.0.x) |
| `v11.2.x` | `v0.0.x`      | [v11.2.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v11.2.x%2Bcog-v0.0.x) |
| `v11.1.x` | `v0.0.x`      | [v11.1.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v11.1.x%2Bcog-v0.0.x) |
| `v11.0.x` | `v0.0.x`      | [v11.0.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v11.0.x%2Bcog-v0.0.x) |
| `v10.4.x` | `v0.0.x`      | [v10.4.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.4.x%2Bcog-v0.0.x) |
| `v10.3.x` | `v0.0.x`      | [v10.3.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.3.x%2Bcog-v0.0.x) |
| `v10.2.x` | `v0.0.x`      | [v10.2.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.2.x%2Bcog-v0.0.x) |
| `v10.1.x` | `v0.0.x`      | [v10.1.x+cog-v0.0.x](https://github.com/grafana/grafana-foundation-sdk/tree/v10.1.x%2Bcog-v0.0.x) |

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)

[cog]: <https://github.com/grafana/cog>
[kind-registry]: <https://github.com/grafana/kind-registry>
