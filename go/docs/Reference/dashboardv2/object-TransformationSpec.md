---
title: <span class="badge object-type-struct"></span> TransformationSpec
---
# <span class="badge object-type-struct"></span> TransformationSpec

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```go
type TransformationSpec struct {
    // Disabled transformations are skipped
    Disabled *bool `json:"disabled,omitempty"`
    // Optional frame matcher. When missing it will be applied to all results
    Filter *dashboardv2.MatcherConfig `json:"filter,omitempty"`
    // Where to pull DataFrames from as input to transformation
    Topic *dashboardv2.DataTopic `json:"topic,omitempty"`
    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    Options any `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TransformationSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (transformationSpec *TransformationSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TransformationSpec` objects.

```go
func (transformationSpec *TransformationSpec) Equals(other TransformationSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TransformationSpec` fields for violations and returns them.

```go
func (transformationSpec *TransformationSpec) Validate() error
```

