---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    public bool $keepTime;

    public bool $includeVars;

    public bool $showStarred;

    public bool $showRecentlyViewed;

    public bool $showSearch;

    public bool $showHeadings;

    public int $maxItems;

    public string $query;

    public ?int $folderId;

    /**
     * @var array<string>
     */
    public array $tags;

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

