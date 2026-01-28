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

### <span class="badge object-method"></span> EnableLogDetails

```go
func (builder *OptionsBuilder) EnableLogDetails(enableLogDetails bool) *OptionsBuilder
```

### <span class="badge object-method"></span> PrettifyLogMessage

```go
func (builder *OptionsBuilder) PrettifyLogMessage(prettifyLogMessage bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowCommonLabels

```go
func (builder *OptionsBuilder) ShowCommonLabels(showCommonLabels bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowLabels

```go
func (builder *OptionsBuilder) ShowLabels(showLabels bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowLogContextToggle

```go
func (builder *OptionsBuilder) ShowLogContextToggle(showLogContextToggle bool) *OptionsBuilder
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
