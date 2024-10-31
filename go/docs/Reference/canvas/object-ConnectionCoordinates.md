---
title: <span class="badge object-type-struct"></span> ConnectionCoordinates
---
# <span class="badge object-type-struct"></span> ConnectionCoordinates

## Definition

```go
type ConnectionCoordinates struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConnectionCoordinates` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (connectionCoordinates *ConnectionCoordinates) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConnectionCoordinates` objects.

```go
func (connectionCoordinates *ConnectionCoordinates) Equals(other ConnectionCoordinates) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConnectionCoordinates` fields for violations and returns them.

```go
func (connectionCoordinates *ConnectionCoordinates) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConnectionCoordinatesBuilder](./builder-ConnectionCoordinatesBuilder.md)
