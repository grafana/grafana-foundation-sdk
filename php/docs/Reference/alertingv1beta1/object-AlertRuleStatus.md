---
title: <span class="badge object-type-class"></span> AlertRuleStatus
---
# <span class="badge object-type-class"></span> AlertRuleStatus

## Definition

```php
class AlertRuleStatus implements \JsonSerializable
{
    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleHealth $health;

    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleState $state;

    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason $stateReason;

    public ?string $lastEvaluationTime;

    /**
     * duration of the last evaluation in seconds
     */
    public ?float $evaluationDuration;

    public ?string $lastError;

    public ?\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals $totals;

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

 * <span class="badge builder"></span> [AlertRuleStatusBuilder](./builder-AlertRuleStatusBuilder.md)
