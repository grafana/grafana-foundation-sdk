---
title: <span class="badge builder"></span> ExprBuilder
---
# <span class="badge builder"></span> ExprBuilder

## Constructor

```go
func NewExprBuilder() *ExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> TypeClassicConditions

```go
func (builder *ExprBuilder) TypeClassicConditions(typeClassicConditions cog.Builder[expr.TypeClassicConditions]) *ExprBuilder
```

### <span class="badge object-method"></span> TypeMath

```go
func (builder *ExprBuilder) TypeMath(typeMath cog.Builder[expr.TypeMath]) *ExprBuilder
```

### <span class="badge object-method"></span> TypeReduce

```go
func (builder *ExprBuilder) TypeReduce(typeReduce cog.Builder[expr.TypeReduce]) *ExprBuilder
```

### <span class="badge object-method"></span> TypeResample

```go
func (builder *ExprBuilder) TypeResample(typeResample cog.Builder[expr.TypeResample]) *ExprBuilder
```

### <span class="badge object-method"></span> TypeSql

```go
func (builder *ExprBuilder) TypeSql(typeSql cog.Builder[expr.TypeSql]) *ExprBuilder
```

### <span class="badge object-method"></span> TypeThreshold

```go
func (builder *ExprBuilder) TypeThreshold(typeThreshold cog.Builder[expr.TypeThreshold]) *ExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
