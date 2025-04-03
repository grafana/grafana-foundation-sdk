---
title: <span class="badge object-type-struct"></span> BackgroundConfig
---
# <span class="badge object-type-struct"></span> BackgroundConfig

## Definition

```go
type BackgroundConfig struct {
    Color *common.ColorDimensionConfig `json:"color,omitempty"`
    Image *common.ResourceDimensionConfig `json:"image,omitempty"`
    Size *canvas.BackgroundImageSize `json:"size,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BackgroundConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (backgroundConfig *BackgroundConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BackgroundConfig` objects.

```go
func (backgroundConfig *BackgroundConfig) Equals(other BackgroundConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BackgroundConfig` fields for violations and returns them.

```go
func (backgroundConfig *BackgroundConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [BackgroundConfigBuilder](./builder-BackgroundConfigBuilder.md)
