---
title: <span class="badge object-type-struct"></span> TabsLayoutKind
---
# <span class="badge object-type-struct"></span> TabsLayoutKind

## Definition

```go
type TabsLayoutKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.TabsLayoutSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tabsLayoutKind *TabsLayoutKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TabsLayoutKind` objects.

```go
func (tabsLayoutKind *TabsLayoutKind) Equals(other TabsLayoutKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TabsLayoutKind` fields for violations and returns them.

```go
func (tabsLayoutKind *TabsLayoutKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [TabsLayoutBuilder](./builder-TabsLayoutBuilder.md)
