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
func (builder *ConstantVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *ConstantVariableBuilder) Description(description string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *ConstantVariableBuilder) Label(label string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *ConstantVariableBuilder) Name(name string) *ConstantVariableBuilder
```

### <span class="badge object-method"></span> Value

Query used to fetch values for a variable

```go
func (builder *ConstantVariableBuilder) Value(query dashboard.StringOrMap) *ConstantVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
