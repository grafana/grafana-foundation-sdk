---
title: <span class="badge object-type-class"></span> CodeOptions
---
# <span class="badge object-type-class"></span> CodeOptions

## Definition

```php
class CodeOptions implements \JsonSerializable
{
    /**
     * The language passed to monaco code editor
     */
    public \Grafana\Foundation\Text\CodeLanguage $language;

    public bool $showLineNumbers;

    public bool $showMiniMap;

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

 * <span class="badge builder"></span> [CodeOptionsBuilder](./builder-CodeOptionsBuilder.md)
