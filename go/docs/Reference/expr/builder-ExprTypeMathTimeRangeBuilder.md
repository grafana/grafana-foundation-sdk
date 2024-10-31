---
title: <span class="badge builder"></span> ExprTypeMathTimeRangeBuilder
---
# <span class="badge builder"></span> ExprTypeMathTimeRangeBuilder

## Constructor

```go
func NewExprTypeMathTimeRangeBuilder() *ExprTypeMathTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeMathTimeRangeBuilder) Build() (ExprTypeMathTimeRange, error)
```

### <span class="badge object-method"></span> From

From is the start time of the query.

```go
func (builder *ExprTypeMathTimeRangeBuilder) From(from string) *ExprTypeMathTimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *ExprTypeMathTimeRangeBuilder) To(to string) *ExprTypeMathTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeMathTimeRange](./object-ExprTypeMathTimeRange.md)
