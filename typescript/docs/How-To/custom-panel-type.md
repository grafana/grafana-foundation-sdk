# Defining a custom panel type

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
