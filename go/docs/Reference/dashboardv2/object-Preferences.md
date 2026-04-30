---
title: <span class="badge object-type-struct"></span> Preferences
---
# <span class="badge object-type-struct"></span> Preferences

Dashboard specific preferences (applied per dashboard = all users using the dashboard)

## Definition

```go
type Preferences struct {
    // default layout template to be used when new containers are created
    Layout *dashboardv2.AutoGridLayoutKindOrGridLayoutKind `json:"layout,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Preferences` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (preferences *Preferences) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Preferences` objects.

```go
func (preferences *Preferences) Equals(other Preferences) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Preferences` fields for violations and returns them.

```go
func (preferences *Preferences) Validate() error
```

## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
