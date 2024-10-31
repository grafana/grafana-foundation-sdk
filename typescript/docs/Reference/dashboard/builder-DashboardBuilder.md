---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```typescript
new DashboardBuilder(title: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> annotation

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```typescript
annotation(annotation: cog.Builder<dashboard.AnnotationQuery>)
```

### <span class="badge object-method"></span> annotations

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```typescript
annotations(annotations: cog.Builder<dashboard.AnnotationQuery>[])
```

### <span class="badge object-method"></span> description

Description of dashboard.

```typescript
description(description: string)
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```typescript
editable()
```

### <span class="badge object-method"></span> fiscalYearStartMonth

The month that the fiscal year starts on.  0 = January, 11 = December

```typescript
fiscalYearStartMonth(fiscalYearStartMonth: number)
```

### <span class="badge object-method"></span> gnetId

ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal

```typescript
gnetId(gnetId: string)
```

### <span class="badge object-method"></span> id

Unique numeric identifier for the dashboard.

`id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.

```typescript
id(id: number | null)
```

### <span class="badge object-method"></span> link

Links with references to other dashboards or external websites.

```typescript
link(link: cog.Builder<dashboard.DashboardLink>)
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

```typescript
links(links: cog.Builder<dashboard.DashboardLink>[])
```

### <span class="badge object-method"></span> liveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data

```typescript
liveNow(liveNow: boolean)
```

### <span class="badge object-method"></span> readonly

Whether a dashboard is editable or not.

```typescript
readonly()
```

### <span class="badge object-method"></span> refresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

```typescript
refresh(refresh: string | false)
```

### <span class="badge object-method"></span> revision

This property should only be used in dashboards defined by plugins.  It is a quick check

to see if the version has changed since the last time.

```typescript
revision(revision: number)
```

### <span class="badge object-method"></span> snapshot

Snapshot options. They are present only if the dashboard is a snapshot.

```typescript
snapshot(snapshot: cog.Builder<dashboard.Snapshot>)
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

```typescript
tags(tags: string[])
```

### <span class="badge object-method"></span> time

Time range for dashboard.

Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.

```typescript
time(time: {
	from: string;
	to: string;
})
```

### <span class="badge object-method"></span> timepicker

Configuration of the time picker shown at the top of a dashboard.

```typescript
timepicker(timepicker: cog.Builder<dashboard.TimePickerConfig>)
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```typescript
timezone(timezone: string)
```

### <span class="badge object-method"></span> title

Title of dashboard.

```typescript
title(title: string)
```

### <span class="badge object-method"></span> tooltip

Configuration of dashboard cursor sync behavior.

Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).

```typescript
tooltip(graphTooltip: dashboard.DashboardCursorSync)
```

### <span class="badge object-method"></span> uid

Unique dashboard identifier that can be generated by anyone. string (8-40)

```typescript
uid(uid: string)
```

### <span class="badge object-method"></span> variables

Configured template variables

```typescript
variables(variables: cog.Builder<dashboard.VariableModel>[])
```

### <span class="badge object-method"></span> version

Version of the dashboard, incremented each time the dashboard is updated.

```typescript
version(version: number)
```

### <span class="badge object-method"></span> weekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```typescript
weekStart(weekStart: string)
```

### <span class="badge object-method"></span> withPanel

```typescript
withPanel(panel: cog.Builder<dashboard.Panel>)
```

### <span class="badge object-method"></span> withRow

```typescript
withRow(rowPanel: cog.Builder<dashboard.RowPanel>)
```

### <span class="badge object-method"></span> withVariable

Configured template variables

```typescript
withVariable(variable: cog.Builder<dashboard.VariableModel>)
```

## Examples

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
          .legendFormat("{{ device }}")
      )
  )
;

console.log(JSON.stringify(builder.build(), null, 2));
```
## See also

 * <span class="badge object-type-interface"></span> [Dashboard](./object-Dashboard.md)
