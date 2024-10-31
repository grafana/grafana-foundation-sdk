---
title: <span class="badge builder"></span> ExprTypeReduceSettingsBuilder
---
# <span class="badge builder"></span> ExprTypeReduceSettingsBuilder

## Constructor

```go
func NewExprTypeReduceSettingsBuilder() *ExprTypeReduceSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeReduceSettingsBuilder) Build() (ExprTypeReduceSettings, error)
```

### <span class="badge object-method"></span> Mode

Non-number reduce behavior

Possible enum values:

 - `"dropNN"` Drop non-numbers

 - `"replaceNN"` Replace non-numbers

```go
func (builder *ExprTypeReduceSettingsBuilder) Mode(mode expr.TypeReduceMode) *ExprTypeReduceSettingsBuilder
```

### <span class="badge object-method"></span> ReplaceWithValue

Only valid when mode is replace

```go
func (builder *ExprTypeReduceSettingsBuilder) ReplaceWithValue(replaceWithValue float64) *ExprTypeReduceSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeReduceSettings](./object-ExprTypeReduceSettings.md)
