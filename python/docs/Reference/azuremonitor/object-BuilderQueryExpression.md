---
title: <span class="badge object-type-class"></span> BuilderQueryExpression
---
# <span class="badge object-type-class"></span> BuilderQueryExpression

## Definition

```python
class BuilderQueryExpression:
    from_val: typing.Optional[azuremonitor.BuilderQueryEditorPropertyExpression]
    columns: typing.Optional[azuremonitor.BuilderQueryEditorColumnsExpression]
    where: typing.Optional[azuremonitor.BuilderQueryEditorWhereExpressionArray]
    reduce: typing.Optional[azuremonitor.BuilderQueryEditorReduceExpressionArray]
    group_by: typing.Optional[azuremonitor.BuilderQueryEditorGroupByExpressionArray]
    limit: typing.Optional[int]
    order_by: typing.Optional[azuremonitor.BuilderQueryEditorOrderByExpressionArray]
    fuzzy_search: typing.Optional[azuremonitor.BuilderQueryEditorWhereExpressionArray]
    time_filter: typing.Optional[azuremonitor.BuilderQueryEditorWhereExpressionArray]
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

 * <span class="badge builder"></span> [BuilderQueryExpression](./builder-BuilderQueryExpression.md)
