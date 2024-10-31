---
title: <span class="badge builder"></span> ContactPointBuilder
---
# <span class="badge builder"></span> ContactPointBuilder

## Constructor

```php
new ContactPointBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> disableResolveMessage

```php
disableResolveMessage(bool $disableResolveMessage)
```

### <span class="badge object-method"></span> name

Name is used as grouping key in the UI. Contact points with the

same name will be grouped in the UI.

```php
name(string $name)
```

### <span class="badge object-method"></span> provenance

```php
provenance(string $provenance)
```

### <span class="badge object-method"></span> settings

@param mixed $settings

```php
settings($settings)
```

### <span class="badge object-method"></span> type

```php
type(\Grafana\Foundation\Alerting\ContactPointType $type)
```

### <span class="badge object-method"></span> uid

UID is the unique identifier of the contact point. The UID can be

set by the user.

```php
uid(string $uid)
```

## See also

 * <span class="badge object-type-class"></span> [ContactPoint](./object-ContactPoint.md)
