---
title: <span class="badge builder"></span> V2TransformationSpecBuilder
---
# <span class="badge builder"></span> V2TransformationSpecBuilder

Dashboard v2 shape: no `id`, it is carried by the parent's `group`.

## Constructor

```go
func NewV2TransformationSpecBuilder() *V2TransformationSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *V2TransformationSpecBuilder) Build() (V2TransformationSpec, error)
```

### <span class="badge object-method"></span> Disabled

Disabled transformations are skipped

```go
func (builder *V2TransformationSpecBuilder) Disabled(disabled bool) *V2TransformationSpecBuilder
```

### <span class="badge object-method"></span> Filter

Optional frame matcher. When missing it will be applied to all results

```go
func (builder *V2TransformationSpecBuilder) Filter(filter dashboardv2beta1.MatcherConfig) *V2TransformationSpecBuilder
```

### <span class="badge object-method"></span> Options

Options to be passed to the transformer

Valid options depend on the transformer id

```go
func (builder *V2TransformationSpecBuilder) Options(options any) *V2TransformationSpecBuilder
```

### <span class="badge object-method"></span> Topic

Where to pull DataFrames from as input to transformation

```go
func (builder *V2TransformationSpecBuilder) Topic(topic dashboardv2beta1.DataTopic) *V2TransformationSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [V2TransformationSpec](./object-V2TransformationSpec.md)
