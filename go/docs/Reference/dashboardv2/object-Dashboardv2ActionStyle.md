---
title: <span class="badge object-type-struct"></span> Dashboardv2ActionStyle
---
# <span class="badge object-type-struct"></span> Dashboardv2ActionStyle

## Definition

```go
type Dashboardv2ActionStyle struct {
    BackgroundColor *string `json:"backgroundColor,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2ActionStyle` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2ActionStyle *Dashboardv2ActionStyle) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2ActionStyle` objects.

```go
func (dashboardv2ActionStyle *Dashboardv2ActionStyle) Equals(other Dashboardv2ActionStyle) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2ActionStyle` fields for violations and returns them.

```go
func (dashboardv2ActionStyle *Dashboardv2ActionStyle) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2ActionStyleBuilder](./builder-Dashboardv2ActionStyleBuilder.md)
