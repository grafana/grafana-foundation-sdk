# Grafana Foundation SDK – TypeScript

A set of tools, types and *builder libraries* for building and manipulating Grafana objects in TypeScript.

> [!NOTE]
> This branch contains **types and builders generated for Grafana {{ .Extra.GrafanaVersion }}.**
> Other supported versions of Grafana can be found at [this repository's root](https://github.com/grafana/grafana-foundation-sdk/).

## Installing

```shell
yarn add '@grafana/grafana-foundation-sdk@~{{ .Extra.GrafanaVersion|registryToSemver }}-cog{{ .Extra.CogVersion }}.{{ .Extra.BuildTimestamp }}'
```

## Example usage

[More examples](https://github.com/grafana/grafana-foundation-sdk/tree/main/examples/typescript) can be found at the repository root.

### Building a dashboard

```typescript
import { DashboardBuilder, RowBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { DataqueryBuilder } from '@grafana/grafana-foundation-sdk/prometheus';
import { PanelBuilder } from '@grafana/grafana-foundation-sdk/timeseries';

const builder = new DashboardBuilder('[TEST] Node Exporter / Raspberry')
  .uid('test-dashboard-raspberry')
  .tags(['generated', 'raspberrypi-node-integration'])

  .refresh('1m')
  .time({from: 'now-30m', to: 'now'})
  .timezone('browser')

  .withRow(new RowBuilder('Overview'))
  .withPanel(
    new PanelBuilder()
      .title('Network Received')
      .unit('bps')
      .min(0)
      .withTarget(
        new DataqueryBuilder()
          .expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
          .legendFormat({{ `"{{ device }}"` }})
      )
  )
;

console.log(JSON.stringify(builder.build(), null, 2));
```

### Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type and a builder for the custom query:

```typescript
// customQuery.ts
import { Builder, Dataquery } from '@grafana/grafana-foundation-sdk/cog';

export interface CustomQuery {
    // refId and hide are expected on all queries
    refId?: string;
    hide?: boolean;


    // query is specific to the CustomQuery type
    query: string;

    // Let cog know that CustomQuery is a Dataquery variant
    _implementsDataqueryVariant(): void;
}

export const defaultCustomQuery = (): CustomQuery => ({
    query: "",
    _implementsDataqueryVariant() {},
});

export class CustomQueryBuilder implements Builder<Dataquery> {
    private readonly internal: CustomQuery;

    constructor(query: string) {
        this.internal = defaultCustomQuery();
        this.internal.query = query;
    }

    build(): CustomQuery {
        return this.internal;
    }

    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
```

The custom query type can now be used as usual to build a dashboard:

```typescript
import { DashboardBuilder, RowBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { PanelBuilder as TimeSeriesBuilder } from "@grafana/grafana-foundation-sdk/timeseries";
import { CustomQueryBuilder } from "./customQuery";

const builder = new DashboardBuilder('Custom query type')
    .uid('test-custom-query-type')

    .refresh('1m')
    .time({ from: 'now-30m', to: 'now' })

    .withRow(new RowBuilder('Overview'))
    .withPanel(
        new TimeSeriesBuilder()
            .title('Sample panel')
            .withTarget(
                new CustomQueryBuilder("query here")
            )
    );

console.log(JSON.stringify(builder.build(), null, 2));
```

### Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type and a builder for the custom panel's options:

```typescript
// customPanel.ts
import * as dashboard from '@grafana/grafana-foundation-sdk/dashboard';

export interface CustomPanelOptions {
    makeBeautiful?: boolean;
}

export const defaultCustomPanelOptions = (): CustomPanelOptions => ({
});

export class CustomPanelBuilder extends dashboard.PanelBuilder {
    constructor() {
        super();
        this.internal.type = "custom-panel"; // panel plugin ID
    }

    makeBeautiful(): this {
        if (!this.internal.options) {
            this.internal.options = defaultCustomPanelOptions();
        }
        this.internal.options.makeBeautiful = true;
        return this;
    }
}
```

The custom panel type can now be used as usual to build a dashboard:

```typescript
import { DashboardBuilder, RowBuilder } from '@grafana/grafana-foundation-sdk/dashboard';
import { CustomPanelBuilder } from "./customPanel";

const builder = new DashboardBuilder('Custom panel type')
    .uid('test-custom-panel-type')

    .refresh('1m')
    .time({ from: 'now-30m', to: 'now' })

    .withRow(new RowBuilder('Overview'))
    .withPanel(
        new CustomPanelBuilder()
            .title('Sample custom panel')
            .makeBeautiful()
    );

console.log(JSON.stringify(builder.build(), null, 2));
```

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)
