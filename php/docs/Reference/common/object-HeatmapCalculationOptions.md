---
title: <span class="badge object-type-class"></span> HeatmapCalculationOptions
---
# <span class="badge object-type-class"></span> HeatmapCalculationOptions

## Definition

```php
class HeatmapCalculationOptions implements \JsonSerializable
{
    /**
     * The number of buckets to use for the xAxis in the heatmap
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $xBuckets;

    /**
     * The number of buckets to use for the yAxis in the heatmap
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $yBuckets;

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

 * <span class="badge builder"></span> [HeatmapCalculationOptionsBuilder](./builder-HeatmapCalculationOptionsBuilder.md)
