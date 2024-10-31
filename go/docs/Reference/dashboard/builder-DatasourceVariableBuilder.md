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
func (builder *DatasourceVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> AllValue

Custom all value

```go
func (builder *DatasourceVariableBuilder) AllValue(allValue string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Current

Shows current selected variable text/value on the dashboard

```go
func (builder *DatasourceVariableBuilder) Current(current dashboard.VariableOption) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *DatasourceVariableBuilder) Description(description string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Hide

Visibility configuration for the variable

```go
func (builder *DatasourceVariableBuilder) Hide(hide dashboard.VariableHide) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> IncludeAll

Whether all value option is available or not

```go
func (builder *DatasourceVariableBuilder) IncludeAll(includeAll bool) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *DatasourceVariableBuilder) Label(label string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Multi

Whether multiple values can be selected or not from variable value list

```go
func (builder *DatasourceVariableBuilder) Multi(multi bool) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *DatasourceVariableBuilder) Name(name string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```go
func (builder *DatasourceVariableBuilder) Regex(regex string) *DatasourceVariableBuilder
```

### <span class="badge object-method"></span> Type

Query used to fetch values for a variable

```go
func (builder *DatasourceVariableBuilder) Type(stringArg string) *DatasourceVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
