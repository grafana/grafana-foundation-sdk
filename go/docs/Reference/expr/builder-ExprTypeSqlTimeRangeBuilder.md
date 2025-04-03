---
title: <span class="badge builder"></span> ExprTypeSqlTimeRangeBuilder
---
# <span class="badge builder"></span> ExprTypeSqlTimeRangeBuilder

## Constructor

```go
func NewExprTypeSqlTimeRangeBuilder() *ExprTypeSqlTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeSqlTimeRangeBuilder) Build() (ExprTypeSqlTimeRange, error)
```

### <span class="badge object-method"></span> From

From is the start time of the query.

```go
func (builder *ExprTypeSqlTimeRangeBuilder) From(from string) *ExprTypeSqlTimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *ExprTypeSqlTimeRangeBuilder) To(to string) *ExprTypeSqlTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeSqlTimeRange](./object-ExprTypeSqlTimeRange.md)
