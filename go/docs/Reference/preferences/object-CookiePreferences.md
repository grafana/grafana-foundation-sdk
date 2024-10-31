---
title: <span class="badge object-type-struct"></span> CookiePreferences
---
# <span class="badge object-type-struct"></span> CookiePreferences

## Definition

```go
type CookiePreferences struct {
    Analytics any `json:"analytics,omitempty"`
    Performance any `json:"performance,omitempty"`
    Functional any `json:"functional,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CookiePreferences` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cookiePreferences *CookiePreferences) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CookiePreferences` objects.

```go
func (cookiePreferences *CookiePreferences) Equals(other CookiePreferences) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CookiePreferences` fields for violations and returns them.

```go
func (cookiePreferences *CookiePreferences) Validate() error
```

## See also

 * <span class="badge builder"></span> [CookiePreferencesBuilder](./builder-CookiePreferencesBuilder.md)
