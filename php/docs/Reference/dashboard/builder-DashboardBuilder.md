---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```php
new DashboardBuilder(string $title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> annotation

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationQuery> $annotation

```php
annotation(\Grafana\Foundation\Cog\Builder $annotation)
```

### <span class="badge object-method"></span> annotations

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationQuery>> $annotations

```php
annotations(array $annotations)
```

### <span class="badge object-method"></span> description

Description of dashboard.

```php
description(string $description)
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```php
editable()
```

### <span class="badge object-method"></span> fiscalYearStartMonth

The month that the fiscal year starts on.  0 = January, 11 = December

```php
fiscalYearStartMonth(int $fiscalYearStartMonth)
```

### <span class="badge object-method"></span> gnetId

ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal

```php
gnetId(string $gnetId)
```

### <span class="badge object-method"></span> id

Unique numeric identifier for the dashboard.

`id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.

TODO eliminate this null option

```php
id(int $id)
```

### <span class="badge object-method"></span> link

Links with references to other dashboards or external websites.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink> $link

```php
link(\Grafana\Foundation\Cog\Builder $link)
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink>> $links

```php
links(array $links)
```

### <span class="badge object-method"></span> liveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data

```php
liveNow(bool $liveNow)
```

### <span class="badge object-method"></span> preload

When set to true, the dashboard will load all panels in the dashboard when it's loaded.

```php
preload(bool $preload)
```

### <span class="badge object-method"></span> readonly

Whether a dashboard is editable or not.

```php
readonly()
```

### <span class="badge object-method"></span> refresh

Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".

```php
refresh(string $refresh)
```

### <span class="badge object-method"></span> revision

This property should only be used in dashboards defined by plugins.  It is a quick check

to see if the version has changed since the last time.

```php
revision(int $revision)
```

### <span class="badge object-method"></span> snapshot

Snapshot options. They are present only if the dashboard is a snapshot.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Snapshot> $snapshot

```php
snapshot(\Grafana\Foundation\Cog\Builder $snapshot)
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

@param array<string> $tags

```php
tags(array $tags)
```

### <span class="badge object-method"></span> time

Time range for dashboard.

Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.

```php
time(string $from, string $to)
```

### <span class="badge object-method"></span> timepicker

Configuration of the time picker shown at the top of a dashboard.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\TimePickerConfig> $timepicker

```php
timepicker(\Grafana\Foundation\Cog\Builder $timepicker)
```

### <span class="badge object-method"></span> timezone

Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".

```php
timezone(string $timezone)
```

### <span class="badge object-method"></span> title

Title of dashboard.

```php
title(string $title)
```

### <span class="badge object-method"></span> tooltip

Configuration of dashboard cursor sync behavior.

Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).

```php
tooltip(\Grafana\Foundation\Dashboard\DashboardCursorSync $graphTooltip)
```

### <span class="badge object-method"></span> uid

Unique dashboard identifier that can be generated by anyone. string (8-40)

```php
uid(string $uid)
```

### <span class="badge object-method"></span> variables

Configured template variables

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel>> $variables

```php
variables(array $variables)
```

### <span class="badge object-method"></span> version

Version of the dashboard, incremented each time the dashboard is updated.

```php
version(int $version)
```

### <span class="badge object-method"></span> weekStart

Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".

```php
weekStart(string $weekStart)
```

### <span class="badge object-method"></span> withPanel

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Panel> $panel

```php
withPanel(\Grafana\Foundation\Cog\Builder $panel)
```

### <span class="badge object-method"></span> withRow

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\RowPanel> $rowPanel

```php
withRow(\Grafana\Foundation\Cog\Builder $rowPanel)
```

### <span class="badge object-method"></span> withVariable

Configured template variables

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\VariableModel> $variable

```php
withVariable(\Grafana\Foundation\Cog\Builder $variable)
```

## Examples

### Building a dashboard

```php
use Grafana\Foundation\Common;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\RowBuilder;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Timeseries;

require_once __DIR__.'/vendor/autoload.php';

$builder = (new DashboardBuilder(title: 'Sample dashboard'))
    ->uid('generated-from-php')
    ->tags(['generated', 'from', 'php'])
    ->refresh('1m')
    ->time('now-30m', 'now')
    ->timezone(Common\Constants::TIME_ZONE_BROWSER)
    ->withRow(new RowBuilder('Overview'))
    ->withPanel(
        (new Timeseries\PanelBuilder())
            ->title('Network received')
            ->unit('bps')
            ->min(0)
            ->withTarget(
                (new Prometheus\DataqueryBuilder())
                    ->expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
                    ->legendFormat('{{ device }}')
            )
    )
;

echo(json_encode($builder->build(), JSON_PRETTY_PRINT).PHP_EOL);
```
## See also

 * <span class="badge object-type-class"></span> [Dashboard](./object-Dashboard.md)
