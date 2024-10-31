---
title: <span class="badge object-type-struct"></span> Team
---
# <span class="badge object-type-struct"></span> Team

## Definition

```go
type Team struct {
    // Name of the team.
    Name string `json:"name"`
    // Email of the team.
    Email *string `json:"email,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Team` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (team *Team) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Team` objects.

```go
func (team *Team) Equals(other Team) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Team` fields for violations and returns them.

```go
func (team *Team) Validate() error
```

## See also

 * <span class="badge builder"></span> [TeamBuilder](./builder-TeamBuilder.md)
