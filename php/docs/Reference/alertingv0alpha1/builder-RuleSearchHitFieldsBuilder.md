---
title: <span class="badge builder"></span> RuleSearchHitFieldsBuilder
---
# <span class="badge builder"></span> RuleSearchHitFieldsBuilder

#RuleSearchHitFields is the per-kind field payload returned on each hit.

It carries the union of alert- and recording-rule search fields; only the

fields relevant to a hit's kind are populated. This maps to the kind's

declared searchFields.

## Constructor

```php
new RuleSearchHitFieldsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> annotations

Alert-rule fields.

@param array<string, string> $annotations

```php
annotations(array $annotations)
```

### <span class="badge object-method"></span> dashboardUID

```php
dashboardUID(string $dashboardUID)
```

### <span class="badge object-method"></span> datasourceUIDs

@param array<string> $datasourceUIDs

```php
datasourceUIDs(array $datasourceUIDs)
```

### <span class="badge object-method"></span> folder

```php
folder(string $folder)
```

### <span class="badge object-method"></span> for

```php
for(string $for)
```

### <span class="badge object-method"></span> interval

```php
interval(string $interval)
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

### <span class="badge object-method"></span> metric

Recording-rule fields.

```php
metric(string $metric)
```

### <span class="badge object-method"></span> notificationType

```php
notificationType(string $notificationType)
```

### <span class="badge object-method"></span> panelID

```php
panelID(int $panelID)
```

### <span class="badge object-method"></span> paused

```php
paused(bool $paused)
```

### <span class="badge object-method"></span> receiver

```php
receiver(string $receiver)
```

### <span class="badge object-method"></span> routingTree

```php
routingTree(string $routingTree)
```

### <span class="badge object-method"></span> targetDatasourceUID

```php
targetDatasourceUID(string $targetDatasourceUID)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

### <span class="badge object-method"></span> type

```php
type(string $type)
```

## See also

 * <span class="badge object-type-class"></span> [RuleSearchHitFields](./object-RuleSearchHitFields.md)
