---
title: <span class="badge object-type-struct"></span> RegexMap
---
# <span class="badge object-type-struct"></span> RegexMap

Maps regular expressions to replacement text and a color.

For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.

## Definition

```go
type RegexMap struct {
    Type string `json:"type"`
    // Regular expression to match against and the result to apply when the value matches the regex
    Options dashboard.DashboardRegexMapOptions `json:"options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RegexMap` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (regexMap *RegexMap) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RegexMap` objects.

```go
func (regexMap *RegexMap) Equals(other RegexMap) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RegexMap` fields for violations and returns them.

```go
func (regexMap *RegexMap) Validate() error
```

## See also

 * <span class="badge builder"></span> [RegexMapBuilder](./builder-RegexMapBuilder.md)
