---
title: <span class="badge builder"></span> SnapshotBuilder
---
# <span class="badge builder"></span> SnapshotBuilder

## Constructor

```php
new SnapshotBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> dashboard

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Dashboard> $dashboard

```php
dashboard(\Grafana\Foundation\Cog\Builder $dashboard)
```

### <span class="badge object-method"></span> expires

Time when the snapshot expires, default is never to expire

```php
expires(string $expires)
```

### <span class="badge object-method"></span> external

Is the snapshot saved in an external grafana instance

```php
external(bool $external)
```

### <span class="badge object-method"></span> externalUrl

external url, if snapshot was shared in external grafana instance

```php
externalUrl(string $externalUrl)
```

### <span class="badge object-method"></span> id

Unique identifier of the snapshot

```php
id(int $id)
```

### <span class="badge object-method"></span> key

Optional, defined the unique key of the snapshot, required if external is true

```php
key(string $key)
```

### <span class="badge object-method"></span> name

Optional, name of the snapshot

```php
name(string $name)
```

### <span class="badge object-method"></span> orgId

org id of the snapshot

```php
orgId(int $orgId)
```

### <span class="badge object-method"></span> originalUrl

original url, url of the dashboard that was snapshotted

```php
originalUrl(string $originalUrl)
```

### <span class="badge object-method"></span> url

url of the snapshot, if snapshot was shared internally

```php
url(string $url)
```

## See also

 * <span class="badge object-type-class"></span> [Snapshot](./object-Snapshot.md)
