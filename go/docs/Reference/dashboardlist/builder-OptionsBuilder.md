---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> FolderId

folderId is deprecated, and migrated to folderUid on panel init

```go
func (builder *OptionsBuilder) FolderId(folderId int64) *OptionsBuilder
```

### <span class="badge object-method"></span> FolderUID

```go
func (builder *OptionsBuilder) FolderUID(folderUID string) *OptionsBuilder
```

### <span class="badge object-method"></span> IncludeVars

```go
func (builder *OptionsBuilder) IncludeVars(includeVars bool) *OptionsBuilder
```

### <span class="badge object-method"></span> KeepTime

```go
func (builder *OptionsBuilder) KeepTime(keepTime bool) *OptionsBuilder
```

### <span class="badge object-method"></span> MaxItems

```go
func (builder *OptionsBuilder) MaxItems(maxItems int64) *OptionsBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *OptionsBuilder) Query(query string) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowHeadings

```go
func (builder *OptionsBuilder) ShowHeadings(showHeadings bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowRecentlyViewed

```go
func (builder *OptionsBuilder) ShowRecentlyViewed(showRecentlyViewed bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowSearch

```go
func (builder *OptionsBuilder) ShowSearch(showSearch bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowStarred

```go
func (builder *OptionsBuilder) ShowStarred(showStarred bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Tags

```go
func (builder *OptionsBuilder) Tags(tags []string) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
