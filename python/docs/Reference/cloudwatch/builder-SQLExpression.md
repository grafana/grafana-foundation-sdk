---
title: <span class="badge builder"></span> SQLExpression
---
# <span class="badge builder"></span> SQLExpression

## Constructor

```python
SQLExpression()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> cloudwatch.SQLExpression
```

### <span class="badge object-method"></span> from_val

FROM part of the SQL expression

```python
def from_val(from_val: typing.Union[cogbuilder.Builder[cloudwatch.QueryEditorPropertyExpression], cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]]) -> typing.Self
```

### <span class="badge object-method"></span> group_by

GROUP BY part of the SQL expression

```python
def group_by(group_by: cogbuilder.Builder[cloudwatch.QueryEditorArrayExpression]) -> typing.Self
```

### <span class="badge object-method"></span> limit

LIMIT part of the SQL expression

```python
def limit(limit: int) -> typing.Self
```

### <span class="badge object-method"></span> order_by

ORDER BY part of the SQL expression

```python
def order_by(order_by: cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]) -> typing.Self
```

### <span class="badge object-method"></span> order_by_direction

The sort order of the SQL expression, `ASC` or `DESC`

```python
def order_by_direction(order_by_direction: str) -> typing.Self
```

### <span class="badge object-method"></span> select

SELECT part of the SQL expression

```python
def select(select: cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]) -> typing.Self
```

### <span class="badge object-method"></span> where

WHERE part of the SQL expression

```python
def where(where: cogbuilder.Builder[cloudwatch.QueryEditorArrayExpression]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
