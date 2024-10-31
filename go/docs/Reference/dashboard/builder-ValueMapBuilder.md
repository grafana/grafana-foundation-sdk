---
title: <span class="badge builder"></span> ValueMapBuilder
---
# <span class="badge builder"></span> ValueMapBuilder

## Constructor

```go
func NewValueMapBuilder() *ValueMapBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ValueMapBuilder) Build() (ValueMap, error)
```

### <span class="badge object-method"></span> Options

Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }

```go
func (builder *ValueMapBuilder) Options(options map[string]dashboard.ValueMappingResult) *ValueMapBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ValueMap](./object-ValueMap.md)
