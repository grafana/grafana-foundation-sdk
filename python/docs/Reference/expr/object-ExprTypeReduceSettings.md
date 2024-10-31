---
title: <span class="badge object-type-class"></span> ExprTypeReduceSettings
---
# <span class="badge object-type-class"></span> ExprTypeReduceSettings

## Definition

```python
class ExprTypeReduceSettings:
    # Non-number reduce behavior
    # Possible enum values:
    #  - `"dropNN"` Drop non-numbers
    #  - `"replaceNN"` Replace non-numbers
    mode: typing.Literal["dropNN", "replaceNN"]
    # Only valid when mode is replace
    replace_with_value: typing.Optional[float]
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

 * <span class="badge builder"></span> [ExprTypeReduceSettings](./builder-ExprTypeReduceSettings.md)
