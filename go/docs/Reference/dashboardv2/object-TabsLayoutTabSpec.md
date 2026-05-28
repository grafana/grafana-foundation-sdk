---
title: <span class="badge object-type-struct"></span> TabsLayoutTabSpec
---
# <span class="badge object-type-struct"></span> TabsLayoutTabSpec

## Definition

```go
type TabsLayoutTabSpec struct {
    Title *string `json:"title,omitempty"`
    Layout dashboardv2.GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind `json:"layout"`
    ConditionalRendering *dashboardv2.ConditionalRenderingGroupKind `json:"conditionalRendering,omitempty"`
    Repeat *dashboardv2.TabRepeatOptions `json:"repeat,omitempty"`
    Variables []dashboardv2.VariableKind `json:"variables,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TabsLayoutTabSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tabsLayoutTabSpec *TabsLayoutTabSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TabsLayoutTabSpec` objects.

```go
func (tabsLayoutTabSpec *TabsLayoutTabSpec) Equals(other TabsLayoutTabSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TabsLayoutTabSpec` fields for violations and returns them.

```go
func (tabsLayoutTabSpec *TabsLayoutTabSpec) Validate() error
```

