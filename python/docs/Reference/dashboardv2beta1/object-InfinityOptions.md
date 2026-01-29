---
title: <span class="badge object-type-class"></span> InfinityOptions
---
# <span class="badge object-type-class"></span> InfinityOptions

## Definition

```python
class InfinityOptions:
    method: dashboardv2beta1.HttpRequestMethod
    url: str
    body: typing.Optional[str]
    # These are 2D arrays of strings, each representing a key-value pair
    # We are defining them this way because we can't generate a go struct that
    # that would have exactly two strings in each sub-array
    query_params: typing.Optional[list[list[str]]]
    datasource_uid: str
    headers: typing.Optional[list[list[str]]]
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

 * <span class="badge builder"></span> [InfinityOptions](./builder-InfinityOptions.md)
