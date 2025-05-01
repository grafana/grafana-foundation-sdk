---
title: <span class="badge builder"></span> SLOQueryBuilder
---
# <span class="badge builder"></span> SLOQueryBuilder

## Constructor

```java
new SLOQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public SLOQuery build()
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```java
public SLOQueryBuilder alignmentPeriod(String alignmentPeriod)
```

### <span class="badge object-method"></span> goal

SLO goal value.

```java
public SLOQueryBuilder goal(Double goal)
```

### <span class="badge object-method"></span> lookbackPeriod

Specific lookback period for the SLO.

```java
public SLOQueryBuilder lookbackPeriod(String lookbackPeriod)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```java
public SLOQueryBuilder perSeriesAligner(String perSeriesAligner)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```java
public SLOQueryBuilder projectName(String projectName)
```

### <span class="badge object-method"></span> selectorName

SLO selector.

```java
public SLOQueryBuilder selectorName(String selectorName)
```

### <span class="badge object-method"></span> serviceId

ID for the service the SLO is in.

```java
public SLOQueryBuilder serviceId(String serviceId)
```

### <span class="badge object-method"></span> serviceName

Name for the service the SLO is in.

```java
public SLOQueryBuilder serviceName(String serviceName)
```

### <span class="badge object-method"></span> sloId

ID for the SLO.

```java
public SLOQueryBuilder sloId(String sloId)
```

### <span class="badge object-method"></span> sloName

Name of the SLO.

```java
public SLOQueryBuilder sloName(String sloName)
```

## See also

 * <span class="badge object-type-class"></span> [SLOQuery](./object-SLOQuery.md)
