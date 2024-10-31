---
title: <span class="badge builder"></span> SLOQueryBuilder
---
# <span class="badge builder"></span> SLOQueryBuilder

## Constructor

```go
func NewSLOQueryBuilder() *SLOQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SLOQueryBuilder) Build() (SLOQuery, error)
```

### <span class="badge object-method"></span> AlignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```go
func (builder *SLOQueryBuilder) AlignmentPeriod(alignmentPeriod string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> Goal

SLO goal value.

```go
func (builder *SLOQueryBuilder) Goal(goal float64) *SLOQueryBuilder
```

### <span class="badge object-method"></span> LookbackPeriod

Specific lookback period for the SLO.

```go
func (builder *SLOQueryBuilder) LookbackPeriod(lookbackPeriod string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> PerSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```go
func (builder *SLOQueryBuilder) PerSeriesAligner(perSeriesAligner string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *SLOQueryBuilder) ProjectName(projectName string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> SelectorName

SLO selector.

```go
func (builder *SLOQueryBuilder) SelectorName(selectorName string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> ServiceId

ID for the service the SLO is in.

```go
func (builder *SLOQueryBuilder) ServiceId(serviceId string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> ServiceName

Name for the service the SLO is in.

```go
func (builder *SLOQueryBuilder) ServiceName(serviceName string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> SloId

ID for the SLO.

```go
func (builder *SLOQueryBuilder) SloId(sloId string) *SLOQueryBuilder
```

### <span class="badge object-method"></span> SloName

Name of the SLO.

```go
func (builder *SLOQueryBuilder) SloName(sloName string) *SLOQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SLOQuery](./object-SLOQuery.md)
