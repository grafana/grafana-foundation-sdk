---
title: <span class="badge builder"></span> BuilderQueryExpressionBuilder
---
# <span class="badge builder"></span> BuilderQueryExpressionBuilder

## Constructor

```go
func NewBuilderQueryExpressionBuilder() *BuilderQueryExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryExpressionBuilder) Build() (BuilderQueryExpression, error)
```

### <span class="badge object-method"></span> Columns

```go
func (builder *BuilderQueryExpressionBuilder) Columns(columns cog.Builder[azuremonitor.BuilderQueryEditorColumnsExpression]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> From

```go
func (builder *BuilderQueryExpressionBuilder) From(from cog.Builder[azuremonitor.BuilderQueryEditorPropertyExpression]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> FuzzySearch

```go
func (builder *BuilderQueryExpressionBuilder) FuzzySearch(fuzzySearch cog.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> GroupBy

```go
func (builder *BuilderQueryExpressionBuilder) GroupBy(groupBy cog.Builder[azuremonitor.BuilderQueryEditorGroupByExpressionArray]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> Limit

```go
func (builder *BuilderQueryExpressionBuilder) Limit(limit int64) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> OrderBy

```go
func (builder *BuilderQueryExpressionBuilder) OrderBy(orderBy cog.Builder[azuremonitor.BuilderQueryEditorOrderByExpressionArray]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> Reduce

```go
func (builder *BuilderQueryExpressionBuilder) Reduce(reduce cog.Builder[azuremonitor.BuilderQueryEditorReduceExpressionArray]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> TimeFilter

```go
func (builder *BuilderQueryExpressionBuilder) TimeFilter(timeFilter cog.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) *BuilderQueryExpressionBuilder
```

### <span class="badge object-method"></span> Where

```go
func (builder *BuilderQueryExpressionBuilder) Where(where cog.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) *BuilderQueryExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryExpression](./object-BuilderQueryExpression.md)
