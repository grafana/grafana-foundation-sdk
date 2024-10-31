---
title: <span class="badge object-type-class"></span> RegexMap
---
# <span class="badge object-type-class"></span> RegexMap

Maps regular expressions to replacement text and a color.

For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.

## Definition

```php
class RegexMap implements \JsonSerializable
{
    public string $type;

    /**
     * Regular expression to match against and the result to apply when the value matches the regex
     */
    public \Grafana\Foundation\Dashboard\DashboardRegexMapOptions $options;

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

 * <span class="badge builder"></span> [RegexMapBuilder](./builder-RegexMapBuilder.md)
