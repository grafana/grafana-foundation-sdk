---
title: <span class="badge object-type-struct"></span> TabsLayoutSpec
---
# <span class="badge object-type-struct"></span> TabsLayoutSpec

## Definition

```go
type TabsLayoutSpec struct {
    Tabs []dashboardv2beta1.TabsLayoutTabKind `json:"tabs"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tabsLayoutSpec *TabsLayoutSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TabsLayoutSpec` objects.

```go
func (tabsLayoutSpec *TabsLayoutSpec) Equals(other TabsLayoutSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TabsLayoutSpec` fields for violations and returns them.

```go
func (tabsLayoutSpec *TabsLayoutSpec) Validate() error
```

