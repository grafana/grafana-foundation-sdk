---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    KeepTime bool `json:"keepTime"`
    IncludeVars bool `json:"includeVars"`
    ShowStarred bool `json:"showStarred"`
    ShowRecentlyViewed bool `json:"showRecentlyViewed"`
    ShowSearch bool `json:"showSearch"`
    ShowHeadings bool `json:"showHeadings"`
    MaxItems int64 `json:"maxItems"`
    Query string `json:"query"`
    Tags []string `json:"tags"`
    // folderId is deprecated, and migrated to folderUid on panel init
    FolderId *int64 `json:"folderId,omitempty"`
    FolderUID *string `json:"folderUID,omitempty"`
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

