---
title: <span class="badge object-type-struct"></span> OptionsWithLegend
---
# <span class="badge object-type-struct"></span> OptionsWithLegend

TODO docs

## Definition

```go
type OptionsWithLegend struct {
    Legend common.VizLegendOptions `json:"legend"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithLegend` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (optionsWithLegend *OptionsWithLegend) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `OptionsWithLegend` objects.

```go
func (optionsWithLegend *OptionsWithLegend) Equals(other OptionsWithLegend) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `OptionsWithLegend` fields for violations and returns them.

```go
func (optionsWithLegend *OptionsWithLegend) Validate() error
```

## See also

 * <span class="badge builder"></span> [OptionsWithLegendBuilder](./builder-OptionsWithLegendBuilder.md)
