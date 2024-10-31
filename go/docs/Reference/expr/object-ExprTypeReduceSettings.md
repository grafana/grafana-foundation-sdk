---
title: <span class="badge object-type-struct"></span> ExprTypeReduceSettings
---
# <span class="badge object-type-struct"></span> ExprTypeReduceSettings

## Definition

```go
type ExprTypeReduceSettings struct {
    // Non-number reduce behavior
    // Possible enum values:
    //  - `"dropNN"` Drop non-numbers
    //  - `"replaceNN"` Replace non-numbers
    Mode expr.TypeReduceMode `json:"mode"`
    // Only valid when mode is replace
    ReplaceWithValue *float64 `json:"replaceWithValue,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeReduceSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeReduceSettings *ExprTypeReduceSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeReduceSettings` objects.

```go
func (exprTypeReduceSettings *ExprTypeReduceSettings) Equals(other ExprTypeReduceSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeReduceSettings` fields for violations and returns them.

```go
func (exprTypeReduceSettings *ExprTypeReduceSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeReduceSettingsBuilder](./builder-ExprTypeReduceSettingsBuilder.md)
