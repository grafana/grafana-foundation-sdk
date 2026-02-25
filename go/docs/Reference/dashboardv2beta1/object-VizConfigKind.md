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
    Spec dashboardv2beta1.VizConfigSpec `json:"spec"`
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

 * <span class="badge builder"></span> [annotationslist.VisualizationBuilder](../annotationslist/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [barchart.VisualizationBuilder](../barchart/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [bargauge.VisualizationBuilder](../bargauge/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [candlestick.VisualizationBuilder](../candlestick/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [canvas.VisualizationBuilder](../canvas/builder-VisualizationBuilder.md)
 * <span class="badge builder"></span> [dashboardlist.VisualizationBuilder](../dashboardlist/builder-VisualizationBuilder.md)
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
