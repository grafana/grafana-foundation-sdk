---
title: <span class="badge object-type-struct"></span> ConnectionArgs
---
# <span class="badge object-type-struct"></span> ConnectionArgs

## Definition

```go
type ConnectionArgs struct {
    Region *string `json:"region,omitempty"`
    Catalog *string `json:"catalog,omitempty"`
    Database *string `json:"database,omitempty"`
    ResultReuseEnabled *bool `json:"resultReuseEnabled,omitempty"`
    ResultReuseMaxAgeInMinutes *float64 `json:"resultReuseMaxAgeInMinutes,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConnectionArgs` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (connectionArgs *ConnectionArgs) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConnectionArgs` objects.

```go
func (connectionArgs *ConnectionArgs) Equals(other ConnectionArgs) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConnectionArgs` fields for violations and returns them.

```go
func (connectionArgs *ConnectionArgs) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConnectionArgsBuilder](./builder-ConnectionArgsBuilder.md)
