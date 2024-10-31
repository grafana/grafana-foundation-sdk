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
func (builder *CustomVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> AllFormat

Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.

```go
func (builder *CustomVariableBuilder) AllFormat(allFormat string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> AllValue

Custom all value

```go
func (builder *CustomVariableBuilder) AllValue(allValue string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Current

Shows current selected variable text/value on the dashboard

```go
func (builder *CustomVariableBuilder) Current(current dashboard.VariableOption) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *CustomVariableBuilder) Description(description string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Hide

Visibility configuration for the variable

```go
func (builder *CustomVariableBuilder) Hide(hide dashboard.VariableHide) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Id

Unique numeric identifier for the variable.

```go
func (builder *CustomVariableBuilder) Id(id string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> IncludeAll

Whether all value option is available or not

```go
func (builder *CustomVariableBuilder) IncludeAll(includeAll bool) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *CustomVariableBuilder) Label(label string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Multi

Whether multiple values can be selected or not from variable value list

```go
func (builder *CustomVariableBuilder) Multi(multi bool) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *CustomVariableBuilder) Name(name string) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Options

Options that can be selected for a variable.

```go
func (builder *CustomVariableBuilder) Options(options []dashboard.VariableOption) *CustomVariableBuilder
```

### <span class="badge object-method"></span> Values

Query used to fetch values for a variable

```go
func (builder *CustomVariableBuilder) Values(query dashboard.StringOrMap) *CustomVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
