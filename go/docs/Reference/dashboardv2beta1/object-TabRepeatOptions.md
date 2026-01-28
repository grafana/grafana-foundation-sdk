---
title: <span class="badge object-type-struct"></span> TabRepeatOptions
---
# <span class="badge object-type-struct"></span> TabRepeatOptions

## Definition

```go
type TabRepeatOptions struct {
    Mode dashboardv2beta1.RepeatMode `json:"mode"`
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabRepeatOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tabRepeatOptions *TabRepeatOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TabRepeatOptions` objects.

```go
func (tabRepeatOptions *TabRepeatOptions) Equals(other TabRepeatOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TabRepeatOptions` fields for violations and returns them.

```go
func (tabRepeatOptions *TabRepeatOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TabRepeatOptionsBuilder](./builder-TabRepeatOptionsBuilder.md)
