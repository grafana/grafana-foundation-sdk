---
title: <span class="badge builder"></span> CandlestickFieldMap
---
# <span class="badge builder"></span> CandlestickFieldMap

## Constructor

```python
CandlestickFieldMap()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> candlestick.CandlestickFieldMap
```

### <span class="badge object-method"></span> close

Corresponds to the final (end) value of the given period

```python
def close(close: str) -> typing.Self
```

### <span class="badge object-method"></span> high

Corresponds to the highest value of the given period

```python
def high(high: str) -> typing.Self
```

### <span class="badge object-method"></span> low

Corresponds to the lowest value of the given period

```python
def low(low: str) -> typing.Self
```

### <span class="badge object-method"></span> open_val

Corresponds to the starting value of the given period

```python
def open_val(open_val: str) -> typing.Self
```

### <span class="badge object-method"></span> volume

Corresponds to the sample count in the given period. (e.g. number of trades)

```python
def volume(volume: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CandlestickFieldMap](./object-CandlestickFieldMap.md)
