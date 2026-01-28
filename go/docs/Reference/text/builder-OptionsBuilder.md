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

### <span class="badge object-method"></span> Code

```go
func (builder *OptionsBuilder) Code(code cog.Builder[text.CodeOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Content

```go
func (builder *OptionsBuilder) Content(content string) *OptionsBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *OptionsBuilder) Mode(mode text.TextMode) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
