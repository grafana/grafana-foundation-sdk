---
title: <span class="badge object-type-class"></span> Playlist
---
# <span class="badge object-type-class"></span> Playlist

## Definition

```php
class Playlist implements \JsonSerializable
{
    /**
     * Unique playlist identifier. Generated on creation, either by the
     * creator of the playlist of by the application.
     */
    public string $uid;

    /**
     * Name of the playlist.
     */
    public string $name;

    /**
     * Interval sets the time between switching views in a playlist.
     * FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
     */
    public string $interval;

    /**
     * The ordered list of items that the playlist will iterate over.
     * FIXME! This should not be optional, but changing it makes the godegen awkward
     * @var array<\Grafana\Foundation\Playlist\PlaylistItem>|null
     */
    public ?array $items;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [PlaylistBuilder](./builder-PlaylistBuilder.md)
