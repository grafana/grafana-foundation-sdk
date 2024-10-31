---
title: <span class="badge builder"></span> PlaylistBuilder
---
# <span class="badge builder"></span> PlaylistBuilder

## Constructor

```php
new PlaylistBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> interval

Interval sets the time between switching views in a playlist.

FIXME: Is this based on a standardized format or what options are available? Can datemath be used?

```php
interval(string $interval)
```

### <span class="badge object-method"></span> items

The ordered list of items that the playlist will iterate over.

FIXME! This should not be optional, but changing it makes the godegen awkward

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlist\PlaylistItem>> $items

```php
items(array $items)
```

### <span class="badge object-method"></span> name

Name of the playlist.

```php
name(string $name)
```

### <span class="badge object-method"></span> uid

Unique playlist identifier. Generated on creation, either by the

creator of the playlist of by the application.

```php
uid(string $uid)
```

## See also

 * <span class="badge object-type-class"></span> [Playlist](./object-Playlist.md)
