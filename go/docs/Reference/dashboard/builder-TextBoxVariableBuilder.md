---
title: <span class="badge builder"></span> TextBoxVariableBuilder
---
# <span class="badge builder"></span> TextBoxVariableBuilder

## Constructor

```go
func NewTextBoxVariableBuilder(name string) *TextBoxVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TextBoxVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> AllowCustomValue

Allow custom values to be entered in the variable

```go
func (builder *TextBoxVariableBuilder) AllowCustomValue(allowCustomValue bool) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Current

Shows current selected variable text/value on the dashboard

```go
func (builder *TextBoxVariableBuilder) Current(current dashboard.VariableOption) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> DefaultValue

Query used to fetch values for a variable

```go
func (builder *TextBoxVariableBuilder) DefaultValue(query dashboard.StringOrMap) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Definition

```go
func (builder *TextBoxVariableBuilder) Definition(definition string) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *TextBoxVariableBuilder) Description(description string) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Hide

Visibility configuration for the variable

```go
func (builder *TextBoxVariableBuilder) Hide(hide dashboard.VariableHide) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *TextBoxVariableBuilder) Label(label string) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *TextBoxVariableBuilder) Name(name string) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> Options

Options that can be selected for a variable.

```go
func (builder *TextBoxVariableBuilder) Options(options []dashboard.VariableOption) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> StaticOptions

Additional static options for query variable

```go
func (builder *TextBoxVariableBuilder) StaticOptions(staticOptions []dashboard.VariableOption) *TextBoxVariableBuilder
```

### <span class="badge object-method"></span> StaticOptionsOrder

Ordering of static options in relation to options returned from data source for query variable

```go
func (builder *TextBoxVariableBuilder) StaticOptionsOrder(staticOptionsOrder dashboard.VariableModelStaticOptionsOrder) *TextBoxVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
