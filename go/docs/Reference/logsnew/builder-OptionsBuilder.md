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

### <span class="badge object-method"></span> DedupStrategy

```go
func (builder *OptionsBuilder) DedupStrategy(dedupStrategy common.LogsDedupStrategy) *OptionsBuilder
```

### <span class="badge object-method"></span> EnableInfiniteScrolling

```go
func (builder *OptionsBuilder) EnableInfiniteScrolling(enableInfiniteScrolling bool) *OptionsBuilder
```

### <span class="badge object-method"></span> EnableLogDetails

```go
func (builder *OptionsBuilder) EnableLogDetails(enableLogDetails bool) *OptionsBuilder
```

### <span class="badge object-method"></span> OnNewLogsReceived

```go
func (builder *OptionsBuilder) OnNewLogsReceived(onNewLogsReceived any) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowTime

```go
func (builder *OptionsBuilder) ShowTime(showTime bool) *OptionsBuilder
```

### <span class="badge object-method"></span> SortOrder

```go
func (builder *OptionsBuilder) SortOrder(sortOrder common.LogsSortOrder) *OptionsBuilder
```

### <span class="badge object-method"></span> WrapLogMessage

```go
func (builder *OptionsBuilder) WrapLogMessage(wrapLogMessage bool) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
