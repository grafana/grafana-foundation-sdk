# Defining a custom query type

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
