---
title: <span class="badge builder"></span> TracesFilterBuilder
---
# <span class="badge builder"></span> TracesFilterBuilder

## Constructor

```go
func NewTracesFilterBuilder() *TracesFilterBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TracesFilterBuilder) Build() (TracesFilter, error)
```

### <span class="badge object-method"></span> Filters

Values to filter by.

```go
func (builder *TracesFilterBuilder) Filters(filters []string) *TracesFilterBuilder
```

### <span class="badge object-method"></span> Operation

Comparison operator to use. Either equals or not equals.

```go
func (builder *TracesFilterBuilder) Operation(operation string) *TracesFilterBuilder
```

### <span class="badge object-method"></span> Property

Property name, auto-populated based on available traces.

```go
func (builder *TracesFilterBuilder) Property(property string) *TracesFilterBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TracesFilter](./object-TracesFilter.md)
