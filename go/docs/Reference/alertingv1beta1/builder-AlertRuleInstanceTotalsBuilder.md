---
title: <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder
---
# <span class="badge builder"></span> AlertRuleInstanceTotalsBuilder

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Constructor

```go
func NewAlertRuleInstanceTotalsBuilder() *AlertRuleInstanceTotalsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AlertRuleInstanceTotalsBuilder) Build() (AlertRuleInstanceTotals, error)
```

### <span class="badge object-method"></span> Error

```go
func (builder *AlertRuleInstanceTotalsBuilder) Error(errorArg int64) *AlertRuleInstanceTotalsBuilder
```

### <span class="badge object-method"></span> Firing

```go
func (builder *AlertRuleInstanceTotalsBuilder) Firing(firing int64) *AlertRuleInstanceTotalsBuilder
```

### <span class="badge object-method"></span> Healthy

```go
func (builder *AlertRuleInstanceTotalsBuilder) Healthy(healthy int64) *AlertRuleInstanceTotalsBuilder
```

### <span class="badge object-method"></span> Nodata

```go
func (builder *AlertRuleInstanceTotalsBuilder) Nodata(nodata int64) *AlertRuleInstanceTotalsBuilder
```

### <span class="badge object-method"></span> Pending

```go
func (builder *AlertRuleInstanceTotalsBuilder) Pending(pending int64) *AlertRuleInstanceTotalsBuilder
```

### <span class="badge object-method"></span> Recovering

```go
func (builder *AlertRuleInstanceTotalsBuilder) Recovering(recovering int64) *AlertRuleInstanceTotalsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AlertRuleInstanceTotals](./object-AlertRuleInstanceTotals.md)
