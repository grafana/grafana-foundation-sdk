---
title: <span class="badge builder"></span> ConstantVariableBuilder
---
# <span class="badge builder"></span> ConstantVariableBuilder

## Constructor

```go
func NewConstantVariableBuilder(name string) *ConstantVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConstantVariableBuilder) Build() (ConstantVariableKind, error)
```

### <span class="badge object-method"></span> Current

```go
func (builder *ConstantVariableBuilder) Current(current dashboardv2beta1.VariableOption) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *ConstantVariableBuilder) Description(description string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *ConstantVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *ConstantVariableBuilder) Label(label string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *ConstantVariableBuilder) Name(name string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *ConstantVariableBuilder) Query(query string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *ConstantVariableBuilder) SkipUrlSync(skipUrlSync bool) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *ConstantVariableBuilder) Spec(spec dashboardv2beta1.ConstantVariableSpec) *ConstantVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConstantVariableKind](./object-ConstantVariableKind.md)
