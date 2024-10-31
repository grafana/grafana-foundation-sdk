---
title: <span class="badge object-type-class"></span> PlaylistItem
---
# <span class="badge object-type-class"></span> PlaylistItem

## Definition

```php
class PlaylistItem implements \JsonSerializable
{
    /**
     * Type of the item.
     */
    public \Grafana\Foundation\Playlist\PlaylistItemType $type;

    /**
     * Value depends on type and describes the playlist item.
     * 
     *  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
     *  is not portable as the numerical identifier is non-deterministic between different instances.
     *  Will be replaced by dashboard_by_uid in the future. (deprecated)
     *  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
     *  dashboards behind the tag will be added to the playlist.
     *  - dashboard_by_uid: The value is the dashboard UID
     */
    public string $value;

    /**
     * Title is an unused property -- it will be removed in the future
     */
    public ?string $title;

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

 * <span class="badge builder"></span> [PlaylistItemBuilder](./builder-PlaylistItemBuilder.md)
