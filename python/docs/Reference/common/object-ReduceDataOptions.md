---
title: <span class="badge object-type-class"></span> ReduceDataOptions
---
# <span class="badge object-type-class"></span> ReduceDataOptions

TODO docs

## Definition

```python
class ReduceDataOptions:
    """
    TODO docs
    """

    # If true show each row value
    values: typing.Optional[bool]
    # if showing all values limit
    limit: typing.Optional[float]
    # When !values, pick one value for the whole field
    calcs: list[str]
    # Which fields to show.  By default this is only numeric fields
    fields: typing.Optional[str]
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

 * <span class="badge builder"></span> [ReduceDataOptions](./builder-ReduceDataOptions.md)
