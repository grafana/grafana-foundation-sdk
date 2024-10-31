---
title: <span class="badge object-type-struct"></span> SLOQuery
---
# <span class="badge object-type-struct"></span> SLOQuery

SLO sub-query properties.

## Definition

```go
type SLOQuery struct {
    // GCP project to execute the query against.
    ProjectName string `json:"projectName"`
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`
    // SLO selector.
    SelectorName string `json:"selectorName"`
    // ID for the service the SLO is in.
    ServiceId string `json:"serviceId"`
    // Name for the service the SLO is in.
    ServiceName string `json:"serviceName"`
    // ID for the SLO.
    SloId string `json:"sloId"`
    // Name of the SLO.
    SloName string `json:"sloName"`
    // SLO goal value.
    Goal *float64 `json:"goal,omitempty"`
    // Specific lookback period for the SLO.
    LookbackPeriod *string `json:"lookbackPeriod,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SLOQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (sLOQuery *SLOQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SLOQuery` objects.

```go
func (sLOQuery *SLOQuery) Equals(other SLOQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SLOQuery` fields for violations and returns them.

```go
func (sLOQuery *SLOQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [SLOQueryBuilder](./builder-SLOQueryBuilder.md)
