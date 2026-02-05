---
title: <span class="badge builder"></span> SwitchVariableSpecBuilder
---
# <span class="badge builder"></span> SwitchVariableSpecBuilder

## Constructor

```go
func NewSwitchVariableSpecBuilder() *SwitchVariableSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SwitchVariableSpecBuilder) Build() (SwitchVariableSpec, error)
```

### <span class="badge object-method"></span> Current

```go
func (builder *SwitchVariableSpecBuilder) Current(current string) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *SwitchVariableSpecBuilder) Description(description string) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> DisabledValue

```go
func (builder *SwitchVariableSpecBuilder) DisabledValue(disabledValue string) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> EnabledValue

```go
func (builder *SwitchVariableSpecBuilder) EnabledValue(enabledValue string) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *SwitchVariableSpecBuilder) Hide(hide dashboardv2beta1.VariableHide) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *SwitchVariableSpecBuilder) Label(label string) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *SwitchVariableSpecBuilder) Name(name string) *SwitchVariableSpecBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *SwitchVariableSpecBuilder) SkipUrlSync(skipUrlSync bool) *SwitchVariableSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SwitchVariableSpec](./object-SwitchVariableSpec.md)
