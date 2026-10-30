---
title: <span class="badge builder"></span> <span class="badge deprecated"></span> RuleGroupBuilder
---
# <span class="badge builder"></span> <span class="badge deprecated"></span> RuleGroupBuilder

<span class="badge deprecated"></span>Prefer using the alertingv1beta1 package instead

## Constructor

```php
new RuleGroupBuilder(string $title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> folderUid

```php
folderUid(string $folderUid)
```

### <span class="badge object-method"></span> interval

The interval, in seconds, at which all rules in the group are evaluated.

If a group contains many rules, the rules are evaluated sequentially.

```php
interval(int $interval)
```

### <span class="badge object-method"></span> rules

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Rule>> $rules

```php
rules(array $rules)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

### <span class="badge object-method"></span> withRule

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Rule> $rule

```php
withRule(\Grafana\Foundation\Cog\Builder $rule)
```

## See also

 * <span class="badge object-type-class"></span> [RuleGroup](./object-RuleGroup.md)
