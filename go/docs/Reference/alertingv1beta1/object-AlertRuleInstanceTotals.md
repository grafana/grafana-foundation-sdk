---
title: <span class="badge object-type-struct"></span> AlertRuleInstanceTotals
---
# <span class="badge object-type-struct"></span> AlertRuleInstanceTotals

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Definition

```go
type AlertRuleInstanceTotals struct {
    Healthy *int64 `json:"healthy,omitempty"`
    Firing *int64 `json:"firing,omitempty"`
    Pending *int64 `json:"pending,omitempty"`
    Recovering *int64 `json:"recovering,omitempty"`
    Nodata *int64 `json:"nodata,omitempty"`
    Error *int64 `json:"error,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AlertRuleInstanceTotals` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (alertRuleInstanceTotals *AlertRuleInstanceTotals) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AlertRuleInstanceTotals` objects.

```go
func (alertRuleInstanceTotals *AlertRuleInstanceTotals) Equals(other AlertRuleInstanceTotals) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AlertRuleInstanceTotals` fields for violations and returns them.

```go
func (alertRuleInstanceTotals *AlertRuleInstanceTotals) Validate() error
```

## See also

 * <span class="badge builder"></span> [AlertRuleInstanceTotalsBuilder](./builder-AlertRuleInstanceTotalsBuilder.md)
