---
title: <span class="badge object-type-class"></span> CanvasElementOptions
---
# <span class="badge object-type-class"></span> CanvasElementOptions

## Definition

```php
class CanvasElementOptions implements \JsonSerializable
{
    public string $name;

    public string $type;

    /**
     * TODO: figure out how to define this (element config(s))
     * @var mixed|null
     */
    public $config;

    public ?\Grafana\Foundation\Canvas\Constraint $constraint;

    public ?\Grafana\Foundation\Canvas\Placement $placement;

    public ?\Grafana\Foundation\Canvas\BackgroundConfig $background;

    public ?\Grafana\Foundation\Canvas\LineConfig $border;

    /**
     * @var array<\Grafana\Foundation\Canvas\CanvasConnection>|null
     */
    public ?array $connections;

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

 * <span class="badge builder"></span> [CanvasElementOptionsBuilder](./builder-CanvasElementOptionsBuilder.md)
