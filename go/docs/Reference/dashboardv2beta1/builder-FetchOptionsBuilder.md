---
title: <span class="badge builder"></span> FetchOptionsBuilder
---
# <span class="badge builder"></span> FetchOptionsBuilder

## Constructor

```go
func NewFetchOptionsBuilder() *FetchOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FetchOptionsBuilder) Build() (FetchOptions, error)
```

### <span class="badge object-method"></span> Body

```go
func (builder *FetchOptionsBuilder) Body(body string) *FetchOptionsBuilder
```

### <span class="badge object-method"></span> Headers

```go
func (builder *FetchOptionsBuilder) Headers(headers [][]string) *FetchOptionsBuilder
```

### <span class="badge object-method"></span> Method

```go
func (builder *FetchOptionsBuilder) Method(method dashboardv2beta1.HttpRequestMethod) *FetchOptionsBuilder
```

### <span class="badge object-method"></span> QueryParams

These are 2D arrays of strings, each representing a key-value pair

We are defining them this way because we can't generate a go struct that

that would have exactly two strings in each sub-array

```go
func (builder *FetchOptionsBuilder) QueryParams(queryParams [][]string) *FetchOptionsBuilder
```

### <span class="badge object-method"></span> Url

```go
func (builder *FetchOptionsBuilder) Url(url string) *FetchOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FetchOptions](./object-FetchOptions.md)
