---
title: <span class="badge object-type-struct"></span> XychartXYSeriesConfigName
---
# <span class="badge object-type-struct"></span> XychartXYSeriesConfigName

## Definition

```go
type XychartXYSeriesConfigName struct {
    Fixed *string `json:"fixed,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigName` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xychartXYSeriesConfigName *XychartXYSeriesConfigName) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XychartXYSeriesConfigName` objects.

```go
func (xychartXYSeriesConfigName *XychartXYSeriesConfigName) Equals(other XychartXYSeriesConfigName) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigName` fields for violations and returns them.

```go
func (xychartXYSeriesConfigName *XychartXYSeriesConfigName) Validate() error
```

## See also

 * <span class="badge builder"></span> [XychartXYSeriesConfigNameBuilder](./builder-XychartXYSeriesConfigNameBuilder.md)
