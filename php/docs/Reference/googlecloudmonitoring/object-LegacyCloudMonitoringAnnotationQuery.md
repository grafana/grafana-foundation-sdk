---
title: <span class="badge object-type-class"></span> LegacyCloudMonitoringAnnotationQuery
---
# <span class="badge object-type-class"></span> LegacyCloudMonitoringAnnotationQuery

@deprecated Use AnnotationQuery instead. Legacy annotation query properties for migration purposes.

## Definition

```php
class LegacyCloudMonitoringAnnotationQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    public string $metricType;

    /**
     * Query refId.
     */
    public string $refId;

    /**
     * Array of filters to query data by. Labels that can be filtered on are defined by the metric.
     * @var array<string>
     */
    public array $filters;

    public \Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind;

    public string $valueType;

    /**
     * Annotation title.
     */
    public string $title;

    /**
     * Annotation text.
     */
    public string $text;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [LegacyCloudMonitoringAnnotationQueryBuilder](./builder-LegacyCloudMonitoringAnnotationQueryBuilder.md)
