---
title: <span class="badge object-type-struct"></span> OptionsWithTimezones
---
# <span class="badge object-type-struct"></span> OptionsWithTimezones

TODO docs

## Definition

```go
type OptionsWithTimezones struct {
    Timezone []common.TimeZone `json:"timezone,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `OptionsWithTimezones` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (optionsWithTimezones *OptionsWithTimezones) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `OptionsWithTimezones` objects.

```go
func (optionsWithTimezones *OptionsWithTimezones) Equals(other OptionsWithTimezones) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `OptionsWithTimezones` fields for violations and returns them.

```go
func (optionsWithTimezones *OptionsWithTimezones) Validate() error
```

## See also

 * <span class="badge builder"></span> [OptionsWithTimezonesBuilder](./builder-OptionsWithTimezonesBuilder.md)
