---
title: <span class="badge object-type-class"></span> MetricQuery
---
# <span class="badge object-type-class"></span> MetricQuery

@deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.

## Definition

```php
class MetricQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $perSeriesAligner;

    /**
     * Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $alignmentPeriod;

    /**
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public ?string $aliasBy;

    public string $editorMode;

    public string $metricType;

    /**
     * Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public string $crossSeriesReducer;

    /**
     * Array of labels to group data by.
     * @var array<string>|null
     */
    public ?array $groupBys;

    /**
     * Array of filters to query data by. Labels that can be filtered on are defined by the metric.
     * @var array<string>|null
     */
    public ?array $filters;

    public ?\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind;

    public ?string $valueType;

    public ?string $view;

    /**
     * MQL query to be executed.
     */
    public string $query;

    /**
     * Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor;

    /**
     * To disable the graphPeriod, it should explictly be set to 'disabled'.
     */
    public ?string $graphPeriod;

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

 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
