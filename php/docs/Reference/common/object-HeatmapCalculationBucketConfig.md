---
title: <span class="badge object-type-class"></span> HeatmapCalculationBucketConfig
---
# <span class="badge object-type-class"></span> HeatmapCalculationBucketConfig

## Definition

```php
class HeatmapCalculationBucketConfig implements \JsonSerializable
{
    /**
     * Sets the bucket calculation mode
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationMode $mode;

    /**
     * The number of buckets to use for the axis in the heatmap
     */
    public ?string $value;

    /**
     * Controls the scale of the buckets
     */
    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scale;

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

 * <span class="badge builder"></span> [HeatmapCalculationBucketConfigBuilder](./builder-HeatmapCalculationBucketConfigBuilder.md)
