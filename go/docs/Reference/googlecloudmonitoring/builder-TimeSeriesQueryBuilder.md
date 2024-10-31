---
title: <span class="badge builder"></span> TimeSeriesQueryBuilder
---
# <span class="badge builder"></span> TimeSeriesQueryBuilder

## Constructor

```go
func NewTimeSeriesQueryBuilder() *TimeSeriesQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimeSeriesQueryBuilder) Build() (TimeSeriesQuery, error)
```

### <span class="badge object-method"></span> GraphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```go
func (builder *TimeSeriesQueryBuilder) GraphPeriod(graphPeriod string) *TimeSeriesQueryBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *TimeSeriesQueryBuilder) ProjectName(projectName string) *TimeSeriesQueryBuilder
```

### <span class="badge object-method"></span> Query

MQL query to be executed.

```go
func (builder *TimeSeriesQueryBuilder) Query(query string) *TimeSeriesQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeSeriesQuery](./object-TimeSeriesQuery.md)
