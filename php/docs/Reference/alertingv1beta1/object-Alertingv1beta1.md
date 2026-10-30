---
title: <span class="badge object-type-class"></span> Alertingv1beta1
---
# <span class="badge object-type-class"></span> Alertingv1beta1

## Definition

```php
class Alertingv1beta1 implements \JsonSerializable
{
}
```
## Methods

### <span class="badge object-method"></span> alertRuleManifest

Creates a resource manifest from an AlertRule.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleSpec> $folder

```php
static alertRuleManifest(string $name, \Grafana\Foundation\Cog\Builder $folder)
```

### <span class="badge object-method"></span> recordingRuleManifest

Creates a resource manifest from a RecordingRule.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RecordingRuleSpec> $folder

```php
static recordingRuleManifest(string $name, \Grafana\Foundation\Cog\Builder $folder)
```

### <span class="badge object-method"></span> ruleSequenceManifest

Creates a resource manifest from a RuleSequence.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec> $folder

```php
static ruleSequenceManifest(string $name, \Grafana\Foundation\Cog\Builder $folder)
```

