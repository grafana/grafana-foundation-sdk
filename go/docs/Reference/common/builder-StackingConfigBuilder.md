---
title: <span class="badge builder"></span> StackingConfigBuilder
---
# <span class="badge builder"></span> StackingConfigBuilder

## Constructor

```go
func NewStackingConfigBuilder() *StackingConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StackingConfigBuilder) Build() (StackingConfig, error)
```

### <span class="badge object-method"></span> Group

```go
func (builder *StackingConfigBuilder) Group(group string) *StackingConfigBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *StackingConfigBuilder) Mode(mode common.StackingMode) *StackingConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [StackingConfig](./object-StackingConfig.md)
