---
title: <span class="badge builder"></span> CodeOptionsBuilder
---
# <span class="badge builder"></span> CodeOptionsBuilder

## Constructor

```go
func NewCodeOptionsBuilder() *CodeOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CodeOptionsBuilder) Build() (CodeOptions, error)
```

### <span class="badge object-method"></span> Language

The language passed to monaco code editor

```go
func (builder *CodeOptionsBuilder) Language(language text.CodeLanguage) *CodeOptionsBuilder
```

### <span class="badge object-method"></span> ShowLineNumbers

```go
func (builder *CodeOptionsBuilder) ShowLineNumbers(showLineNumbers bool) *CodeOptionsBuilder
```

### <span class="badge object-method"></span> ShowMiniMap

```go
func (builder *CodeOptionsBuilder) ShowMiniMap(showMiniMap bool) *CodeOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CodeOptions](./object-CodeOptions.md)
