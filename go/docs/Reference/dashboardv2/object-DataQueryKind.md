---
title: <span class="badge object-type-struct"></span> DataQueryKind
---
# <span class="badge object-type-struct"></span> DataQueryKind

## Definition

```go
type DataQueryKind struct {
    Kind string `json:"kind"`
    Group string `json:"group"`
    Version string `json:"version"`
    Labels map[string]string `json:"labels,omitempty"`
    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    Datasource *dashboardv2.Dashboardv2DataQueryKindDatasource `json:"datasource,omitempty"`
    Spec any `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DataQueryKind` from JSON.

```go
func (dataQueryKind *DataQueryKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataQueryKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataQueryKind *DataQueryKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DataQueryKind` objects.

```go
func (dataQueryKind *DataQueryKind) Equals(other DataQueryKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DataQueryKind` fields for violations and returns them.

```go
func (dataQueryKind *DataQueryKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [athena.QueryV2Builder](../athena/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [azuremonitor.QueryV2Builder](../azuremonitor/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [bigquery.QueryV2Builder](../bigquery/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [cloudwatch.QueryV2Builder](../cloudwatch/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [datasource.QueryV2Builder](../datasource/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [elasticsearch.QueryV2Builder](../elasticsearch/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [expr.QueryV2Builder](../expr/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [googlecloudmonitoring.QueryV2Builder](../googlecloudmonitoring/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [grafanapyroscope.QueryV2Builder](../grafanapyroscope/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [loki.QueryV2Builder](../loki/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [parca.QueryV2Builder](../parca/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [prometheus.QueryV2Builder](../prometheus/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [tempo.QueryV2Builder](../tempo/builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [testdata.QueryV2Builder](../testdata/builder-QueryV2Builder.md)
