---
title: <span class="badge object-type-struct"></span> DataTransformerConfig
---
# <span class="badge object-type-struct"></span> DataTransformerConfig

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```go
type DataTransformerConfig struct {
    // Unique identifier of transformer
    Id string `json:"id"`
    // Disabled transformations are skipped
    Disabled *bool `json:"disabled,omitempty"`
    // Optional frame matcher. When missing it will be applied to all results
    Filter *dashboard.MatcherConfig `json:"filter,omitempty"`
    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    Options any `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DataTransformerConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dataTransformerConfig *DataTransformerConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DataTransformerConfig` objects.

```go
func (dataTransformerConfig *DataTransformerConfig) Equals(other DataTransformerConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DataTransformerConfig` fields for violations and returns them.

```go
func (dataTransformerConfig *DataTransformerConfig) Validate() error
```

