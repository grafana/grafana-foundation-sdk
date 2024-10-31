---
title: <span class="badge builder"></span> SLOQueryBuilder
---
# <span class="badge builder"></span> SLOQueryBuilder

## Constructor

```php
new SLOQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```php
alignmentPeriod(string $alignmentPeriod)
```

### <span class="badge object-method"></span> goal

SLO goal value.

```php
goal(float $goal)
```

### <span class="badge object-method"></span> lookbackPeriod

Specific lookback period for the SLO.

```php
lookbackPeriod(string $lookbackPeriod)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```php
perSeriesAligner(string $perSeriesAligner)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```php
projectName(string $projectName)
```

### <span class="badge object-method"></span> selectorName

SLO selector.

```php
selectorName(string $selectorName)
```

### <span class="badge object-method"></span> serviceId

ID for the service the SLO is in.

```php
serviceId(string $serviceId)
```

### <span class="badge object-method"></span> serviceName

Name for the service the SLO is in.

```php
serviceName(string $serviceName)
```

### <span class="badge object-method"></span> sloId

ID for the SLO.

```php
sloId(string $sloId)
```

### <span class="badge object-method"></span> sloName

Name of the SLO.

```php
sloName(string $sloName)
```

## See also

 * <span class="badge object-type-class"></span> [SLOQuery](./object-SLOQuery.md)
