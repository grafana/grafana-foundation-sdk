---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    // Represents the index of the selected frame
    FrameIndex float64 `json:"frameIndex"`
    // Controls whether the panel should show the header
    ShowHeader bool `json:"showHeader"`
    // Controls whether the header should show icons for the column types
    ShowTypeIcons *bool `json:"showTypeIcons,omitempty"`
    // Used to control row sorting
    SortBy []common.TableSortByFieldState `json:"sortBy,omitempty"`
    // Controls footer options
    Footer *common.TableFooterOptions `json:"footer,omitempty"`
    // Controls the height of the rows
    CellHeight *common.TableCellHeight `json:"cellHeight,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (options *Options) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Options` objects.

```go
func (options *Options) Equals(other Options) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.

```go
func (options *Options) Validate() error
```

