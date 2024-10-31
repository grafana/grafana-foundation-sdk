---
title: <span class="badge builder"></span> ExprTypeThresholdTimeRangeBuilder
---
# <span class="badge builder"></span> ExprTypeThresholdTimeRangeBuilder

## Constructor

```go
func NewExprTypeThresholdTimeRangeBuilder() *ExprTypeThresholdTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeThresholdTimeRangeBuilder) Build() (ExprTypeThresholdTimeRange, error)
```

### <span class="badge object-method"></span> From

From is the start time of the query.

```go
func (builder *ExprTypeThresholdTimeRangeBuilder) From(from string) *ExprTypeThresholdTimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *ExprTypeThresholdTimeRangeBuilder) To(to string) *ExprTypeThresholdTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeThresholdTimeRange](./object-ExprTypeThresholdTimeRange.md)
