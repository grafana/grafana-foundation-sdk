---
title: <span class="badge object-type-class"></span> LibraryPanelKindSpec
---
# <span class="badge object-type-class"></span> LibraryPanelKindSpec

## Definition

```php
class LibraryPanelKindSpec implements \JsonSerializable
{
    /**
     * Panel ID for the library panel in the dashboard
     */
    public float $id;

    /**
     * Title for the library panel in the dashboard
     */
    public string $title;

    public \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef $libraryPanel;

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

 * <span class="badge builder"></span> [LibraryPanelKindSpecBuilder](./builder-LibraryPanelKindSpecBuilder.md)
