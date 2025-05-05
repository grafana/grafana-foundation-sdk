---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```java
new DashboardBuilder(String title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dashboard build()
```

### <span class="badge object-method"></span> annotation

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```java
public DashboardBuilder annotation(com.grafana.foundation.cog.Builder<AnnotationQuery> annotation)
```

### <span class="badge object-method"></span> annotations

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

```java
public DashboardBuilder annotations(List<com.grafana.foundation.cog.Builder<AnnotationQuery>> annotations)
```

### <span class="badge object-method"></span> description

Description of dashboard.

```java
public DashboardBuilder description(String description)
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```java
public DashboardBuilder editable()
```

### <span class="badge object-method"></span> fiscalYearStartMonth

The month that the fiscal year starts on.  0 = January, 11 = December

```java
public DashboardBuilder fiscalYearStartMonth(Integer fiscalYearStartMonth)
```

### <span class="badge object-method"></span> gnetId

ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal

```java
public DashboardBuilder gnetId(String gnetId)
```

### <span class="badge object-method"></span> id

Unique numeric identifier for the dashboard.

`id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.

```java
public DashboardBuilder id(Long id)
```

### <span class="badge object-method"></span> link

Links with references to other dashboards or external websites.

```java
public DashboardBuilder link(com.grafana.foundation.cog.Builder<DashboardLink> link)
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

```java
public DashboardBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links)
```

### <span class="badge object-method"></span> liveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data

```java
public DashboardBuilder liveNow(Boolean liveNow)
```

### <span class="badge object-method"></span> readonly

Whether a dashboard is editable or not.

```java
public DashboardBuilder readonly()
```

### <span class="badge object-method"></span> refresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

```java
public DashboardBuilder refresh(String refresh)
```

### <span class="badge object-method"></span> revision

This property should only be used in dashboards defined by plugins.  It is a quick check

to see if the version has changed since the last time.

```java
public DashboardBuilder revision(Long revision)
```

### <span class="badge object-method"></span> snapshot

Snapshot options. They are present only if the dashboard is a snapshot.

```java
public DashboardBuilder snapshot(com.grafana.foundation.cog.Builder<Snapshot> snapshot)
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

```java
public DashboardBuilder tags(List<String> tags)
```

### <span class="badge object-method"></span> time

Time range for dashboard.

Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.

```java
public DashboardBuilder time(com.grafana.foundation.cog.Builder<DashboardDashboardTime> time)
```

### <span class="badge object-method"></span> timepicker

Configuration of the time picker shown at the top of a dashboard.

```java
public DashboardBuilder timepicker(com.grafana.foundation.cog.Builder<TimePickerConfig> timepicker)
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```java
public DashboardBuilder timezone(String timezone)
```

### <span class="badge object-method"></span> title

Title of dashboard.

```java
public DashboardBuilder title(String title)
```

### <span class="badge object-method"></span> tooltip

Configuration of dashboard cursor sync behavior.

Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).

```java
public DashboardBuilder tooltip(DashboardCursorSync graphTooltip)
```

### <span class="badge object-method"></span> uid

Unique dashboard identifier that can be generated by anyone. string (8-40)

```java
public DashboardBuilder uid(String uid)
```

### <span class="badge object-method"></span> variables

Configured template variables

```java
public DashboardBuilder variables(List<com.grafana.foundation.cog.Builder<VariableModel>> variables)
```

### <span class="badge object-method"></span> version

Version of the dashboard, incremented each time the dashboard is updated.

```java
public DashboardBuilder version(Integer version)
```

### <span class="badge object-method"></span> weekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```java
public DashboardBuilder weekStart(String weekStart)
```

### <span class="badge object-method"></span> withPanel

```java
public DashboardBuilder withPanel(com.grafana.foundation.cog.Builder<Panel> panel)
```

### <span class="badge object-method"></span> withRow

```java
public DashboardBuilder withRow(com.grafana.foundation.cog.Builder<RowPanel> rowPanel)
```

### <span class="badge object-method"></span> withVariable

Configured template variables

```java
public DashboardBuilder withVariable(com.grafana.foundation.cog.Builder<VariableModel> variable)
```

## Examples

### Building a dashboard

```java
package example;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.dashboard.Dashboard;
import com.grafana.foundation.common.Constants;
import com.grafana.foundation.dashboard.DashboardBuilder;
import com.grafana.foundation.dashboard.DashboardDashboardTimeBuilder;
import com.grafana.foundation.dashboard.RowBuilder;
import com.grafana.foundation.prometheus.DataqueryBuilder;
import com.grafana.foundation.timeseries.PanelBuilder;

import java.util.List;

public class Main {
    public static void main(String[] args) throws JsonProcessingException {
        Dashboard dashboard = new DashboardBuilder("Sample Dashboard").
                uid("generated-from-java").
                tags(List.of("generated", "from", "java")).
                refresh("1m").
                time(new DashboardDashboardTimeBuilder().
                        from("now-30m").
                        to("now")
                ).
                timezone(Constants.TimeZoneBrowser).
                withRow(new RowBuilder("Overview")).
                withPanel(new PanelBuilder().
                        title("Network Received").
                        unit("bps").
                        min(0.0).
                        withTarget(new DataqueryBuilder().
                                expr("rate(node_network_receive_bytes_total{job=\"integrations/raspberrypi-node\", device!=\"lo\"}[$__rate_interval]) * 8").
                                legendFormat("{{ device }}")
                        )
                ).build();

        System.out.println(dashboard.toJSON());
    }
}
```
## See also

 * <span class="badge object-type-class"></span> [Dashboard](./object-Dashboard.md)
