---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    ShowLabels bool `json:"showLabels"`
    ShowCommonLabels bool `json:"showCommonLabels"`
    ShowTime bool `json:"showTime"`
    ShowLogContextToggle bool `json:"showLogContextToggle"`
    WrapLogMessage bool `json:"wrapLogMessage"`
    PrettifyLogMessage bool `json:"prettifyLogMessage"`
    EnableLogDetails bool `json:"enableLogDetails"`
    SortOrder common.LogsSortOrder `json:"sortOrder"`
    DedupStrategy common.LogsDedupStrategy `json:"dedupStrategy"`
    EnableInfiniteScrolling *bool `json:"enableInfiniteScrolling,omitempty"`
    // Display controls to jump to the last or first log line, and filters by log level.
    ShowControls *bool `json:"showControls,omitempty"`
    // Show a component to manage the displayed fields from the logs.
    ShowFieldSelector *bool `json:"showFieldSelector,omitempty"`
    // Use a predefined coloring scheme to highlight relevant parts of the log lines.
    SyntaxHighlighting *bool `json:"syntaxHighlighting,omitempty"`
    FontSize *logs.OptionsFontSize `json:"fontSize,omitempty"`
    DetailsMode *logs.OptionsDetailsMode `json:"detailsMode,omitempty"`
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

