---
title: <span class="badge object-type-struct"></span> MarkdownCellContentSpec
---
# <span class="badge object-type-struct"></span> MarkdownCellContentSpec

## Definition

```go
type MarkdownCellContentSpec struct {
    Text string `json:"text"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MarkdownCellContentSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (markdownCellContentSpec *MarkdownCellContentSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MarkdownCellContentSpec` objects.

```go
func (markdownCellContentSpec *MarkdownCellContentSpec) Equals(other MarkdownCellContentSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MarkdownCellContentSpec` fields for violations and returns them.

```go
func (markdownCellContentSpec *MarkdownCellContentSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [MarkdownCellContentSpecBuilder](./builder-MarkdownCellContentSpecBuilder.md)
