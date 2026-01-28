---
title: <span class="badge builder"></span> TransformationBuilder
---
# <span class="badge builder"></span> TransformationBuilder

## Constructor

```go
func NewTransformationBuilder() *TransformationBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TransformationBuilder) Build() (TransformationKind, error)
```

### <span class="badge object-method"></span> Disabled

Disabled transformations are skipped

```go
func (builder *TransformationBuilder) Disabled(disabled bool) *TransformationBuilder
```

### <span class="badge object-method"></span> Filter

Optional frame matcher. When missing it will be applied to all results

```go
func (builder *TransformationBuilder) Filter(filter dashboardv2beta1.MatcherConfig) *TransformationBuilder
```

### <span class="badge object-method"></span> Id

Unique identifier of transformer

```go
func (builder *TransformationBuilder) Id(id string) *TransformationBuilder
```

### <span class="badge object-method"></span> Kind

The kind of a TransformationKind is the transformation ID

```go
func (builder *TransformationBuilder) Kind(kind string) *TransformationBuilder
```

### <span class="badge object-method"></span> Options

Options to be passed to the transformer

Valid options depend on the transformer id

```go
func (builder *TransformationBuilder) Options(options any) *TransformationBuilder
```

### <span class="badge object-method"></span> Topic

Where to pull DataFrames from as input to transformation

```go
func (builder *TransformationBuilder) Topic(topic dashboardv2beta1.DataTopic) *TransformationBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TransformationKind](./object-TransformationKind.md)
