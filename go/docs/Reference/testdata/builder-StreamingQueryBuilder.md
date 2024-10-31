---
title: <span class="badge builder"></span> StreamingQueryBuilder
---
# <span class="badge builder"></span> StreamingQueryBuilder

## Constructor

```go
func NewStreamingQueryBuilder() *StreamingQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StreamingQueryBuilder) Build() (StreamingQuery, error)
```

### <span class="badge object-method"></span> Bands

```go
func (builder *StreamingQueryBuilder) Bands(bands int64) *StreamingQueryBuilder
```

### <span class="badge object-method"></span> Noise

```go
func (builder *StreamingQueryBuilder) Noise(noise float64) *StreamingQueryBuilder
```

### <span class="badge object-method"></span> Speed

```go
func (builder *StreamingQueryBuilder) Speed(speed float64) *StreamingQueryBuilder
```

### <span class="badge object-method"></span> Spread

```go
func (builder *StreamingQueryBuilder) Spread(spread float64) *StreamingQueryBuilder
```

### <span class="badge object-method"></span> Type

Possible enum values:

 - `"fetch"` 

 - `"logs"` 

 - `"signal"` 

 - `"traces"` 

```go
func (builder *StreamingQueryBuilder) Type(typeArg testdata.StreamingQueryType) *StreamingQueryBuilder
```

### <span class="badge object-method"></span> Url

```go
func (builder *StreamingQueryBuilder) Url(url string) *StreamingQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [StreamingQuery](./object-StreamingQuery.md)
