---
title: <span class="badge object-type-class"></span> CandlestickFieldMap
---
# <span class="badge object-type-class"></span> CandlestickFieldMap

## Definition

```python
class CandlestickFieldMap:
    # Corresponds to the starting value of the given period
    open_val: typing.Optional[str]
    # Corresponds to the highest value of the given period
    high: typing.Optional[str]
    # Corresponds to the lowest value of the given period
    low: typing.Optional[str]
    # Corresponds to the final (end) value of the given period
    close: typing.Optional[str]
    # Corresponds to the sample count in the given period. (e.g. number of trades)
    volume: typing.Optional[str]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [CandlestickFieldMap](./builder-CandlestickFieldMap.md)
