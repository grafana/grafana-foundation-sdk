---
title: <span class="badge object-type-struct"></span> PanelOrRowPanel
---
# <span class="badge object-type-struct"></span> PanelOrRowPanel

## Definition

```go
type PanelOrRowPanel struct {
    Panel *dashboard.Panel `json:"Panel,omitempty"`
    RowPanel *dashboard.RowPanel `json:"RowPanel,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `PanelOrRowPanel` as JSON.

```go
func (panelOrRowPanel *PanelOrRowPanel) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `PanelOrRowPanel` from JSON.

```go
func (panelOrRowPanel *PanelOrRowPanel) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelOrRowPanel` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelOrRowPanel *PanelOrRowPanel) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelOrRowPanel` objects.

```go
func (panelOrRowPanel *PanelOrRowPanel) Equals(other PanelOrRowPanel) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelOrRowPanel` fields for violations and returns them.

```go
func (panelOrRowPanel *PanelOrRowPanel) Validate() error
```

