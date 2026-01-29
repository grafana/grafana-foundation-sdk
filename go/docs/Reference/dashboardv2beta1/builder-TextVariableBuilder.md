---
title: <span class="badge builder"></span> TextVariableBuilder
---
# <span class="badge builder"></span> TextVariableBuilder

## Constructor

```go
func NewTextVariableBuilder(name string) *TextVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TextVariableBuilder) Build() (TextVariableKind, error)
```

### <span class="badge object-method"></span> Current

```go
func (builder *TextVariableBuilder) Current(current dashboardv2beta1.VariableOption) *TextVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *TextVariableBuilder) Description(description string) *TextVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *TextVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *TextVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *TextVariableBuilder) Label(label string) *TextVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *TextVariableBuilder) Name(name string) *TextVariableBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *TextVariableBuilder) Query(query string) *TextVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *TextVariableBuilder) SkipUrlSync(skipUrlSync bool) *TextVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *TextVariableBuilder) Spec(spec dashboardv2beta1.TextVariableSpec) *TextVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TextVariableKind](./object-TextVariableKind.md)
