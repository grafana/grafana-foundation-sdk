---
title: <span class="badge builder"></span> RuleBuilder
---
# <span class="badge builder"></span> RuleBuilder

## Constructor

```php
new RuleBuilder(string $title)
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

### <span class="badge object-method"></span> condition

```php
condition(string $condition)
```

### <span class="badge object-method"></span> execErrState

```php
execErrState(\Grafana\Foundation\Alerting\RuleExecErrState $execErrState)
```

### <span class="badge object-method"></span> folderUID

```php
folderUID(string $folderUID)
```

### <span class="badge object-method"></span> for

The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.

Before this time has elapsed, the rule is only considered to be Pending.

```php
for(string $for)
```

### <span class="badge object-method"></span> id

```php
id(int $id)
```

### <span class="badge object-method"></span> isPaused

```php
isPaused(bool $isPaused)
```

### <span class="badge object-method"></span> labels

@param array<string, string> $labels

```php
labels(array $labels)
```

### <span class="badge object-method"></span> noDataState

```php
noDataState(\Grafana\Foundation\Alerting\RuleNoDataState $noDataState)
```

### <span class="badge object-method"></span> notificationSettings

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationSettings> $notificationSettings

```php
notificationSettings(\Grafana\Foundation\Cog\Builder $notificationSettings)
```

### <span class="badge object-method"></span> orgID

```php
orgID(int $orgID)
```

### <span class="badge object-method"></span> provenance

```php
provenance(string $provenance)
```

### <span class="badge object-method"></span> queries

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Query>> $data

```php
queries(array $data)
```

### <span class="badge object-method"></span> ruleGroup

```php
ruleGroup(string $ruleGroup)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

### <span class="badge object-method"></span> uid

```php
uid(string $uid)
```

### <span class="badge object-method"></span> updated

```php
updated(string $updated)
```

### <span class="badge object-method"></span> withQuery

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Query> $data

```php
withQuery(\Grafana\Foundation\Cog\Builder $data)
```

## See also

 * <span class="badge object-type-class"></span> [Rule](./object-Rule.md)
