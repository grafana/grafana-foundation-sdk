---
title: <span class="badge builder"></span> AlertRuleBuilder
---
# <span class="badge builder"></span> AlertRuleBuilder

## Constructor

```php
new AlertRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> annotations

@param array<string, string> $annotations

```php
annotations(array $annotations)
```

### <span class="badge object-method"></span> execErrState

```php
execErrState(\Grafana\Foundation\Alertingv0alpha1\ExecErrState $execErrState)
```

### <span class="badge object-method"></span> expressions

@param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\Expression>> $expressions

```php
expressions(array $expressions)
```

### <span class="badge object-method"></span> for

```php
for(string $for)
```

### <span class="badge object-method"></span> keepFiringFor

```php
keepFiringFor(string $keepFiringFor)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> missingSeriesEvalsToResolve

```php
missingSeriesEvalsToResolve(int $missingSeriesEvalsToResolve)
```

### <span class="badge object-method"></span> noDataState

```php
noDataState(\Grafana\Foundation\Alertingv0alpha1\NoDataState $noDataState)
```

### <span class="badge object-method"></span> notificationSettings

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SimplifiedRouting>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\NamedRoutingTree> $notificationSettings

```php
notificationSettings($notificationSettings)
```

### <span class="badge object-method"></span> panelRef

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\PanelRef> $panelRef

```php
panelRef(\Grafana\Foundation\Cog\Builder $panelRef)
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

 * <span class="badge object-type-class"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
