---
title: <span class="badge object-type-class"></span> Folder
---
# <span class="badge object-type-class"></span> Folder

TODO:

common metadata will soon support setting the parent folder in the metadata

## Definition

```php
class Folder implements \JsonSerializable
{
    /**
     * Unique folder id. (will be k8s name)
     */
    public string $uid;

    /**
     * Folder title
     */
    public string $title;

    /**
     * Description of the folder.
     */
    public ?string $description;

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

 * <span class="badge builder"></span> [FolderBuilder](./builder-FolderBuilder.md)
