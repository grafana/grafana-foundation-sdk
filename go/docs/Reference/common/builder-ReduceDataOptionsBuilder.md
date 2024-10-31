---
title: <span class="badge builder"></span> ReduceDataOptionsBuilder
---
# <span class="badge builder"></span> ReduceDataOptionsBuilder

## Constructor

```go
func NewReduceDataOptionsBuilder() *ReduceDataOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ReduceDataOptionsBuilder) Build() (ReduceDataOptions, error)
```

### <span class="badge object-method"></span> Calcs

When !values, pick one value for the whole field

```go
func (builder *ReduceDataOptionsBuilder) Calcs(calcs []string) *ReduceDataOptionsBuilder
```

### <span class="badge object-method"></span> Fields

Which fields to show.  By default this is only numeric fields

```go
func (builder *ReduceDataOptionsBuilder) Fields(fields string) *ReduceDataOptionsBuilder
```

### <span class="badge object-method"></span> Limit

if showing all values limit

```go
func (builder *ReduceDataOptionsBuilder) Limit(limit float64) *ReduceDataOptionsBuilder
```

### <span class="badge object-method"></span> Values

If true show each row value

```go
func (builder *ReduceDataOptionsBuilder) Values(values bool) *ReduceDataOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ReduceDataOptions](./object-ReduceDataOptions.md)
