---
title: <span class="badge object-type-struct"></span> ExtendedStat
---
# <span class="badge object-type-struct"></span> ExtendedStat

## Definition

```go
type ExtendedStat struct {
    Label string `json:"label"`
    Value elasticsearch.ExtendedStatMetaType `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExtendedStat` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (extendedStat *ExtendedStat) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExtendedStat` objects.

```go
func (extendedStat *ExtendedStat) Equals(other ExtendedStat) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExtendedStat` fields for violations and returns them.

```go
func (extendedStat *ExtendedStat) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExtendedStatBuilder](./builder-ExtendedStatBuilder.md)
