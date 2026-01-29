---
title: <span class="badge object-type-struct"></span> Dashboardv2beta1ActionStyle
---
# <span class="badge object-type-struct"></span> Dashboardv2beta1ActionStyle

## Definition

```go
type Dashboardv2beta1ActionStyle struct {
    BackgroundColor *string `json:"backgroundColor,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2beta1ActionStyle` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2beta1ActionStyle *Dashboardv2beta1ActionStyle) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2beta1ActionStyle` objects.

```go
func (dashboardv2beta1ActionStyle *Dashboardv2beta1ActionStyle) Equals(other Dashboardv2beta1ActionStyle) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2beta1ActionStyle` fields for violations and returns them.

```go
func (dashboardv2beta1ActionStyle *Dashboardv2beta1ActionStyle) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2beta1ActionStyleBuilder](./builder-Dashboardv2beta1ActionStyleBuilder.md)
