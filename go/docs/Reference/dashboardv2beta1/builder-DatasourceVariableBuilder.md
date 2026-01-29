---
title: <span class="badge builder"></span> DatasourceVariableBuilder
---
# <span class="badge builder"></span> DatasourceVariableBuilder

## Constructor

```go
func NewDatasourceVariableBuilder(name string) *DatasourceVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DatasourceVariableBuilder) Build() (DatasourceVariableKind, error)
```

### <span class="badge object-method"></span> AllValue

```go
func (builder *DatasourceVariableBuilder) AllValue(allValue string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> AllowCustomValue

```go
func (builder *DatasourceVariableBuilder) AllowCustomValue(allowCustomValue bool) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Current

```go
func (builder *DatasourceVariableBuilder) Current(current dashboardv2beta1.VariableOption) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *DatasourceVariableBuilder) Description(description string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *DatasourceVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> IncludeAll

```go
func (builder *DatasourceVariableBuilder) IncludeAll(includeAll bool) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *DatasourceVariableBuilder) Label(label string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Multi

```go
func (builder *DatasourceVariableBuilder) Multi(multi bool) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *DatasourceVariableBuilder) Name(name string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Options

```go
func (builder *DatasourceVariableBuilder) Options(options []dashboardv2beta1.VariableOption) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> PluginId

```go
func (builder *DatasourceVariableBuilder) PluginId(pluginId string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Refresh

```go
func (builder *DatasourceVariableBuilder) Refresh(refresh dashboardv2beta1.VariableRefresh) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Regex

```go
func (builder *DatasourceVariableBuilder) Regex(regex string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *DatasourceVariableBuilder) SkipUrlSync(skipUrlSync bool) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *DatasourceVariableBuilder) Spec(spec dashboardv2beta1.DatasourceVariableSpec) *DatasourceVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DatasourceVariableKind](./object-DatasourceVariableKind.md)
