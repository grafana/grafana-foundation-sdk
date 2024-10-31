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
func (builder *QueryVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> AllValue

Custom all value

```go
func (builder *QueryVariableBuilder) AllValue(allValue string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Current

Shows current selected variable text/value on the dashboard

```go
func (builder *QueryVariableBuilder) Current(current dashboard.VariableOption) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```go
func (builder *QueryVariableBuilder) Datasource(datasource dashboard.DataSourceRef) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *QueryVariableBuilder) Description(description string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Hide

Visibility configuration for the variable

```go
func (builder *QueryVariableBuilder) Hide(hide dashboard.VariableHide) *QueryVariableBuilder
```

### <span class="badge object-method"></span> IncludeAll

Whether all value option is available or not

```go
func (builder *QueryVariableBuilder) IncludeAll(includeAll bool) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *QueryVariableBuilder) Label(label string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Multi

Whether multiple values can be selected or not from variable value list

```go
func (builder *QueryVariableBuilder) Multi(multi bool) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *QueryVariableBuilder) Name(name string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Options

Options that can be selected for a variable.

```go
func (builder *QueryVariableBuilder) Options(options []dashboard.VariableOption) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Query

Query used to fetch values for a variable

```go
func (builder *QueryVariableBuilder) Query(query dashboard.StringOrMap) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Refresh

Options to config when to refresh a variable

```go
func (builder *QueryVariableBuilder) Refresh(refresh dashboard.VariableRefresh) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Regex

Optional field, if you want to extract part of a series name or metric node segment.

Named capture groups can be used to separate the display text and value.

```go
func (builder *QueryVariableBuilder) Regex(regex string) *QueryVariableBuilder
```

### <span class="badge object-method"></span> Sort

Options sort order

```go
func (builder *QueryVariableBuilder) Sort(sort dashboard.VariableSort) *QueryVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
