---
title: <span class="badge object-type-struct"></span> TableFooterOptions
---
# <span class="badge object-type-struct"></span> TableFooterOptions

Footer options

## Definition

```go
type TableFooterOptions struct {
    Show bool `json:"show"`
    // actually 1 value
    Reducer []string `json:"reducer"`
    Fields []string `json:"fields,omitempty"`
    EnablePagination *bool `json:"enablePagination,omitempty"`
    CountRows *bool `json:"countRows,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableFooterOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableFooterOptions *TableFooterOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableFooterOptions` objects.

```go
func (tableFooterOptions *TableFooterOptions) Equals(other TableFooterOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableFooterOptions` fields for violations and returns them.

```go
func (tableFooterOptions *TableFooterOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableFooterOptionsBuilder](./builder-TableFooterOptionsBuilder.md)
