---
title: <span class="badge builder"></span> NotificationPolicyBuilder
---
# <span class="badge builder"></span> NotificationPolicyBuilder

## Constructor

```php
new NotificationPolicyBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> activeTimeIntervals

@param array<string> $activeTimeIntervals

```php
activeTimeIntervals(array $activeTimeIntervals)
```

### <span class="badge object-method"></span> continue

```php
continue(bool $continue)
```

### <span class="badge object-method"></span> groupBy

@param array<string> $groupBy

```php
groupBy(array $groupBy)
```

### <span class="badge object-method"></span> groupInterval

```php
groupInterval(string $groupInterval)
```

### <span class="badge object-method"></span> groupWait

```php
groupWait(string $groupWait)
```

### <span class="badge object-method"></span> match

Deprecated. Remove before v1.0 release.

@param array<string, string> $match

```php
match(array $match)
```

### <span class="badge object-method"></span> matchRe

@param array<string, string> $matchRe

```php
matchRe(array $matchRe)
```

### <span class="badge object-method"></span> matchers

Matchers is a slice of Matchers that is sortable, implements Stringer, and

provides a Matches method to match a LabelSet against all Matchers in the

slice. Note that some users of Matchers might require it to be sorted.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Matcher>> $matchers

```php
matchers(array $matchers)
```

### <span class="badge object-method"></span> muteTimeIntervals

@param array<string> $muteTimeIntervals

```php
muteTimeIntervals(array $muteTimeIntervals)
```

### <span class="badge object-method"></span> receiver

```php
receiver(string $receiver)
```

### <span class="badge object-method"></span> repeatInterval

```php
repeatInterval(string $repeatInterval)
```

### <span class="badge object-method"></span> routes

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationPolicy>> $routes

```php
routes(array $routes)
```

## See also

 * <span class="badge object-type-class"></span> [NotificationPolicy](./object-NotificationPolicy.md)
