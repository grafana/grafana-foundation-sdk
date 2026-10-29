---
title: <span class="badge builder"></span> ExpressionBuilder
---
# <span class="badge builder"></span> ExpressionBuilder

## Constructor

```go
func NewExpressionBuilder() *ExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExpressionBuilder) Build() (Expression, error)
```

### <span class="badge object-method"></span> DatasourceUID

The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource

```go
func (builder *ExpressionBuilder) DatasourceUID(datasourceUID alertingv0alpha1.DatasourceUID) *ExpressionBuilder
```

### <span class="badge object-method"></span> Model

```go
func (builder *ExpressionBuilder) Model(model any) *ExpressionBuilder
```

### <span class="badge object-method"></span> QueryType

The type of query if this is a query expression

```go
func (builder *ExpressionBuilder) QueryType(queryType string) *ExpressionBuilder
```

### <span class="badge object-method"></span> RelativeTimeRange

```go
func (builder *ExpressionBuilder) RelativeTimeRange(relativeTimeRange cog.Builder[alertingv0alpha1.RelativeTimeRange]) *ExpressionBuilder
```

### <span class="badge object-method"></span> Source

Used to mark the expression to be used as the final source for the rule evaluation

Only one expression in a rule can be marked as the source

For AlertRules, this is the expression that will be evaluated against the alerting condition

For RecordingRules, this is the expression that will be recorded

```go
func (builder *ExpressionBuilder) Source(source bool) *ExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Expression](./object-Expression.md)
