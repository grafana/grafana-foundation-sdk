---
title: <span class="badge builder"></span> TraceqlFilterBuilder
---
# <span class="badge builder"></span> TraceqlFilterBuilder

## Constructor

```go
func NewTraceqlFilterBuilder() *TraceqlFilterBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TraceqlFilterBuilder) Build() (TraceqlFilter, error)
```

### <span class="badge object-method"></span> Id

Uniquely identify the filter, will not be used in the query generation

```go
func (builder *TraceqlFilterBuilder) Id(id string) *TraceqlFilterBuilder
```

### <span class="badge object-method"></span> Operator

The operator that connects the tag to the value, for example: =, >, !=, =~

```go
func (builder *TraceqlFilterBuilder) Operator(operator string) *TraceqlFilterBuilder
```

### <span class="badge object-method"></span> Scope

The scope of the filter, can either be unscoped/all scopes, resource or span

```go
func (builder *TraceqlFilterBuilder) Scope(scope tempo.TraceqlSearchScope) *TraceqlFilterBuilder
```

### <span class="badge object-method"></span> Tag

The tag for the search filter, for example: .http.status_code, .service.name, status

```go
func (builder *TraceqlFilterBuilder) Tag(tag string) *TraceqlFilterBuilder
```

### <span class="badge object-method"></span> Value

The value for the search filter

```go
func (builder *TraceqlFilterBuilder) Value(values []string) *TraceqlFilterBuilder
```

### <span class="badge object-method"></span> ValueType

The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query

```go
func (builder *TraceqlFilterBuilder) ValueType(valueType string) *TraceqlFilterBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TraceqlFilter](./object-TraceqlFilter.md)
