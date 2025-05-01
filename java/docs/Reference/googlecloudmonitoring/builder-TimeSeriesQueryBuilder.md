---
title: <span class="badge builder"></span> TimeSeriesQueryBuilder
---
# <span class="badge builder"></span> TimeSeriesQueryBuilder

## Constructor

```java
new TimeSeriesQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TimeSeriesQuery build()
```

### <span class="badge object-method"></span> graphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```java
public TimeSeriesQueryBuilder graphPeriod(String graphPeriod)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```java
public TimeSeriesQueryBuilder projectName(String projectName)
```

### <span class="badge object-method"></span> query

MQL query to be executed.

```java
public TimeSeriesQueryBuilder query(String query)
```

## See also

 * <span class="badge object-type-class"></span> [TimeSeriesQuery](./object-TimeSeriesQuery.md)
