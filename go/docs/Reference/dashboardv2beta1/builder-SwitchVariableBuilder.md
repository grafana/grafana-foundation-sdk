---
title: <span class="badge builder"></span> SwitchVariableBuilder
---
# <span class="badge builder"></span> SwitchVariableBuilder

## Constructor

```go
func NewSwitchVariableBuilder(name string) *SwitchVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SwitchVariableBuilder) Build() (SwitchVariableKind, error)
```

### <span class="badge object-method"></span> Current

```go
func (builder *SwitchVariableBuilder) Current(current string) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *SwitchVariableBuilder) Description(description string) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> DisabledValue

```go
func (builder *SwitchVariableBuilder) DisabledValue(disabledValue string) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> EnabledValue

```go
func (builder *SwitchVariableBuilder) EnabledValue(enabledValue string) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *SwitchVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *SwitchVariableBuilder) Label(label string) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *SwitchVariableBuilder) Name(name string) *SwitchVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *SwitchVariableBuilder) SkipUrlSync(skipUrlSync bool) *SwitchVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SwitchVariableKind](./object-SwitchVariableKind.md)
