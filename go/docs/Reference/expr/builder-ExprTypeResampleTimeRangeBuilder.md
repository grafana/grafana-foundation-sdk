---
title: <span class="badge builder"></span> ExprTypeResampleTimeRangeBuilder
---
# <span class="badge builder"></span> ExprTypeResampleTimeRangeBuilder

## Constructor

```go
func NewExprTypeResampleTimeRangeBuilder() *ExprTypeResampleTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeResampleTimeRangeBuilder) Build() (ExprTypeResampleTimeRange, error)
```

### <span class="badge object-method"></span> From

From is the start time of the query.

```go
func (builder *ExprTypeResampleTimeRangeBuilder) From(from string) *ExprTypeResampleTimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *ExprTypeResampleTimeRangeBuilder) To(to string) *ExprTypeResampleTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeResampleTimeRange](./object-ExprTypeResampleTimeRange.md)
