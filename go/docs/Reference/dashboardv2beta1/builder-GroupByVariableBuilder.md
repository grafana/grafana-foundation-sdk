---
title: <span class="badge builder"></span> GroupByVariableBuilder
---
# <span class="badge builder"></span> GroupByVariableBuilder

## Constructor

```go
func NewGroupByVariableBuilder(name string) *GroupByVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GroupByVariableBuilder) Build() (GroupByVariableKind, error)
```

### <span class="badge object-method"></span> Current

```go
func (builder *GroupByVariableBuilder) Current(current dashboardv2beta1.VariableOption) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Datasource

```go
func (builder *GroupByVariableBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource]) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> DefaultValue

```go
func (builder *GroupByVariableBuilder) DefaultValue(defaultValue dashboardv2beta1.VariableOption) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *GroupByVariableBuilder) Description(description string) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Group

```go
func (builder *GroupByVariableBuilder) Group(group string) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *GroupByVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *GroupByVariableBuilder) Label(label string) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Multi

```go
func (builder *GroupByVariableBuilder) Multi(multi bool) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *GroupByVariableBuilder) Name(name string) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> Options

```go
func (builder *GroupByVariableBuilder) Options(options []dashboardv2beta1.VariableOption) *GroupByVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *GroupByVariableBuilder) SkipUrlSync(skipUrlSync bool) *GroupByVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GroupByVariableKind](./object-GroupByVariableKind.md)
