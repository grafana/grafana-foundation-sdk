---
title: <span class="badge builder"></span> LegacyCloudMonitoringAnnotationQueryBuilder
---
# <span class="badge builder"></span> LegacyCloudMonitoringAnnotationQueryBuilder

## Constructor

```php
new LegacyCloudMonitoringAnnotationQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

@param array<string> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> metricKind

```php
metricKind(\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind)
```

### <span class="badge object-method"></span> metricType

```php
metricType(string $metricType)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```php
projectName(string $projectName)
```

### <span class="badge object-method"></span> refId

Query refId.

```php
refId(string $refId)
```

### <span class="badge object-method"></span> text

Annotation text.

```php
text(string $text)
```

### <span class="badge object-method"></span> title

Annotation title.

```php
title(string $title)
```

### <span class="badge object-method"></span> valueType

```php
valueType(string $valueType)
```

## See also

 * <span class="badge object-type-class"></span> [LegacyCloudMonitoringAnnotationQuery](./object-LegacyCloudMonitoringAnnotationQuery.md)
