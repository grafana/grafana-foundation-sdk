---
title: <span class="badge object-type-struct"></span> Dashboardv2RegexMapOptions
---
# <span class="badge object-type-struct"></span> Dashboardv2RegexMapOptions

## Definition

```go
type Dashboardv2RegexMapOptions struct {
    // Regular expression to match against
    Pattern string `json:"pattern"`
    // Config to apply when the value matches the regex
    Result dashboardv2.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboardv2RegexMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardv2RegexMapOptions *Dashboardv2RegexMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboardv2RegexMapOptions` objects.

```go
func (dashboardv2RegexMapOptions *Dashboardv2RegexMapOptions) Equals(other Dashboardv2RegexMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboardv2RegexMapOptions` fields for violations and returns them.

```go
func (dashboardv2RegexMapOptions *Dashboardv2RegexMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [Dashboardv2RegexMapOptionsBuilder](./builder-Dashboardv2RegexMapOptionsBuilder.md)
