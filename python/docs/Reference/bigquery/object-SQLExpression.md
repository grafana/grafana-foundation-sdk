---
title: <span class="badge object-type-class"></span> SQLExpression
---
# <span class="badge object-type-class"></span> SQLExpression

## Definition

```python
class SQLExpression:
    columns: typing.Optional[list[bigquery.QueryEditorFunctionExpression]]
    from_val: typing.Optional[str]
    # whereJsonTree?: _
    where_string: typing.Optional[str]
    group_by: typing.Optional[list[bigquery.QueryEditorGroupByExpression]]
    order_by: typing.Optional[bigquery.QueryEditorPropertyExpression]
    order_by_direction: typing.Optional[bigquery.OrderByDirection]
    limit: typing.Optional[int]
    offset: typing.Optional[int]
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
