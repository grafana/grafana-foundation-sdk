---
title: <span class="badge builder"></span> AzureResourceGraphQueryBuilder
---
# <span class="badge builder"></span> AzureResourceGraphQueryBuilder

## Constructor

```go
func NewAzureResourceGraphQueryBuilder() *AzureResourceGraphQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureResourceGraphQueryBuilder) Build() (AzureResourceGraphQuery, error)
```

### <span class="badge object-method"></span> Query

Azure Resource Graph KQL query to be executed.

```go
func (builder *AzureResourceGraphQueryBuilder) Query(query string) *AzureResourceGraphQueryBuilder
```

### <span class="badge object-method"></span> ResultFormat

Specifies the format results should be returned as. Defaults to table.

```go
func (builder *AzureResourceGraphQueryBuilder) ResultFormat(resultFormat string) *AzureResourceGraphQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureResourceGraphQuery](./object-AzureResourceGraphQuery.md)
