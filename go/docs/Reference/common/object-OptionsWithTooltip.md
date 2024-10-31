---
title: <span class="badge object-type-struct"></span> OptionsWithTooltip
---
# <span class="badge object-type-struct"></span> OptionsWithTooltip

TODO docs

## Definition

```go
type OptionsWithTooltip struct {
    Tooltip common.VizTooltipOptions `json:"tooltip"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithTooltip` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (optionsWithTooltip *OptionsWithTooltip) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `OptionsWithTooltip` objects.

```go
func (optionsWithTooltip *OptionsWithTooltip) Equals(other OptionsWithTooltip) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `OptionsWithTooltip` fields for violations and returns them.

```go
func (optionsWithTooltip *OptionsWithTooltip) Validate() error
```

## See also

 * <span class="badge builder"></span> [OptionsWithTooltipBuilder](./builder-OptionsWithTooltipBuilder.md)
