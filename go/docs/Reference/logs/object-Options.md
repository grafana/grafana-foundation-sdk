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
    ShowControls *bool `json:"showControls,omitempty"`
    ControlsStorageKey *string `json:"controlsStorageKey,omitempty"`
    WrapLogMessage bool `json:"wrapLogMessage"`
    PrettifyLogMessage bool `json:"prettifyLogMessage"`
    EnableLogDetails bool `json:"enableLogDetails"`
    SortOrder common.LogsSortOrder `json:"sortOrder"`
    DedupStrategy common.LogsDedupStrategy `json:"dedupStrategy"`
    EnableInfiniteScrolling *bool `json:"enableInfiniteScrolling,omitempty"`
    // TODO: figure out how to define callbacks
    OnClickFilterLabel any `json:"onClickFilterLabel,omitempty"`
    OnClickFilterOutLabel any `json:"onClickFilterOutLabel,omitempty"`
    IsFilterLabelActive any `json:"isFilterLabelActive,omitempty"`
    OnClickFilterString any `json:"onClickFilterString,omitempty"`
    OnClickFilterOutString any `json:"onClickFilterOutString,omitempty"`
    OnClickShowField any `json:"onClickShowField,omitempty"`
    OnClickHideField any `json:"onClickHideField,omitempty"`
    OnLogOptionsChange any `json:"onLogOptionsChange,omitempty"`
    LogRowMenuIconsBefore any `json:"logRowMenuIconsBefore,omitempty"`
    LogRowMenuIconsAfter any `json:"logRowMenuIconsAfter,omitempty"`
    OnNewLogsReceived any `json:"onNewLogsReceived,omitempty"`
    DisplayedFields []string `json:"displayedFields,omitempty"`
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

