---
title: <span class="badge builder"></span> UpdateConfigBuilder
---
# <span class="badge builder"></span> UpdateConfigBuilder

## Constructor

```go
func NewUpdateConfigBuilder() *UpdateConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *UpdateConfigBuilder) Build() (UpdateConfig, error)
```

### <span class="badge object-method"></span> DataChanged

```go
func (builder *UpdateConfigBuilder) DataChanged(dataChanged bool) *UpdateConfigBuilder
```

### <span class="badge object-method"></span> Render

```go
func (builder *UpdateConfigBuilder) Render(render bool) *UpdateConfigBuilder
```

### <span class="badge object-method"></span> SchemaChanged

```go
func (builder *UpdateConfigBuilder) SchemaChanged(schemaChanged bool) *UpdateConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [UpdateConfig](./object-UpdateConfig.md)
