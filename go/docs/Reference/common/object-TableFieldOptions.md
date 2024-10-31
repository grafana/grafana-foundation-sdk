---
title: <span class="badge object-type-struct"></span> TableFieldOptions
---
# <span class="badge object-type-struct"></span> TableFieldOptions

Field options for each field within a table (e.g 10, "The String", 64.20, etc.)

Generally defines alignment, filtering capabilties, display options, etc.

## Definition

```go
type TableFieldOptions struct {
    Width *float64 `json:"width,omitempty"`
    MinWidth *float64 `json:"minWidth,omitempty"`
    Align common.FieldTextAlignment `json:"align"`
    // This field is deprecated in favor of using cellOptions
    DisplayMode *common.TableCellDisplayMode `json:"displayMode,omitempty"`
    CellOptions *common.TableCellOptions `json:"cellOptions,omitempty"`
    // ?? default is missing or false ??
    Hidden *bool `json:"hidden,omitempty"`
    Inspect bool `json:"inspect"`
    Filterable *bool `json:"filterable,omitempty"`
    // Hides any header for a column, useful for columns that show some static content or buttons.
    HideHeader *bool `json:"hideHeader,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableFieldOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableFieldOptions *TableFieldOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableFieldOptions` objects.

```go
func (tableFieldOptions *TableFieldOptions) Equals(other TableFieldOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableFieldOptions` fields for violations and returns them.

```go
func (tableFieldOptions *TableFieldOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableFieldOptionsBuilder](./builder-TableFieldOptionsBuilder.md)
