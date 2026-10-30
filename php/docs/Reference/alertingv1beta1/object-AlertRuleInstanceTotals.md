---
title: <span class="badge object-type-class"></span> AlertRuleInstanceTotals
---
# <span class="badge object-type-class"></span> AlertRuleInstanceTotals

Count of alert instances per state. error also counts instances whose evaluation

errored but were mapped to another state via execErrState, so it can overlap.

## Definition

```php
class AlertRuleInstanceTotals implements \JsonSerializable
{
    public ?int $healthy;

    public ?int $firing;

    public ?int $pending;

    public ?int $recovering;

    public ?int $nodata;

    public ?int $error;

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

 * <span class="badge builder"></span> [AlertRuleInstanceTotalsBuilder](./builder-AlertRuleInstanceTotalsBuilder.md)
