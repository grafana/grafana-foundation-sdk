---
title: <span class="badge object-type-struct"></span> VizConfigKind
---
# <span class="badge object-type-struct"></span> VizConfigKind

## Definition

```go
type VizConfigKind struct {
    Kind string `json:"kind"`
    // The group is the plugin ID
    Group string `json:"group"`
    Version string `json:"version"`
    Spec dashboardv2.VizConfigSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `VizConfigKind` from JSON.

```go
func (vizConfigKind *VizConfigKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `VizConfigKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (vizConfigKind *VizConfigKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `VizConfigKind` objects.

```go
func (vizConfigKind *VizConfigKind) Equals(other VizConfigKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VizConfigKind` fields for violations and returns them.

```go
func (vizConfigKind *VizConfigKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [annotationslist.VisualizationV2Builder](../annotationslist/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [barchart.VisualizationV2Builder](../barchart/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [bargauge.VisualizationV2Builder](../bargauge/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [candlestick.VisualizationV2Builder](../candlestick/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [canvas.VisualizationV2Builder](../canvas/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [dashboardlist.VisualizationV2Builder](../dashboardlist/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [VizConfigKindBuilder](./builder-VizConfigKindBuilder.md)
 * <span class="badge builder"></span> [datagrid.VisualizationV2Builder](../datagrid/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [debug.VisualizationV2Builder](../debug/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [gauge.VisualizationV2Builder](../gauge/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [geomap.VisualizationV2Builder](../geomap/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [heatmap.VisualizationV2Builder](../heatmap/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [histogram.VisualizationV2Builder](../histogram/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [logs.VisualizationV2Builder](../logs/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [news.VisualizationV2Builder](../news/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [nodegraph.VisualizationV2Builder](../nodegraph/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [piechart.VisualizationV2Builder](../piechart/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [stat.VisualizationV2Builder](../stat/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [statetimeline.VisualizationV2Builder](../statetimeline/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [statushistory.VisualizationV2Builder](../statushistory/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [table.VisualizationV2Builder](../table/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [text.VisualizationV2Builder](../text/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [timeseries.VisualizationV2Builder](../timeseries/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [trend.VisualizationV2Builder](../trend/builder-VisualizationV2Builder.md)
 * <span class="badge builder"></span> [xychart.VisualizationV2Builder](../xychart/builder-VisualizationV2Builder.md)
