---
title: <span class="badge builder"></span> NodesQueryBuilder
---
# <span class="badge builder"></span> NodesQueryBuilder

## Constructor

```go
func NewNodesQueryBuilder() *NodesQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NodesQueryBuilder) Build() (NodesQuery, error)
```

### <span class="badge object-method"></span> Count

```go
func (builder *NodesQueryBuilder) Count(count int64) *NodesQueryBuilder
```

### <span class="badge object-method"></span> Seed

```go
func (builder *NodesQueryBuilder) Seed(seed int64) *NodesQueryBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *NodesQueryBuilder) Type(typeArg testdata.NodesQueryType) *NodesQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NodesQuery](./object-NodesQuery.md)
