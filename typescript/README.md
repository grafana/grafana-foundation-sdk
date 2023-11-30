# Grafana Foundation SDK – TypeScript

A set of tools, types and libraries for building and manipulating Grafana objects in TypeScript.

## Maturity

> _The code in this repository should be considered experimental. Documentation is only
available alongside the code. It comes with no support, but we are keen to receive
feedback on the product and suggestions on how to improve it, though we cannot commit
to resolution of any particular issue. No SLAs are available. It is not meant to be used
in production environments, and the risks are unknown/high._

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

## Installing

```shell
yarn add @grafana/grafana-foundation-sdk
```

## Example usage

```typescript
import { DashboardBuilder, TimePickerBuilder } from "@grafana/grafana-foundation-sdk/dashboard";

const builder = new DashboardBuilder("Sample dashboard")
    .uid("generated-from-typescript")
    .tags(["generated", "from", "typescript"])

    .refresh("30s")
    .time({from: "now-30m", to: "now"})
    .timezone("browser")

    .timepicker(
        new TimePickerBuilder()
            .refresh_intervals(["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"])
            .time_options(["5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"]),
    )
;

console.log(JSON.stringify(builder.build(), null, 2));
```

## License

[Apache 2.0 License](./LICENSE)
