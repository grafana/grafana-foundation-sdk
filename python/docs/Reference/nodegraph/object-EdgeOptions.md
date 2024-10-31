---
title: <span class="badge object-type-class"></span> EdgeOptions
---
# <span class="badge object-type-class"></span> EdgeOptions

## Definition

```python
class EdgeOptions:
    # Unit for the main stat to override what ever is set in the data frame.
    main_stat_unit: typing.Optional[str]
    # Unit for the secondary stat to override what ever is set in the data frame.
    secondary_stat_unit: typing.Optional[str]
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

 * <span class="badge builder"></span> [EdgeOptions](./builder-EdgeOptions.md)
