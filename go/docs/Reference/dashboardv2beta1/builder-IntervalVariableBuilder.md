---
title: <span class="badge builder"></span> IntervalVariableBuilder
---
# <span class="badge builder"></span> IntervalVariableBuilder

## Constructor

```go
func NewIntervalVariableBuilder(name string) *IntervalVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *IntervalVariableBuilder) Build() (IntervalVariableKind, error)
```

### <span class="badge object-method"></span> Auto

```go
func (builder *IntervalVariableBuilder) Auto(auto bool) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> AutoCount

```go
func (builder *IntervalVariableBuilder) AutoCount(autoCount int64) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> AutoMin

```go
func (builder *IntervalVariableBuilder) AutoMin(autoMin string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Current

```go
func (builder *IntervalVariableBuilder) Current(current dashboardv2beta1.VariableOption) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Description

```go
func (builder *IntervalVariableBuilder) Description(description string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *IntervalVariableBuilder) Hide(hide dashboardv2beta1.VariableHide) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *IntervalVariableBuilder) Label(label string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *IntervalVariableBuilder) Name(name string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Options

```go
func (builder *IntervalVariableBuilder) Options(options []dashboardv2beta1.VariableOption) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *IntervalVariableBuilder) Query(query string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Refresh

```go
func (builder *IntervalVariableBuilder) Refresh(refresh dashboardv2beta1.VariableRefresh) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> SkipUrlSync

```go
func (builder *IntervalVariableBuilder) SkipUrlSync(skipUrlSync bool) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *IntervalVariableBuilder) Spec(spec dashboardv2beta1.IntervalVariableSpec) *IntervalVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [IntervalVariableKind](./object-IntervalVariableKind.md)
