---
title: <span class="badge builder"></span> AlertRuleStatusBuilder
---
# <span class="badge builder"></span> AlertRuleStatusBuilder

## Constructor

```php
new AlertRuleStatusBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> evaluationDuration

duration of the last evaluation in seconds

```php
evaluationDuration(float $evaluationDuration)
```

### <span class="badge object-method"></span> health

```php
health(\Grafana\Foundation\Alertingv1beta1\AlertRuleHealth $health)
```

### <span class="badge object-method"></span> lastError

```php
lastError(string $lastError)
```

### <span class="badge object-method"></span> lastEvaluationTime

```php
lastEvaluationTime(string $lastEvaluationTime)
```

### <span class="badge object-method"></span> state

```php
state(\Grafana\Foundation\Alertingv1beta1\AlertRuleState $state)
```

### <span class="badge object-method"></span> stateReason

```php
stateReason(\Grafana\Foundation\Alertingv1beta1\AlertRuleStateReason $stateReason)
```

### <span class="badge object-method"></span> totals

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals> $totals

```php
totals(\Grafana\Foundation\Cog\Builder $totals)
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleStatus](./object-AlertRuleStatus.md)
