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
def build() -> bigquery.SQLExpression
```

### <span class="badge object-method"></span> columns

```python
def columns(columns: list[cogbuilder.Builder[bigquery.QueryEditorFunctionExpression]]) -> typing.Self
```

### <span class="badge object-method"></span> from_val

```python
def from_val(from_val: str) -> typing.Self
```

### <span class="badge object-method"></span> group_by

```python
def group_by(group_by: list[cogbuilder.Builder[bigquery.QueryEditorGroupByExpression]]) -> typing.Self
```

### <span class="badge object-method"></span> limit

```python
def limit(limit: int) -> typing.Self
```

### <span class="badge object-method"></span> offset

```python
def offset(offset: int) -> typing.Self
```

### <span class="badge object-method"></span> order_by

```python
def order_by(order_by: cogbuilder.Builder[bigquery.QueryEditorPropertyExpression]) -> typing.Self
```

### <span class="badge object-method"></span> order_by_direction

```python
def order_by_direction(order_by_direction: bigquery.OrderByDirection) -> typing.Self
```

### <span class="badge object-method"></span> where_string

whereJsonTree?: _

```python
def where_string(where_string: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
