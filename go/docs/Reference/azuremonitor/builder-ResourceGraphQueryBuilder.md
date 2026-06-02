---
title: <span class="badge builder"></span> ResourceGraphQueryBuilder
---
# <span class="badge builder"></span> ResourceGraphQueryBuilder

## Constructor

```go
func NewResourceGraphQueryBuilder() *ResourceGraphQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ResourceGraphQueryBuilder) Build() (ResourceGraphQuery, error)
```

### <span class="badge object-method"></span> Query

Azure Resource Graph KQL query to be executed.

```go
func (builder *ResourceGraphQueryBuilder) Query(query string) *ResourceGraphQueryBuilder
```

### <span class="badge object-method"></span> ResultFormat

Specifies the format results should be returned as. Defaults to table.

```go
func (builder *ResourceGraphQueryBuilder) ResultFormat(resultFormat string) *ResourceGraphQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ResourceGraphQuery](./object-ResourceGraphQuery.md)
