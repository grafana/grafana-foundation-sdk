---
title: <span class="badge object-type-struct"></span> SingleStatBaseOptions
---
# <span class="badge object-type-struct"></span> SingleStatBaseOptions

TODO docs

## Definition

```go
type SingleStatBaseOptions struct {
    ReduceOptions common.ReduceDataOptions `json:"reduceOptions"`
    Text *common.VizTextDisplayOptions `json:"text,omitempty"`
    Orientation common.VizOrientation `json:"orientation"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SingleStatBaseOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (singleStatBaseOptions *SingleStatBaseOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SingleStatBaseOptions` objects.

```go
func (singleStatBaseOptions *SingleStatBaseOptions) Equals(other SingleStatBaseOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SingleStatBaseOptions` fields for violations and returns them.

```go
func (singleStatBaseOptions *SingleStatBaseOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [SingleStatBaseOptionsBuilder](./builder-SingleStatBaseOptionsBuilder.md)
