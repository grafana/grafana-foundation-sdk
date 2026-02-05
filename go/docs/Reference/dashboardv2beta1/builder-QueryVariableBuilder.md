---
title: <span class="badge builder"></span> QueryVariableBuilder
---
# <span class="badge builder"></span> QueryVariableBuilder

## Constructor

```go
func NewQueryVariableBuilder(name string) *QueryVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryVariableBuilder) Build() (QueryVariableKind, error)
```

### <span class="badge object-method"></span> AllValue

```go
func (builder *QueryVariableBuilder) AllValue(allValue string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> AllowCustomValue

```go
func (builder *QueryVariableBuilder) AllowCustomValue(allowCustomValue bool) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Current

```go
func (builder *QueryVariableBuilder) Current(current dashboardv2beta1.VariableOption) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Definition

```go
func (builder *QueryVariableBuilder) Definition(definition string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *QueryVariableBuilder) Description(description string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *QueryVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *QueryVariableBuilder
```

### <span class="badge object-method"></span> IncludeAll

```go
func (builder *QueryVariableBuilder) IncludeAll(includeAll bool) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *QueryVariableBuilder) Label(label string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Multi

```go
func (builder *QueryVariableBuilder) Multi(multi bool) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *QueryVariableBuilder) Name(name string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Options

```go
func (builder *QueryVariableBuilder) Options(options []dashboardv2beta1.VariableOption) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Placeholder

```go
func (builder *QueryVariableBuilder) Placeholder(placeholder string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *QueryVariableBuilder) Query(query cog.Builder[dashboardv2beta1.DataQueryKind]) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Refresh

```go
func (builder *QueryVariableBuilder) Refresh(refresh dashboardv2beta1.VariableRefresh) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Regex

```go
func (builder *QueryVariableBuilder) Regex(regex string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *QueryVariableBuilder) SkipUrlSync(skipUrlSync bool) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Sort

```go
func (builder *QueryVariableBuilder) Sort(sort dashboardv2beta1.VariableSort) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *QueryVariableBuilder) Spec(spec dashboardv2beta1.QueryVariableSpec) *QueryVariableBuilder
```

### <span class="badge object-method"></span> StaticOptions

```go
func (builder *QueryVariableBuilder) StaticOptions(staticOptions []dashboardv2beta1.VariableOption) *QueryVariableBuilder
```

### <span class="badge object-method"></span> StaticOptionsOrder

```go
func (builder *QueryVariableBuilder) StaticOptionsOrder(staticOptionsOrder dashboardv2beta1.QueryVariableSpecStaticOptionsOrder) *QueryVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryVariableKind](./object-QueryVariableKind.md)
