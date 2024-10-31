---
title: <span class="badge builder"></span> Dashboard
---
# <span class="badge builder"></span> Dashboard

## Constructor

```python
Dashboard(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.Dashboard
```

### <span class="badge object-method"></span> annotation

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```python
def annotation(annotation: cogbuilder.Builder[dashboard.AnnotationQuery]) -> typing.Self
```

### <span class="badge object-method"></span> annotations

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```python
def annotations(annotations: list[cogbuilder.Builder[dashboard.AnnotationQuery]]) -> typing.Self
```

### <span class="badge object-method"></span> description

Description of dashboard.

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```python
def editable() -> typing.Self
```

### <span class="badge object-method"></span> fiscal_year_start_month

The month that the fiscal year starts on.  0 = January, 11 = December

```python
def fiscal_year_start_month(fiscal_year_start_month: int) -> typing.Self
```

### <span class="badge object-method"></span> gnet_id

ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal

```python
def gnet_id(gnet_id: str) -> typing.Self
```

### <span class="badge object-method"></span> id_val

Unique numeric identifier for the dashboard.

`id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.

```python
def id_val(id_val: int) -> typing.Self
```

### <span class="badge object-method"></span> link

Links with references to other dashboards or external websites.

```python
def link(link: cogbuilder.Builder[dashboard.DashboardLink]) -> typing.Self
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

```python
def links(links: list[cogbuilder.Builder[dashboard.DashboardLink]]) -> typing.Self
```

### <span class="badge object-method"></span> live_now

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data

```python
def live_now(live_now: bool) -> typing.Self
```

### <span class="badge object-method"></span> readonly

Whether a dashboard is editable or not.

```python
def readonly() -> typing.Self
```

### <span class="badge object-method"></span> refresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

```python
def refresh(refresh: str) -> typing.Self
```

### <span class="badge object-method"></span> revision

This property should only be used in dashboards defined by plugins.  It is a quick check

to see if the version has changed since the last time.

```python
def revision(revision: int) -> typing.Self
```

### <span class="badge object-method"></span> snapshot

Snapshot options. They are present only if the dashboard is a snapshot.

```python
def snapshot(snapshot: cogbuilder.Builder[dashboard.Snapshot]) -> typing.Self
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

```python
def tags(tags: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> time

Time range for dashboard.

Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.

```python
def time(from_val: str, to: str) -> typing.Self
```

### <span class="badge object-method"></span> timepicker

Configuration of the time picker shown at the top of a dashboard.

```python
def timepicker(timepicker: cogbuilder.Builder[dashboard.TimePickerConfig]) -> typing.Self
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```python
def timezone(timezone: str) -> typing.Self
```

### <span class="badge object-method"></span> title

Title of dashboard.

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

Configuration of dashboard cursor sync behavior.

Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).

```python
def tooltip(graph_tooltip: dashboard.DashboardCursorSync) -> typing.Self
```

### <span class="badge object-method"></span> uid

Unique dashboard identifier that can be generated by anyone. string (8-40)

```python
def uid(uid: str) -> typing.Self
```

### <span class="badge object-method"></span> variables

Configured template variables

```python
def variables(variables: list[cogbuilder.Builder[dashboard.VariableModel]]) -> typing.Self
```

### <span class="badge object-method"></span> version

Version of the dashboard, incremented each time the dashboard is updated.

```python
def version(version: int) -> typing.Self
```

### <span class="badge object-method"></span> week_start

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```python
def week_start(week_start: str) -> typing.Self
```

### <span class="badge object-method"></span> with_panel

```python
def with_panel(panel: cogbuilder.Builder[dashboard.Panel]) -> typing.Self
```

### <span class="badge object-method"></span> with_row

```python
def with_row(row_panel: cogbuilder.Builder[dashboard.RowPanel]) -> typing.Self
```

### <span class="badge object-method"></span> with_variable

Configured template variables

```python
def with_variable(variable: cogbuilder.Builder[dashboard.VariableModel]) -> typing.Self
```

## Examples

### Building a dashboard

```python
from grafana_foundation_sdk.builders.dashboard import Dashboard, Row
from grafana_foundation_sdk.builders.prometheus import Dataquery as PrometheusQuery
from grafana_foundation_sdk.builders.timeseries import Panel as Timeseries
from grafana_foundation_sdk.models.common import TimeZoneBrowser

def build_dashboard() -> Dashboard:
    return (
        Dashboard("[TEST] Node Exporter / Raspberry")
        .uid("test-dashboard-raspberry")
        .tags(["generated", "raspberrypi-node-integration"])

        .refresh("1m")
        .time("now-30m", "now")
        .timezone(TimeZoneBrowser)

        .with_row(Row("Overview"))
        .with_panel(
            Timeseries()
            .title("Network Received")
            .unit("bps")
            .min_val(0)
            .with_target(
                PrometheusQuery()
                .expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
                .legend_format("{{ device }}")
            )
        )
    )
```
## See also

 * <span class="badge object-type-class"></span> [Dashboard](./object-Dashboard.md)
