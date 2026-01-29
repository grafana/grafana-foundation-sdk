---
title: <span class="badge object-type-class"></span> VizConfigKind
---
# <span class="badge object-type-class"></span> VizConfigKind

## Definition

```php
class VizConfigKind implements \JsonSerializable
{
    public string $kind;

    /**
     * The group is the plugin ID
     */
    public string $group;

    public string $version;

    public \Grafana\Foundation\Dashboardv2beta1\VizConfigSpec $spec;

}
```
## Methods

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [annotationslist.VisualizationBuilder](../annotationslist/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [barchart.VisualizationBuilder](../barchart/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [bargauge.VisualizationBuilder](../bargauge/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [candlestick.VisualizationBuilder](../candlestick/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [canvas.VisualizationBuilder](../canvas/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [dashboardlist.VisualizationBuilder](../dashboardlist/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [VizConfigKindBuilder](./builder-VizConfigKindBuilder.md)
 * <span class="badge builder"></span> [datagrid.VisualizationBuilder](../datagrid/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [debug.VisualizationBuilder](../debug/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [gauge.VisualizationBuilder](../gauge/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [geomap.VisualizationBuilder](../geomap/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [heatmap.VisualizationBuilder](../heatmap/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [histogram.VisualizationBuilder](../histogram/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [logs.VisualizationBuilder](../logs/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [logsnew.VisualizationBuilder](../logsnew/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [news.VisualizationBuilder](../news/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [nodegraph.VisualizationBuilder](../nodegraph/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [piechart.VisualizationBuilder](../piechart/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [stat.VisualizationBuilder](../stat/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [statetimeline.VisualizationBuilder](../statetimeline/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [statushistory.VisualizationBuilder](../statushistory/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [table.VisualizationBuilder](../table/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [text.VisualizationBuilder](../text/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [timeseries.VisualizationBuilder](../timeseries/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [trend.VisualizationBuilder](../trend/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [xychart.VisualizationBuilder](../xychart/builder-VisualizationBuilder.md)
