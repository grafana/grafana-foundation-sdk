---
title: <span class="badge object-type-struct"></span> CodeOptions
---
# <span class="badge object-type-struct"></span> CodeOptions

## Definition

```go
type CodeOptions struct {
    // The language passed to monaco code editor
    Language text.CodeLanguage `json:"language"`
    ShowLineNumbers bool `json:"showLineNumbers"`
    ShowMiniMap bool `json:"showMiniMap"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CodeOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (codeOptions *CodeOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CodeOptions` objects.

```go
func (codeOptions *CodeOptions) Equals(other CodeOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CodeOptions` fields for violations and returns them.

```go
func (codeOptions *CodeOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [CodeOptionsBuilder](./builder-CodeOptionsBuilder.md)
