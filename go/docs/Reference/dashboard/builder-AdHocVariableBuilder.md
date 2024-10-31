---
title: <span class="badge builder"></span> AdHocVariableBuilder
---
# <span class="badge builder"></span> AdHocVariableBuilder

## Constructor

```go
func NewAdHocVariableBuilder(name string) *AdHocVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AdHocVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> AllFormat

Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.

```go
func (builder *AdHocVariableBuilder) AllFormat(allFormat string) *AdHocVariableBuilder
```

### <span class="badge object-method"></span> Datasource

Data source used to fetch values for a variable. It can be defined but `null`.

```go
func (builder *AdHocVariableBuilder) Datasource(datasource dashboard.DataSourceRef) *AdHocVariableBuilder
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *AdHocVariableBuilder) Description(description string) *AdHocVariableBuilder
```

### <span class="badge object-method"></span> Hide

Visibility configuration for the variable

```go
func (builder *AdHocVariableBuilder) Hide(hide dashboard.VariableHide) *AdHocVariableBuilder
```

### <span class="badge object-method"></span> Id

Unique numeric identifier for the variable.

```go
func (builder *AdHocVariableBuilder) Id(id string) *AdHocVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *AdHocVariableBuilder) Label(label string) *AdHocVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *AdHocVariableBuilder) Name(name string) *AdHocVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
