---
title: <span class="badge object-type-class"></span> AnnotationQuery
---
# <span class="badge object-type-class"></span> AnnotationQuery

Annotation sub-query properties.

## Definition

```php
class AnnotationQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public string $crossSeriesReducer;

    /**
     * Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $alignmentPeriod;

    /**
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $perSeriesAligner;

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

    /**
     * Data view, defaults to FULL.
     */
    public ?string $view;

    /**
     * Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public ?string $secondaryCrossSeriesReducer;

    /**
     * Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $secondaryAlignmentPeriod;

    /**
     * Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $secondaryPerSeriesAligner;

    /**
     * Only present if a preprocessor is selected. Array of labels to group data by.
     * @var array<string>|null
     */
    public ?array $secondaryGroupBys;

    /**
     * Annotation title.
     */
    public ?string $title;

    /**
     * Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor;

    /**
     * Annotation text.
     */
    public ?string $text;

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

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
