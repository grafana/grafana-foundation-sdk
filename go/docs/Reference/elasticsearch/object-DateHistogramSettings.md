---
title: <span class="badge object-type-struct"></span> DateHistogramSettings
---
# <span class="badge object-type-struct"></span> DateHistogramSettings

## Definition

```go
type DateHistogramSettings struct {
    Interval *string `json:"interval,omitempty"`
    MinDocCount *string `json:"min_doc_count,omitempty"`
    TrimEdges *string `json:"trimEdges,omitempty"`
    Offset *string `json:"offset,omitempty"`
    TimeZone *string `json:"timeZone,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DateHistogramSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dateHistogramSettings *DateHistogramSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DateHistogramSettings` objects.

```go
func (dateHistogramSettings *DateHistogramSettings) Equals(other DateHistogramSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DateHistogramSettings` fields for violations and returns them.

```go
func (dateHistogramSettings *DateHistogramSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [DateHistogramSettingsBuilder](./builder-DateHistogramSettingsBuilder.md)
