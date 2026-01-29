---
title: <span class="badge object-type-struct"></span> TabsLayoutTabKind
---
# <span class="badge object-type-struct"></span> TabsLayoutTabKind

## Definition

```go
type TabsLayoutTabKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.TabsLayoutTabSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutTabKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tabsLayoutTabKind *TabsLayoutTabKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TabsLayoutTabKind` objects.

```go
func (tabsLayoutTabKind *TabsLayoutTabKind) Equals(other TabsLayoutTabKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TabsLayoutTabKind` fields for violations and returns them.

```go
func (tabsLayoutTabKind *TabsLayoutTabKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [TabsLayoutTabBuilder](./builder-TabsLayoutTabBuilder.md)
