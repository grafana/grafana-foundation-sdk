---
title: <span class="badge object-type-struct"></span> CodeCellContentSpec
---
# <span class="badge object-type-struct"></span> CodeCellContentSpec

## Definition

```go
type CodeCellContentSpec struct {
    Language string `json:"language"`
    Code string `json:"code"`
    Highlight []int64 `json:"highlight,omitempty"`
    Annotation *string `json:"annotation,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CodeCellContentSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (codeCellContentSpec *CodeCellContentSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CodeCellContentSpec` objects.

```go
func (codeCellContentSpec *CodeCellContentSpec) Equals(other CodeCellContentSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CodeCellContentSpec` fields for violations and returns them.

```go
func (codeCellContentSpec *CodeCellContentSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [CodeCellContentSpecBuilder](./builder-CodeCellContentSpecBuilder.md)
