---
title: <span class="badge object-type-struct"></span> OptionsWithTextFormatting
---
# <span class="badge object-type-struct"></span> OptionsWithTextFormatting

TODO docs

## Definition

```go
type OptionsWithTextFormatting struct {
    Text *common.VizTextDisplayOptions `json:"text,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithTextFormatting` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (optionsWithTextFormatting *OptionsWithTextFormatting) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `OptionsWithTextFormatting` objects.

```go
func (optionsWithTextFormatting *OptionsWithTextFormatting) Equals(other OptionsWithTextFormatting) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `OptionsWithTextFormatting` fields for violations and returns them.

```go
func (optionsWithTextFormatting *OptionsWithTextFormatting) Validate() error
```

## See also

 * <span class="badge builder"></span> [OptionsWithTextFormattingBuilder](./builder-OptionsWithTextFormattingBuilder.md)
