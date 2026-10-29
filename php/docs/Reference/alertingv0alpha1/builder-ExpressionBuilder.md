---
title: <span class="badge builder"></span> ExpressionBuilder
---
# <span class="badge builder"></span> ExpressionBuilder

## Constructor

```php
new ExpressionBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasourceUID

The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource

```php
datasourceUID(string $datasourceUID)
```

### <span class="badge object-method"></span> model

@param mixed $model

```php
model($model)
```

### <span class="badge object-method"></span> queryType

The type of query if this is a query expression

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> relativeTimeRange

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange> $relativeTimeRange

```php
relativeTimeRange(\Grafana\Foundation\Cog\Builder $relativeTimeRange)
```

### <span class="badge object-method"></span> source

Used to mark the expression to be used as the final source for the rule evaluation

Only one expression in a rule can be marked as the source

For AlertRules, this is the expression that will be evaluated against the alerting condition

For RecordingRules, this is the expression that will be recorded

```php
source(bool $source)
```

## See also

 * <span class="badge object-type-class"></span> [Expression](./object-Expression.md)
