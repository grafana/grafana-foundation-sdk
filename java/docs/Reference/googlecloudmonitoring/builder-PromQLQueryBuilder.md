---
title: <span class="badge builder"></span> PromQLQueryBuilder
---
# <span class="badge builder"></span> PromQLQueryBuilder

## Constructor

```java
new PromQLQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public PromQLQuery build()
```

### <span class="badge object-method"></span> expr

PromQL expression/query to be executed.

```java
public PromQLQueryBuilder expr(String expr)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```java
public PromQLQueryBuilder projectName(String projectName)
```

### <span class="badge object-method"></span> step

PromQL min step

```java
public PromQLQueryBuilder step(String step)
```

## See also

 * <span class="badge object-type-class"></span> [PromQLQuery](./object-PromQLQuery.md)
