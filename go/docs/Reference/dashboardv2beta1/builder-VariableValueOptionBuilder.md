---
title: <span class="badge builder"></span> VariableValueOptionBuilder
---
# <span class="badge builder"></span> VariableValueOptionBuilder

## Constructor

```go
func NewVariableValueOptionBuilder() *VariableValueOptionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VariableValueOptionBuilder) Build() (VariableValueOption, error)
```

### <span class="badge object-method"></span> Group

```go
func (builder *VariableValueOptionBuilder) Group(group string) *VariableValueOptionBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *VariableValueOptionBuilder) Label(label string) *VariableValueOptionBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *VariableValueOptionBuilder) Value(value cog.Builder[dashboardv2beta1.VariableValueSingle]) *VariableValueOptionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableValueOption](./object-VariableValueOption.md)
