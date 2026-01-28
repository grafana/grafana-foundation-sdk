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

### <span class="badge object-method"></span> DisplayedFields

```go
func (builder *OptionsBuilder) DisplayedFields(displayedFields []string) *OptionsBuilder
```

### <span class="badge object-method"></span> EnableLogDetails

```go
func (builder *OptionsBuilder) EnableLogDetails(enableLogDetails bool) *OptionsBuilder
```

### <span class="badge object-method"></span> IsFilterLabelActive

```go
func (builder *OptionsBuilder) IsFilterLabelActive(isFilterLabelActive any) *OptionsBuilder
```

### <span class="badge object-method"></span> OnClickFilterLabel

TODO: figure out how to define callbacks

```go
func (builder *OptionsBuilder) OnClickFilterLabel(onClickFilterLabel any) *OptionsBuilder
```

### <span class="badge object-method"></span> OnClickFilterOutLabel

```go
func (builder *OptionsBuilder) OnClickFilterOutLabel(onClickFilterOutLabel any) *OptionsBuilder
```

### <span class="badge object-method"></span> OnClickFilterOutString

```go
func (builder *OptionsBuilder) OnClickFilterOutString(onClickFilterOutString any) *OptionsBuilder
```

### <span class="badge object-method"></span> OnClickFilterString

```go
func (builder *OptionsBuilder) OnClickFilterString(onClickFilterString any) *OptionsBuilder
```

### <span class="badge object-method"></span> OnClickHideField

```go
func (builder *OptionsBuilder) OnClickHideField(onClickHideField any) *OptionsBuilder
```

### <span class="badge object-method"></span> OnClickShowField

```go
func (builder *OptionsBuilder) OnClickShowField(onClickShowField any) *OptionsBuilder
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
