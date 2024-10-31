---
title: <span class="badge builder"></span> PromQLQueryBuilder
---
# <span class="badge builder"></span> PromQLQueryBuilder

## Constructor

```go
func NewPromQLQueryBuilder() *PromQLQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PromQLQueryBuilder) Build() (PromQLQuery, error)
```

### <span class="badge object-method"></span> Expr

PromQL expression/query to be executed.

```go
func (builder *PromQLQueryBuilder) Expr(expr string) *PromQLQueryBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *PromQLQueryBuilder) ProjectName(projectName string) *PromQLQueryBuilder
```

### <span class="badge object-method"></span> Step

PromQL min step

```go
func (builder *PromQLQueryBuilder) Step(step string) *PromQLQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PromQLQuery](./object-PromQLQuery.md)
