---
title: <span class="badge builder"></span> SimulationQueryBuilder
---
# <span class="badge builder"></span> SimulationQueryBuilder

## Constructor

```go
func NewSimulationQueryBuilder() *SimulationQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SimulationQueryBuilder) Build() (SimulationQuery, error)
```

### <span class="badge object-method"></span> Config

```go
func (builder *SimulationQueryBuilder) Config(config map[string]any) *SimulationQueryBuilder
```

### <span class="badge object-method"></span> Key

```go
func (builder *SimulationQueryBuilder) Key(key cog.Builder[testdata.Key]) *SimulationQueryBuilder
```

### <span class="badge object-method"></span> Last

```go
func (builder *SimulationQueryBuilder) Last(last bool) *SimulationQueryBuilder
```

### <span class="badge object-method"></span> Stream

```go
func (builder *SimulationQueryBuilder) Stream(stream bool) *SimulationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SimulationQuery](./object-SimulationQuery.md)
