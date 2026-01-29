---
title: <span class="badge builder"></span> CustomFormatterVariableBuilder
---
# <span class="badge builder"></span> CustomFormatterVariableBuilder

## Constructor

```go
func NewCustomFormatterVariableBuilder() *CustomFormatterVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CustomFormatterVariableBuilder) Build() (CustomFormatterVariable, error)
```

### <span class="badge object-method"></span> IncludeAll

```go
func (builder *CustomFormatterVariableBuilder) IncludeAll(includeAll bool) *CustomFormatterVariableBuilder
```

### <span class="badge object-method"></span> Multi

```go
func (builder *CustomFormatterVariableBuilder) Multi(multi bool) *CustomFormatterVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *CustomFormatterVariableBuilder) Name(name string) *CustomFormatterVariableBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *CustomFormatterVariableBuilder) Type(typeArg dashboardv2beta1.VariableType) *CustomFormatterVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CustomFormatterVariable](./object-CustomFormatterVariable.md)
