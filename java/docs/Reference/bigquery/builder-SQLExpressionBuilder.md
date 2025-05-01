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

### <span class="badge object-method"></span> columns

```java
public SQLExpressionBuilder columns(List<QueryEditorFunctionExpression> columns)
```

### <span class="badge object-method"></span> from

```java
public SQLExpressionBuilder from(String from)
```

### <span class="badge object-method"></span> groupBy

```java
public SQLExpressionBuilder groupBy(List<QueryEditorGroupByExpression> groupBy)
```

### <span class="badge object-method"></span> limit

```java
public SQLExpressionBuilder limit(Long limit)
```

### <span class="badge object-method"></span> offset

```java
public SQLExpressionBuilder offset(Long offset)
```

### <span class="badge object-method"></span> orderBy

```java
public SQLExpressionBuilder orderBy(QueryEditorPropertyExpression orderBy)
```

### <span class="badge object-method"></span> orderByDirection

```java
public SQLExpressionBuilder orderByDirection(OrderByDirection orderByDirection)
```

### <span class="badge object-method"></span> whereString

whereJsonTree?: _

```java
public SQLExpressionBuilder whereString(String whereString)
```

## See also

 * <span class="badge object-type-class"></span> [SQLExpression](./object-SQLExpression.md)
