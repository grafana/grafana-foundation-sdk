---
title: <span class="badge object-type-class"></span> VariableOption
---
# <span class="badge object-type-class"></span> VariableOption

Option to be selected in a variable.

## Definition

```php
class VariableOption implements \JsonSerializable
{
    /**
     * Whether the option is selected or not
     */
    public ?bool $selected;

    /**
     * Text to be displayed for the option
     * @var string|array<string>
     */
    public $text;

    /**
     * Value of the option
     * @var string|array<string>
     */
    public $value;

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

