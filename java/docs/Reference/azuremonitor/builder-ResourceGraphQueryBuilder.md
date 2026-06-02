---
title: <span class="badge builder"></span> ResourceGraphQueryBuilder
---
# <span class="badge builder"></span> ResourceGraphQueryBuilder

## Constructor

```java
new ResourceGraphQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public ResourceGraphQuery build()
```

### <span class="badge object-method"></span> query

Azure Resource Graph KQL query to be executed.

```java
public ResourceGraphQueryBuilder query(String query)
```

### <span class="badge object-method"></span> resultFormat

Specifies the format results should be returned as. Defaults to table.

```java
public ResourceGraphQueryBuilder resultFormat(String resultFormat)
```

## See also

 * <span class="badge object-type-class"></span> [ResourceGraphQuery](./object-ResourceGraphQuery.md)
