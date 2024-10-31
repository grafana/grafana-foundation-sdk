---
title: <span class="badge builder"></span> StreamingQuery
---
# <span class="badge builder"></span> StreamingQuery

## Constructor

```python
StreamingQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> testdata.StreamingQuery
```

### <span class="badge object-method"></span> bands

```python
def bands(bands: int) -> typing.Self
```

### <span class="badge object-method"></span> noise

```python
def noise(noise: float) -> typing.Self
```

### <span class="badge object-method"></span> speed

```python
def speed(speed: float) -> typing.Self
```

### <span class="badge object-method"></span> spread

```python
def spread(spread: float) -> typing.Self
```

### <span class="badge object-method"></span> type_val

Possible enum values:

 - `"fetch"` 

 - `"logs"` 

 - `"signal"` 

 - `"traces"` 

```python
def type_val(type_val: typing.Literal["fetch", "logs", "signal", "traces"]) -> typing.Self
```

### <span class="badge object-method"></span> url

```python
def url(url: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [StreamingQuery](./object-StreamingQuery.md)
