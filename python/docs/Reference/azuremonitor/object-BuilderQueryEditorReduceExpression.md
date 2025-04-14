---
title: <span class="badge object-type-class"></span> BuilderQueryEditorReduceExpression
---
# <span class="badge object-type-class"></span> BuilderQueryEditorReduceExpression

## Definition

```python
class BuilderQueryEditorReduceExpression:
    property_val: typing.Optional[azuremonitor.BuilderQueryEditorProperty]
    reduce: typing.Optional[azuremonitor.BuilderQueryEditorProperty]
    parameters: typing.Optional[list[azuremonitor.BuilderQueryEditorFunctionParameterExpression]]
    focus: typing.Optional[bool]
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

 * <span class="badge builder"></span> [BuilderQueryEditorReduceExpression](./builder-BuilderQueryEditorReduceExpression.md)
