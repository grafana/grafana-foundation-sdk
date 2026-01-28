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

### <span class="badge object-method"></span> Limit

```go
func (builder *OptionsBuilder) Limit(limit uint32) *OptionsBuilder
```

### <span class="badge object-method"></span> NavigateAfter

```go
func (builder *OptionsBuilder) NavigateAfter(navigateAfter string) *OptionsBuilder
```

### <span class="badge object-method"></span> NavigateBefore

```go
func (builder *OptionsBuilder) NavigateBefore(navigateBefore string) *OptionsBuilder
```

### <span class="badge object-method"></span> NavigateToPanel

```go
func (builder *OptionsBuilder) NavigateToPanel(navigateToPanel bool) *OptionsBuilder
```

### <span class="badge object-method"></span> OnlyFromThisDashboard

```go
func (builder *OptionsBuilder) OnlyFromThisDashboard(onlyFromThisDashboard bool) *OptionsBuilder
```

### <span class="badge object-method"></span> OnlyInTimeRange

```go
func (builder *OptionsBuilder) OnlyInTimeRange(onlyInTimeRange bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowTags

```go
func (builder *OptionsBuilder) ShowTags(showTags bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowTime

```go
func (builder *OptionsBuilder) ShowTime(showTime bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowUser

```go
func (builder *OptionsBuilder) ShowUser(showUser bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Tags

```go
func (builder *OptionsBuilder) Tags(tags []string) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
