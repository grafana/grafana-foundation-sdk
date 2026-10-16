---
title: <span class="badge object-type-struct"></span> MarkdownCellContentKindOrCodeCellContentKind
---
# <span class="badge object-type-struct"></span> MarkdownCellContentKindOrCodeCellContentKind

## Definition

```go
type MarkdownCellContentKindOrCodeCellContentKind struct {
    MarkdownCellContentKind *dashboardv2beta1.MarkdownCellContentKind `json:"MarkdownCellContentKind,omitempty"`
    CodeCellContentKind *dashboardv2beta1.CodeCellContentKind `json:"CodeCellContentKind,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `MarkdownCellContentKindOrCodeCellContentKind` as JSON.

```go
func (markdownCellContentKindOrCodeCellContentKind *MarkdownCellContentKindOrCodeCellContentKind) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `MarkdownCellContentKindOrCodeCellContentKind` from JSON.

```go
func (markdownCellContentKindOrCodeCellContentKind *MarkdownCellContentKindOrCodeCellContentKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MarkdownCellContentKindOrCodeCellContentKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (markdownCellContentKindOrCodeCellContentKind *MarkdownCellContentKindOrCodeCellContentKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MarkdownCellContentKindOrCodeCellContentKind` objects.

```go
func (markdownCellContentKindOrCodeCellContentKind *MarkdownCellContentKindOrCodeCellContentKind) Equals(other MarkdownCellContentKindOrCodeCellContentKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MarkdownCellContentKindOrCodeCellContentKind` fields for violations and returns them.

```go
func (markdownCellContentKindOrCodeCellContentKind *MarkdownCellContentKindOrCodeCellContentKind) Validate() error
```

