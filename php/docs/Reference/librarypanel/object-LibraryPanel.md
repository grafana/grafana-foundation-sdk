---
title: <span class="badge object-type-class"></span> LibraryPanel
---
# <span class="badge object-type-class"></span> LibraryPanel

## Definition

```php
class LibraryPanel implements \JsonSerializable
{
    /**
     * Folder UID
     */
    public ?string $folderUid;

    /**
     * Library element UID
     */
    public string $uid;

    /**
     * Panel name (also saved in the model)
     */
    public string $name;

    /**
     * Panel description
     */
    public ?string $description;

    /**
     * The panel type (from inside the model)
     */
    public string $type;

    /**
     * Dashboard version when this was saved (zero if unknown)
     */
    public ?int $schemaVersion;

    /**
     * panel version, incremented each time the dashboard is updated.
     */
    public int $version;

    /**
     * TODO: should be the same panel schema defined in dashboard
     * Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
     */
    public \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel $model;

    /**
     * Object storage metadata
     */
    public ?\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta $meta;

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

 * <span class="badge builder"></span> [LibraryPanelBuilder](./builder-LibraryPanelBuilder.md)
