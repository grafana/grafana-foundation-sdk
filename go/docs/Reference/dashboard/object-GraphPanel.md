---
title: <span class="badge object-type-struct"></span> GraphPanel
---
# <span class="badge object-type-struct"></span> GraphPanel

Support for legacy graph panel.

@deprecated this a deprecated panel type

## Definition

```go
type GraphPanel struct {
    Type string `json:"type"`
    // @deprecated this is part of deprecated graph panel
    Legend *dashboard.DashboardGraphPanelLegend `json:"legend,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GraphPanel` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (graphPanel *GraphPanel) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GraphPanel` objects.

```go
func (graphPanel *GraphPanel) Equals(other GraphPanel) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GraphPanel` fields for violations and returns them.

```go
func (graphPanel *GraphPanel) Validate() error
```

## See also

 * <span class="badge builder"></span> [GraphPanelBuilder](./builder-GraphPanelBuilder.md)
