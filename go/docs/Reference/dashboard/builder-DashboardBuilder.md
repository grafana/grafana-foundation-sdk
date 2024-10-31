---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```go
func NewDashboardBuilder(title string) *DashboardBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardBuilder) Build() (Dashboard, error)
```

### <span class="badge object-method"></span> Annotation

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```go
func (builder *DashboardBuilder) Annotation(annotation cog.Builder[dashboard.AnnotationQuery]) *DashboardBuilder
```

### <span class="badge object-method"></span> Annotations

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```go
func (builder *DashboardBuilder) Annotations(annotations []cog.Builder[dashboard.AnnotationQuery]) *DashboardBuilder
```

### <span class="badge object-method"></span> Description

Description of dashboard.

```go
func (builder *DashboardBuilder) Description(description string) *DashboardBuilder
```

### <span class="badge object-method"></span> Editable

Whether a dashboard is editable or not.

```go
func (builder *DashboardBuilder) Editable() *DashboardBuilder
```

### <span class="badge object-method"></span> FiscalYearStartMonth

The month that the fiscal year starts on.  0 = January, 11 = December

```go
func (builder *DashboardBuilder) FiscalYearStartMonth(fiscalYearStartMonth uint8) *DashboardBuilder
```

### <span class="badge object-method"></span> GnetId

ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal

```go
func (builder *DashboardBuilder) GnetId(gnetId string) *DashboardBuilder
```

### <span class="badge object-method"></span> Id

Unique numeric identifier for the dashboard.

`id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.

```go
func (builder *DashboardBuilder) Id(id int64) *DashboardBuilder
```

### <span class="badge object-method"></span> Link

Links with references to other dashboards or external websites.

```go
func (builder *DashboardBuilder) Link(link cog.Builder[dashboard.DashboardLink]) *DashboardBuilder
```

### <span class="badge object-method"></span> Links

Links with references to other dashboards or external websites.

```go
func (builder *DashboardBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *DashboardBuilder
```

### <span class="badge object-method"></span> LiveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data

```go
func (builder *DashboardBuilder) LiveNow(liveNow bool) *DashboardBuilder
```

### <span class="badge object-method"></span> Readonly

Whether a dashboard is editable or not.

```go
func (builder *DashboardBuilder) Readonly() *DashboardBuilder
```

### <span class="badge object-method"></span> Refresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

```go
func (builder *DashboardBuilder) Refresh(refresh string) *DashboardBuilder
```

### <span class="badge object-method"></span> Revision

This property should only be used in dashboards defined by plugins.  It is a quick check

to see if the version has changed since the last time.

```go
func (builder *DashboardBuilder) Revision(revision int64) *DashboardBuilder
```

### <span class="badge object-method"></span> Snapshot

Snapshot options. They are present only if the dashboard is a snapshot.

```go
func (builder *DashboardBuilder) Snapshot(snapshot cog.Builder[dashboard.Snapshot]) *DashboardBuilder
```

### <span class="badge object-method"></span> Tags

Tags associated with dashboard.

```go
func (builder *DashboardBuilder) Tags(tags []string) *DashboardBuilder
```

### <span class="badge object-method"></span> Time

Time range for dashboard.

Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.

```go
func (builder *DashboardBuilder) Time(from string, to string) *DashboardBuilder
```

### <span class="badge object-method"></span> Timepicker

Configuration of the time picker shown at the top of a dashboard.

```go
func (builder *DashboardBuilder) Timepicker(timepicker cog.Builder[dashboard.TimePickerConfig]) *DashboardBuilder
```

### <span class="badge object-method"></span> Timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```go
func (builder *DashboardBuilder) Timezone(timezone string) *DashboardBuilder
```

### <span class="badge object-method"></span> Title

Title of dashboard.

```go
func (builder *DashboardBuilder) Title(title string) *DashboardBuilder
```

### <span class="badge object-method"></span> Tooltip

Configuration of dashboard cursor sync behavior.

Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).

```go
func (builder *DashboardBuilder) Tooltip(graphTooltip dashboard.DashboardCursorSync) *DashboardBuilder
```

### <span class="badge object-method"></span> Uid

Unique dashboard identifier that can be generated by anyone. string (8-40)

```go
func (builder *DashboardBuilder) Uid(uid string) *DashboardBuilder
```

### <span class="badge object-method"></span> Variables

Configured template variables

```go
func (builder *DashboardBuilder) Variables(variables []cog.Builder[dashboard.VariableModel]) *DashboardBuilder
```

### <span class="badge object-method"></span> Version

Version of the dashboard, incremented each time the dashboard is updated.

```go
func (builder *DashboardBuilder) Version(version uint32) *DashboardBuilder
```

### <span class="badge object-method"></span> WeekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```go
func (builder *DashboardBuilder) WeekStart(weekStart string) *DashboardBuilder
```

### <span class="badge object-method"></span> WithPanel

```go
func (builder *DashboardBuilder) WithPanel(panel cog.Builder[dashboard.Panel]) *DashboardBuilder
```

### <span class="badge object-method"></span> WithRow

```go
func (builder *DashboardBuilder) WithRow(rowPanel cog.Builder[dashboard.RowPanel]) *DashboardBuilder
```

### <span class="badge object-method"></span> WithVariable

Configured template variables

```go
func (builder *DashboardBuilder) WithVariable(variable cog.Builder[dashboard.VariableModel]) *DashboardBuilder
```

## Examples

### Building a dashboard

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/grafana/grafana-foundation-sdk/go/common"
    "github.com/grafana/grafana-foundation-sdk/go/dashboard"
    "github.com/grafana/grafana-foundation-sdk/go/prometheus"
    "github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func main() {
    builder := dashboard.NewDashboardBuilder("Sample dashboard").
        Uid("generated-from-go").
        Tags([]string{"generated", "from", "go"}).
        Refresh("1m").
        Time("now-30m", "now").
        Timezone(common.TimeZoneBrowser).
        WithRow(dashboard.NewRowBuilder("Overview")).
        WithPanel(
            timeseries.NewPanelBuilder().
                Title("Network Received").
                Unit("bps").
                Min(0).
                WithTarget(
                    prometheus.NewDataqueryBuilder().
                        Expr(`rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8`).
                        LegendFormat("{{ device }}"),
                ),
        )

    sampleDashboard, err := builder.Build()
    if err != nil {
        panic(err)
    }
    dashboardJson, err := json.MarshalIndent(sampleDashboard, "", "  ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(dashboardJson))
}
```
## See also

 * <span class="badge object-type-struct"></span> [Dashboard](./object-Dashboard.md)
