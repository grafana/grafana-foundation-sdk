---
title: <span class="badge object-type-struct"></span> GeoHashGridSettings
---
# <span class="badge object-type-struct"></span> GeoHashGridSettings

## Definition

```go
type GeoHashGridSettings struct {
    Precision *string `json:"precision,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GeoHashGridSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (geoHashGridSettings *GeoHashGridSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GeoHashGridSettings` objects.

```go
func (geoHashGridSettings *GeoHashGridSettings) Equals(other GeoHashGridSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GeoHashGridSettings` fields for violations and returns them.

```go
func (geoHashGridSettings *GeoHashGridSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [GeoHashGridSettingsBuilder](./builder-GeoHashGridSettingsBuilder.md)
