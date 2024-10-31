---
title: <span class="badge object-type-struct"></span> NavbarPreference
---
# <span class="badge object-type-struct"></span> NavbarPreference

## Definition

```go
type NavbarPreference struct {
    BookmarkUrls []string `json:"bookmarkUrls"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NavbarPreference` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (navbarPreference *NavbarPreference) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NavbarPreference` objects.

```go
func (navbarPreference *NavbarPreference) Equals(other NavbarPreference) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NavbarPreference` fields for violations and returns them.

```go
func (navbarPreference *NavbarPreference) Validate() error
```

## See also

 * <span class="badge builder"></span> [NavbarPreferenceBuilder](./builder-NavbarPreferenceBuilder.md)
