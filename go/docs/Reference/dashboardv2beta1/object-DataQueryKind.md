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
    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    Datasource *dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource `json:"datasource,omitempty"`
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

 * <span class="badge builder"></span> [athena.QueryBuilder](../athena/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [azuremonitor.QueryBuilder](../azuremonitor/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [bigquery.QueryBuilder](../bigquery/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [cloudwatch.QueryBuilder](../cloudwatch/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [datasource.QueryBuilder](../datasource/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [elasticsearch.QueryBuilder](../elasticsearch/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [expr.QueryBuilder](../expr/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [googlecloudmonitoring.QueryBuilder](../googlecloudmonitoring/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [grafanapyroscope.QueryBuilder](../grafanapyroscope/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [loki.QueryBuilder](../loki/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [parca.QueryBuilder](../parca/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [prometheus.QueryBuilder](../prometheus/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [tempo.QueryBuilder](../tempo/builder-QueryBuilder.md)
 * <span class="badge builder"></span> [testdata.QueryBuilder](../testdata/builder-QueryBuilder.md)
