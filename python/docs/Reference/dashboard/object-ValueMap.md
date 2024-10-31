---
title: <span class="badge object-type-class"></span> ValueMap
---
# <span class="badge object-type-class"></span> ValueMap

Maps text values to a color or different display text and color.

For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.

## Definition

```python
class ValueMap:
    """
    Maps text values to a color or different display text and color.
    For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
    """

    type_val: typing.Literal["value"]
    # Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
    options: dict[str, dashboard.ValueMappingResult]
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

 * <span class="badge builder"></span> [ValueMap](./builder-ValueMap.md)
