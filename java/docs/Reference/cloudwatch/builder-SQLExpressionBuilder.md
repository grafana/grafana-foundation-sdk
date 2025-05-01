---
title: <span class="badge builder"></span> SQLExpressionBuilder
---
# <span class="badge builder"></span> SQLExpressionBuilder

## Constructor

```java
new SQLExpressionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public SQLExpression build()
```

### <span class="badge object-method"></span> from

FROM part of the SQL expression

```java
public SQLExpressionBuilder from(QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from)
```

### <span class="badge object-method"></span> groupBy

GROUP BY part of the SQL expression

```java
public SQLExpressionBuilder groupBy(QueryEditorArrayExpression groupBy)
```

### <span class="badge object-method"></span> limit

LIMIT part of the SQL expression

```java
public SQLExpressionBuilder limit(Long limit)
```

### <span class="badge object-method"></span> orderBy

ORDER BY part of the SQL expression

```java
public SQLExpressionBuilder orderBy(QueryEditorFunctionExpression orderBy)
```

### <span class="badge object-method"></span> orderByDirection

The sort order of the SQL expression, `ASC` or `DESC`

```java
public SQLExpressionBuilder orderByDirection(String orderByDirection)
```

### <span class="badge object-method"></span> select

SELECT part of the SQL expression

```java
public SQLExpressionBuilder select(QueryEditorFunctionExpression select)
```

### <span class="badge object-method"></span> where

WHERE part of the SQL expression

```java
public SQLExpressionBuilder where(QueryEditorArrayExpression where)
```

## See also

 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
