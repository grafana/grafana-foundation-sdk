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
    spec: dashboardv2beta1.VizConfigSpec
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

 * <span class="badge builder"></span> [annotationslist.Visualization](../annotationslist/builder-Visualization.md)
 * <span class="badge builder"></span> [barchart.Visualization](../barchart/builder-Visualization.md)
 * <span class="badge builder"></span> [bargauge.Visualization](../bargauge/builder-Visualization.md)
 * <span class="badge builder"></span> [candlestick.Visualization](../candlestick/builder-Visualization.md)
 * <span class="badge builder"></span> [canvas.Visualization](../canvas/builder-Visualization.md)
 * <span class="badge builder"></span> [dashboardlist.Visualization](../dashboardlist/builder-Visualization.md)
 * <span class="badge builder"></span> [VizConfigKind](./builder-VizConfigKind.md)
 * <span class="badge builder"></span> [datagrid.Visualization](../datagrid/builder-Visualization.md)
 * <span class="badge builder"></span> [debug.Visualization](../debug/builder-Visualization.md)
 * <span class="badge builder"></span> [gauge.Visualization](../gauge/builder-Visualization.md)
 * <span class="badge builder"></span> [geomap.Visualization](../geomap/builder-Visualization.md)
 * <span class="badge builder"></span> [heatmap.Visualization](../heatmap/builder-Visualization.md)
 * <span class="badge builder"></span> [histogram.Visualization](../histogram/builder-Visualization.md)
 * <span class="badge builder"></span> [logs.Visualization](../logs/builder-Visualization.md)
 * <span class="badge builder"></span> [logsnew.Visualization](../logsnew/builder-Visualization.md)
 * <span class="badge builder"></span> [news.Visualization](../news/builder-Visualization.md)
 * <span class="badge builder"></span> [nodegraph.Visualization](../nodegraph/builder-Visualization.md)
 * <span class="badge builder"></span> [piechart.Visualization](../piechart/builder-Visualization.md)
 * <span class="badge builder"></span> [stat.Visualization](../stat/builder-Visualization.md)
 * <span class="badge builder"></span> [statetimeline.Visualization](../statetimeline/builder-Visualization.md)
 * <span class="badge builder"></span> [statushistory.Visualization](../statushistory/builder-Visualization.md)
 * <span class="badge builder"></span> [table.Visualization](../table/builder-Visualization.md)
 * <span class="badge builder"></span> [text.Visualization](../text/builder-Visualization.md)
 * <span class="badge builder"></span> [timeseries.Visualization](../timeseries/builder-Visualization.md)
 * <span class="badge builder"></span> [trend.Visualization](../trend/builder-Visualization.md)
 * <span class="badge builder"></span> [xychart.Visualization](../xychart/builder-Visualization.md)
