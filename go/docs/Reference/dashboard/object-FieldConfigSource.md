---
title: <span class="badge object-type-struct"></span> FieldConfigSource
---
# <span class="badge object-type-struct"></span> FieldConfigSource

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```go
type FieldConfigSource struct {
    // Defaults are the options applied to all fields.
    Defaults dashboard.FieldConfig `json:"defaults"`
    // Overrides are the options applied to specific fields overriding the defaults.
    Overrides []dashboard.DashboardFieldConfigSourceOverrides `json:"overrides"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfigSource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fieldConfigSource *FieldConfigSource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FieldConfigSource` objects.

```go
func (fieldConfigSource *FieldConfigSource) Equals(other FieldConfigSource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FieldConfigSource` fields for violations and returns them.

```go
func (fieldConfigSource *FieldConfigSource) Validate() error
```

