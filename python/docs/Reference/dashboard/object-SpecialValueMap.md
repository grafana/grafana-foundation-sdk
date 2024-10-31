---
title: <span class="badge object-type-class"></span> SpecialValueMap
---
# <span class="badge object-type-class"></span> SpecialValueMap

Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.

See SpecialValueMatch to see the list of special values.

For example, you can configure a special value mapping so that null values appear as N/A.

## Definition

```python
class SpecialValueMap:
    """
    Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
    See SpecialValueMatch to see the list of special values.
    For example, you can configure a special value mapping so that null values appear as N/A.
    """

    type_val: typing.Literal["special"]
    options: dashboard.DashboardSpecialValueMapOptions
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

 * <span class="badge builder"></span> [SpecialValueMap](./builder-SpecialValueMap.md)
