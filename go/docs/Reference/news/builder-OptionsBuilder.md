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

### <span class="badge object-method"></span> FeedUrl

empty/missing will default to grafana blog

```go
func (builder *OptionsBuilder) FeedUrl(feedUrl string) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowImage

```go
func (builder *OptionsBuilder) ShowImage(showImage bool) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
