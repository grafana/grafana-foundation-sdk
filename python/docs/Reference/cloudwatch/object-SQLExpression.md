---
title: <span class="badge object-type-class"></span> SQLExpression
---
# <span class="badge object-type-class"></span> SQLExpression

## Definition

```python
class SQLExpression:
    # SELECT part of the SQL expression
    select: typing.Optional[cloudwatch.QueryEditorFunctionExpression]
    # FROM part of the SQL expression
    from_val: typing.Optional[typing.Union[cloudwatch.QueryEditorPropertyExpression, cloudwatch.QueryEditorFunctionExpression]]
    # WHERE part of the SQL expression
    where: typing.Optional[cloudwatch.QueryEditorArrayExpression]
    # GROUP BY part of the SQL expression
    group_by: typing.Optional[cloudwatch.QueryEditorArrayExpression]
    # ORDER BY part of the SQL expression
    order_by: typing.Optional[cloudwatch.QueryEditorFunctionExpression]
    # The sort order of the SQL expression, `ASC` or `DESC`
    order_by_direction: typing.Optional[str]
    # LIMIT part of the SQL expression
    limit: typing.Optional[int]
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

 * <span class="badge builder"></span> [SQLExpression](./builder-SQLExpression.md)
