---
title: <span class="badge object-type-struct"></span> RawData
---
# <span class="badge object-type-struct"></span> RawData

## Definition

```go
type RawData struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchRawDataSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RawData` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rawData *RawData) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RawData` objects.

```go
func (rawData *RawData) Equals(other RawData) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RawData` fields for violations and returns them.

```go
func (rawData *RawData) Validate() error
```

## See also

 * <span class="badge builder"></span> [RawDataBuilder](./builder-RawDataBuilder.md)
