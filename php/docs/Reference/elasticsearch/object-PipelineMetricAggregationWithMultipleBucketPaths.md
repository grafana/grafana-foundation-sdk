---
title: <span class="badge object-type-class"></span> PipelineMetricAggregationWithMultipleBucketPaths
---
# <span class="badge object-type-class"></span> PipelineMetricAggregationWithMultipleBucketPaths

## Definition

```php
class PipelineMetricAggregationWithMultipleBucketPaths implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Elasticsearch\PipelineVariable>|null
     */
    public ?array $pipelineVariables;

    /**
     * @var string|\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType
     */
    public $type;

    public string $id;

    public ?bool $hide;

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

 * <span class="badge builder"></span> [PipelineMetricAggregationWithMultipleBucketPathsBuilder](./builder-PipelineMetricAggregationWithMultipleBucketPathsBuilder.md)
