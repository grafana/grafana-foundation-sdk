---
title: <span class="badge builder"></span> ValueMappingResultBuilder
---
# <span class="badge builder"></span> ValueMappingResultBuilder

## Constructor

```go
func NewValueMappingResultBuilder() *ValueMappingResultBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ValueMappingResultBuilder) Build() (ValueMappingResult, error)
```

### <span class="badge object-method"></span> Color

Text to use when the value matches

```go
func (builder *ValueMappingResultBuilder) Color(color string) *ValueMappingResultBuilder
```

### <span class="badge object-method"></span> Icon

Icon to display when the value matches. Only specific visualizations.

```go
func (builder *ValueMappingResultBuilder) Icon(icon string) *ValueMappingResultBuilder
```

### <span class="badge object-method"></span> Index

Position in the mapping array. Only used internally.

```go
func (builder *ValueMappingResultBuilder) Index(index int32) *ValueMappingResultBuilder
```

### <span class="badge object-method"></span> Text

Text to display when the value matches

```go
func (builder *ValueMappingResultBuilder) Text(text string) *ValueMappingResultBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ValueMappingResult](./object-ValueMappingResult.md)
