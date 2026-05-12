---
title: <span class="badge object-type-class"></span> VizConfigKind
---
# <span class="badge object-type-class"></span> VizConfigKind

## Definition

```python
class VizConfigKind:
    kind: typing.Literal["VizConfig"]
    # The group is the plugin ID
    group: str
    version: str
    spec: dashboardv2.VizConfigSpec
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [annotationslist.VisualizationV2](../annotationslist/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [barchart.VisualizationV2](../barchart/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [bargauge.VisualizationV2](../bargauge/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [candlestick.VisualizationV2](../candlestick/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [canvas.VisualizationV2](../canvas/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [dashboardlist.VisualizationV2](../dashboardlist/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [VizConfigKind](./builder-VizConfigKind.md)
 * <span class="badge builder"></span> [datagrid.VisualizationV2](../datagrid/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [debug.VisualizationV2](../debug/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [gauge.VisualizationV2](../gauge/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [geomap.VisualizationV2](../geomap/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [heatmap.VisualizationV2](../heatmap/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [histogram.VisualizationV2](../histogram/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [logs.VisualizationV2](../logs/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [news.VisualizationV2](../news/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [nodegraph.VisualizationV2](../nodegraph/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [piechart.VisualizationV2](../piechart/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [stat.VisualizationV2](../stat/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [statetimeline.VisualizationV2](../statetimeline/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [statushistory.VisualizationV2](../statushistory/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [table.VisualizationV2](../table/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [text.VisualizationV2](../text/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [timeseries.VisualizationV2](../timeseries/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [trend.VisualizationV2](../trend/builder-VisualizationV2.md)
 * <span class="badge builder"></span> [xychart.VisualizationV2](../xychart/builder-VisualizationV2.md)
