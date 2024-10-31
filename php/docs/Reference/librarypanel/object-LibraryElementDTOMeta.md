---
title: <span class="badge object-type-class"></span> LibraryElementDTOMeta
---
# <span class="badge object-type-class"></span> LibraryElementDTOMeta

## Definition

```php
class LibraryElementDTOMeta implements \JsonSerializable
{
    public string $folderName;

    public string $folderUid;

    public int $connectedDashboards;

    public string $created;

    public string $updated;

    public \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $createdBy;

    public \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $updatedBy;

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

 * <span class="badge builder"></span> [LibraryElementDTOMetaBuilder](./builder-LibraryElementDTOMetaBuilder.md)
