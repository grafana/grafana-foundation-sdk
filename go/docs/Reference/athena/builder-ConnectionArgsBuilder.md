---
title: <span class="badge builder"></span> ConnectionArgsBuilder
---
# <span class="badge builder"></span> ConnectionArgsBuilder

## Constructor

```go
func NewConnectionArgsBuilder() *ConnectionArgsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConnectionArgsBuilder) Build() (ConnectionArgs, error)
```

### <span class="badge object-method"></span> Catalog

```go
func (builder *ConnectionArgsBuilder) Catalog(catalog string) *ConnectionArgsBuilder
```

### <span class="badge object-method"></span> Database

```go
func (builder *ConnectionArgsBuilder) Database(database string) *ConnectionArgsBuilder
```

### <span class="badge object-method"></span> Region

```go
func (builder *ConnectionArgsBuilder) Region(region string) *ConnectionArgsBuilder
```

### <span class="badge object-method"></span> ResultReuseEnabled

```go
func (builder *ConnectionArgsBuilder) ResultReuseEnabled(resultReuseEnabled bool) *ConnectionArgsBuilder
```

### <span class="badge object-method"></span> ResultReuseMaxAgeInMinutes

```go
func (builder *ConnectionArgsBuilder) ResultReuseMaxAgeInMinutes(resultReuseMaxAgeInMinutes float64) *ConnectionArgsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConnectionArgs](./object-ConnectionArgs.md)
