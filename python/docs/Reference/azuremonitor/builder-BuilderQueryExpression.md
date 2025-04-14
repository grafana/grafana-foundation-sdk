---
title: <span class="badge builder"></span> BuilderQueryExpression
---
# <span class="badge builder"></span> BuilderQueryExpression

## Constructor

```python
BuilderQueryExpression()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.BuilderQueryExpression
```

### <span class="badge object-method"></span> columns

```python
def columns(columns: cogbuilder.Builder[azuremonitor.BuilderQueryEditorColumnsExpression]) -> typing.Self
```

### <span class="badge object-method"></span> from_val

```python
def from_val(from_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorPropertyExpression]) -> typing.Self
```

### <span class="badge object-method"></span> fuzzy_search

```python
def fuzzy_search(fuzzy_search: cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) -> typing.Self
```

### <span class="badge object-method"></span> group_by

```python
def group_by(group_by: cogbuilder.Builder[azuremonitor.BuilderQueryEditorGroupByExpressionArray]) -> typing.Self
```

### <span class="badge object-method"></span> limit

```python
def limit(limit: int) -> typing.Self
```

### <span class="badge object-method"></span> order_by

```python
def order_by(order_by: cogbuilder.Builder[azuremonitor.BuilderQueryEditorOrderByExpressionArray]) -> typing.Self
```

### <span class="badge object-method"></span> reduce

```python
def reduce(reduce: cogbuilder.Builder[azuremonitor.BuilderQueryEditorReduceExpressionArray]) -> typing.Self
```

### <span class="badge object-method"></span> time_filter

```python
def time_filter(time_filter: cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) -> typing.Self
```

### <span class="badge object-method"></span> where

```python
def where(where: cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [BuilderQueryExpression](./object-BuilderQueryExpression.md)
