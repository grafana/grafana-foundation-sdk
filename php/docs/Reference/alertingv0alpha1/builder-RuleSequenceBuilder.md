---
title: <span class="badge builder"></span> RuleSequenceBuilder
---
# <span class="badge builder"></span> RuleSequenceBuilder

## Constructor

```php
new RuleSequenceBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> alertingRules

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleRef>> $alertingRules

```php
alertingRules(array $alertingRules)
```

### <span class="badge object-method"></span> recordingRules

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleRef>> $recordingRules

```php
recordingRules(array $recordingRules)
```

### <span class="badge object-method"></span> trigger

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\IntervalTrigger> $trigger

```php
trigger(\Grafana\Foundation\Cog\Builder $trigger)
```

## See also

 * <span class="badge object-type-class"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
