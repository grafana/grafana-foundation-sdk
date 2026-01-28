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

### <span class="badge object-method"></span> Counters

```go
func (builder *OptionsBuilder) Counters(counters cog.Builder[debug.UpdateConfig]) *OptionsBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *OptionsBuilder) Mode(mode debug.DebugMode) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
