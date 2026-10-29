---
title: <span class="badge builder"></span> RuleSpecBuilder
---
# <span class="badge builder"></span> RuleSpecBuilder

## Constructor

```php
new RuleSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> expressions

@param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\Expression>> $expressions

```php
expressions(array $expressions)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> paused

```php
paused(bool $paused)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

### <span class="badge object-method"></span> trigger

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\IntervalTrigger> $trigger

```php
trigger(\Grafana\Foundation\Cog\Builder $trigger)
```

## See also

 * <span class="badge object-type-class"></span> [RuleSpec](./object-RuleSpec.md)
