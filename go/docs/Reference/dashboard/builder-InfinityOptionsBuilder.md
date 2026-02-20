---
title: <span class="badge builder"></span> InfinityOptionsBuilder
---
# <span class="badge builder"></span> InfinityOptionsBuilder

## Constructor

```go
func NewInfinityOptionsBuilder() *InfinityOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *InfinityOptionsBuilder) Build() (InfinityOptions, error)
```

### <span class="badge object-method"></span> Body

```go
func (builder *InfinityOptionsBuilder) Body(body string) *InfinityOptionsBuilder
```

### <span class="badge object-method"></span> DatasourceUid

```go
func (builder *InfinityOptionsBuilder) DatasourceUid(datasourceUid string) *InfinityOptionsBuilder
```

### <span class="badge object-method"></span> Headers

```go
func (builder *InfinityOptionsBuilder) Headers(headers [][]string) *InfinityOptionsBuilder
```

### <span class="badge object-method"></span> Method

```go
func (builder *InfinityOptionsBuilder) Method(method dashboard.HttpRequestMethod) *InfinityOptionsBuilder
```

### <span class="badge object-method"></span> QueryParams

These are 2D arrays of strings, each representing a key-value pair

We are defining them this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

```go
func (builder *InfinityOptionsBuilder) QueryParams(queryParams [][]string) *InfinityOptionsBuilder
```

### <span class="badge object-method"></span> Url

```go
func (builder *InfinityOptionsBuilder) Url(url string) *InfinityOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [InfinityOptions](./object-InfinityOptions.md)
