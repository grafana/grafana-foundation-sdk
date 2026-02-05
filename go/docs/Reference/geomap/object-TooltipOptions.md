---
title: <span class="badge object-type-struct"></span> TooltipOptions
---
# <span class="badge object-type-struct"></span> TooltipOptions

## Definition

```go
type TooltipOptions struct {
    Mode geomap.TooltipMode `json:"mode"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TooltipOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tooltipOptions *TooltipOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TooltipOptions` objects.

```go
func (tooltipOptions *TooltipOptions) Equals(other TooltipOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TooltipOptions` fields for violations and returns them.

```go
func (tooltipOptions *TooltipOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TooltipOptionsBuilder](./builder-TooltipOptionsBuilder.md)
