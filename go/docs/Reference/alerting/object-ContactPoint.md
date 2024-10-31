---
title: <span class="badge object-type-struct"></span> ContactPoint
---
# <span class="badge object-type-struct"></span> ContactPoint

EmbeddedContactPoint is the contact point type that is used

by grafanas embedded alertmanager implementation.

## Definition

```go
type ContactPoint struct {
    DisableResolveMessage *bool `json:"disableResolveMessage,omitempty"`
    // Name is used as grouping key in the UI. Contact points with the
    // same name will be grouped in the UI.
    Name *string `json:"name,omitempty"`
    Provenance *string `json:"provenance,omitempty"`
    Settings alerting.Json `json:"settings"`
    Type alerting.ContactPointType `json:"type"`
    // UID is the unique identifier of the contact point. The UID can be
    // set by the user.
    Uid *string `json:"uid,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ContactPoint` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (contactPoint *ContactPoint) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ContactPoint` objects.

```go
func (contactPoint *ContactPoint) Equals(other ContactPoint) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ContactPoint` fields for violations and returns them.

```go
func (contactPoint *ContactPoint) Validate() error
```

## See also

 * <span class="badge builder"></span> [ContactPointBuilder](./builder-ContactPointBuilder.md)
