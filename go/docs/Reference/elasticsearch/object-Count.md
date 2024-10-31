---
title: <span class="badge object-type-struct"></span> Count
---
# <span class="badge object-type-struct"></span> Count

## Definition

```go
type Count struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Count` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (count *Count) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Count` objects.

```go
func (count *Count) Equals(other Count) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Count` fields for violations and returns them.

```go
func (count *Count) Validate() error
```

## See also

 * <span class="badge builder"></span> [CountBuilder](./builder-CountBuilder.md)
