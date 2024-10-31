---
title: <span class="badge builder"></span> ExprTypeReduceTimeRangeBuilder
---
# <span class="badge builder"></span> ExprTypeReduceTimeRangeBuilder

## Constructor

```go
func NewExprTypeReduceTimeRangeBuilder() *ExprTypeReduceTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeReduceTimeRangeBuilder) Build() (ExprTypeReduceTimeRange, error)
```

### <span class="badge object-method"></span> From

From is the start time of the query.

```go
func (builder *ExprTypeReduceTimeRangeBuilder) From(from string) *ExprTypeReduceTimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *ExprTypeReduceTimeRangeBuilder) To(to string) *ExprTypeReduceTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeReduceTimeRange](./object-ExprTypeReduceTimeRange.md)
