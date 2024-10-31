---
title: <span class="badge builder"></span> CandlestickFieldMapBuilder
---
# <span class="badge builder"></span> CandlestickFieldMapBuilder

## Constructor

```go
func NewCandlestickFieldMapBuilder() *CandlestickFieldMapBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CandlestickFieldMapBuilder) Build() (CandlestickFieldMap, error)
```

### <span class="badge object-method"></span> Close

Corresponds to the final (end) value of the given period

```go
func (builder *CandlestickFieldMapBuilder) Close(close string) *CandlestickFieldMapBuilder
```

### <span class="badge object-method"></span> High

Corresponds to the highest value of the given period

```go
func (builder *CandlestickFieldMapBuilder) High(high string) *CandlestickFieldMapBuilder
```

### <span class="badge object-method"></span> Low

Corresponds to the lowest value of the given period

```go
func (builder *CandlestickFieldMapBuilder) Low(low string) *CandlestickFieldMapBuilder
```

### <span class="badge object-method"></span> Open

Corresponds to the starting value of the given period

```go
func (builder *CandlestickFieldMapBuilder) Open(open string) *CandlestickFieldMapBuilder
```

### <span class="badge object-method"></span> Volume

Corresponds to the sample count in the given period. (e.g. number of trades)

```go
func (builder *CandlestickFieldMapBuilder) Volume(volume string) *CandlestickFieldMapBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CandlestickFieldMap](./object-CandlestickFieldMap.md)
