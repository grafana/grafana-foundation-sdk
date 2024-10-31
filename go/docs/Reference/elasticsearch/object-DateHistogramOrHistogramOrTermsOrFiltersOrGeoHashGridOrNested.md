---
title: <span class="badge object-type-struct"></span> DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested
---
# <span class="badge object-type-struct"></span> DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested

## Definition

```go
type DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested struct {
    DateHistogram *elasticsearch.DateHistogram `json:"DateHistogram,omitempty"`
    Histogram *elasticsearch.Histogram `json:"Histogram,omitempty"`
    Terms *elasticsearch.Terms `json:"Terms,omitempty"`
    Filters *elasticsearch.Filters `json:"Filters,omitempty"`
    GeoHashGrid *elasticsearch.GeoHashGrid `json:"GeoHashGrid,omitempty"`
    Nested *elasticsearch.Nested `json:"Nested,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` as JSON.

```go
func (dateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` from JSON.

```go
func (dateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` objects.

```go
func (dateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) Equals(other DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested` fields for violations and returns them.

```go
func (dateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) Validate() error
```

