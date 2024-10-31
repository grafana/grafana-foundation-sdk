---
title: <span class="badge object-type-struct"></span> HistogramSettings
---
# <span class="badge object-type-struct"></span> HistogramSettings

## Definition

```go
type HistogramSettings struct {
    Interval *string `json:"interval,omitempty"`
    MinDocCount *string `json:"min_doc_count,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HistogramSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (histogramSettings *HistogramSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HistogramSettings` objects.

```go
func (histogramSettings *HistogramSettings) Equals(other HistogramSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HistogramSettings` fields for violations and returns them.

```go
func (histogramSettings *HistogramSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [HistogramSettingsBuilder](./builder-HistogramSettingsBuilder.md)
