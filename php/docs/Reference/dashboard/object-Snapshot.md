---
title: <span class="badge object-type-class"></span> Snapshot
---
# <span class="badge object-type-class"></span> Snapshot

A dashboard snapshot shares an interactive dashboard publicly.

It is a read-only version of a dashboard, and is not editable.

It is possible to create a snapshot of a snapshot.

Grafana strips away all sensitive information from the dashboard.

Sensitive information stripped: queries (metric, template,annotation) and panel links.

## Definition

```php
class Snapshot implements \JsonSerializable
{
    /**
     * Time when the snapshot was created
     */
    public string $created;

    /**
     * Time when the snapshot expires, default is never to expire
     */
    public string $expires;

    /**
     * Is the snapshot saved in an external grafana instance
     */
    public bool $external;

    /**
     * external url, if snapshot was shared in external grafana instance
     */
    public string $externalUrl;

    /**
     * original url, url of the dashboard that was snapshotted
     */
    public string $originalUrl;

    /**
     * Unique identifier of the snapshot
     */
    public int $id;

    /**
     * Optional, defined the unique key of the snapshot, required if external is true
     */
    public string $key;

    /**
     * Optional, name of the snapshot
     */
    public string $name;

    /**
     * org id of the snapshot
     */
    public int $orgId;

    /**
     * last time when the snapshot was updated
     */
    public string $updated;

    /**
     * url of the snapshot, if snapshot was shared internally
     */
    public ?string $url;

    /**
     * user id of the snapshot creator
     */
    public int $userId;

    public ?\Grafana\Foundation\Dashboard\Dashboard $dashboard;

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

 * <span class="badge builder"></span> [SnapshotBuilder](./builder-SnapshotBuilder.md)
