---
title: <span class="badge builder"></span> AdhocVariableBuilder
---
# <span class="badge builder"></span> AdhocVariableBuilder

## Constructor

```go
func NewAdhocVariableBuilder(name string) *AdhocVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AdhocVariableBuilder) Build() (AdhocVariableKind, error)
```

### <span class="badge object-method"></span> AllowCustomValue

```go
func (builder *AdhocVariableBuilder) AllowCustomValue(allowCustomValue bool) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> BaseFilters

```go
func (builder *AdhocVariableBuilder) BaseFilters(baseFilters []cog.Builder[dashboardv2.AdHocFilterWithLabels]) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Datasource

```go
func (builder *AdhocVariableBuilder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2AdhocVariableKindDatasource]) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> DefaultKeys

```go
func (builder *AdhocVariableBuilder) DefaultKeys(defaultKeys []cog.Builder[dashboardv2.MetricFindValue]) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *AdhocVariableBuilder) Description(description string) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> EnableGroupBy

Whether the group-by operator is enabled in the ad hoc filter combobox.

```go
func (builder *AdhocVariableBuilder) EnableGroupBy(enableGroupBy bool) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Filters

```go
func (builder *AdhocVariableBuilder) Filters(filters []cog.Builder[dashboardv2.AdHocFilterWithLabels]) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Group

```go
func (builder *AdhocVariableBuilder) Group(group string) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *AdhocVariableBuilder) Hide(hide dashboardv2.VariableHide) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *AdhocVariableBuilder) Label(label string) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *AdhocVariableBuilder) Labels(labels map[string]string) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *AdhocVariableBuilder) Name(name string) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> Origin

```go
func (builder *AdhocVariableBuilder) Origin(origin cog.Builder[dashboardv2.ControlSourceRef]) *AdhocVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *AdhocVariableBuilder) SkipUrlSync(skipUrlSync bool) *AdhocVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AdhocVariableKind](./object-AdhocVariableKind.md)
