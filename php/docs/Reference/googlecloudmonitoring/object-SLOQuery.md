---
title: <span class="badge object-type-class"></span> SLOQuery
---
# <span class="badge object-type-class"></span> SLOQuery

SLO sub-query properties.

## Definition

```php
class SLOQuery implements \JsonSerializable
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
     * SLO selector.
     */
    public string $selectorName;

    /**
     * ID for the service the SLO is in.
     */
    public string $serviceId;

    /**
     * Name for the service the SLO is in.
     */
    public string $serviceName;

    /**
     * ID for the SLO.
     */
    public string $sloId;

    /**
     * Name of the SLO.
     */
    public string $sloName;

    /**
     * SLO goal value.
     */
    public ?float $goal;

    /**
     * Specific lookback period for the SLO.
     */
    public ?string $lookbackPeriod;

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

 * <span class="badge builder"></span> [SLOQueryBuilder](./builder-SLOQueryBuilder.md)
