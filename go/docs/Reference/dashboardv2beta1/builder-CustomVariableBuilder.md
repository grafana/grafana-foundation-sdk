---
title: <span class="badge builder"></span> CustomVariableBuilder
---
# <span class="badge builder"></span> CustomVariableBuilder

## Constructor

```go
func NewCustomVariableBuilder(name string) *CustomVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CustomVariableBuilder) Build() (CustomVariableKind, error)
```

### <span class="badge object-method"></span> AllValue

```go
func (builder *CustomVariableBuilder) AllValue(allValue string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> AllowCustomValue

```go
func (builder *CustomVariableBuilder) AllowCustomValue(allowCustomValue bool) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Current

```go
func (builder *CustomVariableBuilder) Current(current dashboardv2beta1.VariableOption) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *CustomVariableBuilder) Description(description string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *CustomVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *CustomVariableBuilder
```

### <span class="badge object-method"></span> IncludeAll

```go
func (builder *CustomVariableBuilder) IncludeAll(includeAll bool) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *CustomVariableBuilder) Label(label string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Multi

```go
func (builder *CustomVariableBuilder) Multi(multi bool) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *CustomVariableBuilder) Name(name string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Options

```go
func (builder *CustomVariableBuilder) Options(options []dashboardv2beta1.VariableOption) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *CustomVariableBuilder) Query(query string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *CustomVariableBuilder) SkipUrlSync(skipUrlSync bool) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *CustomVariableBuilder) Spec(spec dashboardv2beta1.CustomVariableSpec) *CustomVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CustomVariableKind](./object-CustomVariableKind.md)
