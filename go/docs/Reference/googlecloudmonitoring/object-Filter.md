---
title: <span class="badge object-type-struct"></span> Filter
---
# <span class="badge object-type-struct"></span> Filter

Query filter representation.

## Definition

```go
type Filter struct {
    // Filter key.
    Key string `json:"key"`
    // Filter operator.
    Operator string `json:"operator"`
    // Filter value.
    Value string `json:"value"`
    // Filter condition.
    Condition *string `json:"condition,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Filter` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (filter *Filter) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Filter` objects.

```go
func (filter *Filter) Equals(other Filter) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Filter` fields for violations and returns them.

```go
func (filter *Filter) Validate() error
```

## See also

 * <span class="badge builder"></span> [FilterBuilder](./builder-FilterBuilder.md)
