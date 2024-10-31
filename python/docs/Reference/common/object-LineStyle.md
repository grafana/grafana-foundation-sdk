---
title: <span class="badge object-type-class"></span> LineStyle
---
# <span class="badge object-type-class"></span> LineStyle

TODO docs

## Definition

```python
class LineStyle:
    """
    TODO docs
    """

    fill: typing.Optional[typing.Literal["solid", "dash", "dot", "square"]]
    dash: typing.Optional[list[float]]
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

 * <span class="badge builder"></span> [LineStyle](./builder-LineStyle.md)
