---
title: <span class="badge builder"></span> ExpressionBuilder
---
# <span class="badge builder"></span> ExpressionBuilder

## Constructor

```java
new ExpressionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Expression build()
```

### <span class="badge object-method"></span> datasourceUID

The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource

```java
public ExpressionBuilder datasourceUID(String datasourceUID)
```

### <span class="badge object-method"></span> model

```java
public ExpressionBuilder model(Object model)
```

### <span class="badge object-method"></span> queryType

The type of query if this is a query expression

```java
public ExpressionBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> relativeTimeRange

```java
public ExpressionBuilder relativeTimeRange(com.grafana.foundation.cog.Builder<RelativeTimeRange> relativeTimeRange)
```

### <span class="badge object-method"></span> source

Used to mark the expression to be used as the final source for the rule evaluation

Only one expression in a rule can be marked as the source

For AlertRules, this is the expression that will be evaluated against the alerting condition

For RecordingRules, this is the expression that will be recorded

```java
public ExpressionBuilder source(Boolean source)
```

## See also

 * <span class="badge object-type-class"></span> [Expression](./object-Expression.md)
