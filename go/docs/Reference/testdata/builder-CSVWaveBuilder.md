---
title: <span class="badge builder"></span> CSVWaveBuilder
---
# <span class="badge builder"></span> CSVWaveBuilder

## Constructor

```go
func NewCSVWaveBuilder() *CSVWaveBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CSVWaveBuilder) Build() (CSVWave, error)
```

### <span class="badge object-method"></span> Labels

```go
func (builder *CSVWaveBuilder) Labels(labels string) *CSVWaveBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *CSVWaveBuilder) Name(name string) *CSVWaveBuilder
```

### <span class="badge object-method"></span> TimeStep

```go
func (builder *CSVWaveBuilder) TimeStep(timeStep int64) *CSVWaveBuilder
```

### <span class="badge object-method"></span> ValuesCSV

```go
func (builder *CSVWaveBuilder) ValuesCSV(valuesCSV string) *CSVWaveBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CSVWave](./object-CSVWave.md)
