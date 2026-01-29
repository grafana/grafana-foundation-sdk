---
title: <span class="badge object-type-class"></span> TextVariableSpec
---
# <span class="badge object-type-class"></span> TextVariableSpec

Text variable specification

## Definition

```php
class TextVariableSpec implements \JsonSerializable
{
    public string $name;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    public string $query;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

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

