---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorOperator
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorOperator

## Definition

```go
type BuilderQueryEditorOperator struct {
    Name string `json:"name"`
    Value string `json:"value"`
    LabelValue *string `json:"labelValue,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorOperator` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorOperator *BuilderQueryEditorOperator) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorOperator` objects.

```go
func (builderQueryEditorOperator *BuilderQueryEditorOperator) Equals(other BuilderQueryEditorOperator) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorOperator` fields for violations and returns them.

```go
func (builderQueryEditorOperator *BuilderQueryEditorOperator) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorOperatorBuilder](./builder-BuilderQueryEditorOperatorBuilder.md)
