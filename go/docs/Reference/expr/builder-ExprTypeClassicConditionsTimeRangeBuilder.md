---
title: <span class="badge builder"></span> ExprTypeClassicConditionsTimeRangeBuilder
---
# <span class="badge builder"></span> ExprTypeClassicConditionsTimeRangeBuilder

## Constructor

```go
func NewExprTypeClassicConditionsTimeRangeBuilder() *ExprTypeClassicConditionsTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) Build() (ExprTypeClassicConditionsTimeRange, error)
```

### <span class="badge object-method"></span> From

From is the start time of the query.

```go
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) From(from string) *ExprTypeClassicConditionsTimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) To(to string) *ExprTypeClassicConditionsTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeClassicConditionsTimeRange](./object-ExprTypeClassicConditionsTimeRange.md)
