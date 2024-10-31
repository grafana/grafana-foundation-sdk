---
title: <span class="badge object-type-class"></span> TimePickerConfig
---
# <span class="badge object-type-class"></span> TimePickerConfig

Time picker configuration

It defines the default config for the time picker and the refresh picker for the specific dashboard.

## Definition

```php
class TimePickerConfig implements \JsonSerializable
{
    /**
     * Whether timepicker is visible or not.
     */
    public bool $hidden;

    /**
     * Interval options available in the refresh picker dropdown.
     * @var array<string>
     */
    public array $refreshIntervals;

    /**
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * @var array<string>
     */
    public array $timeOptions;

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

 * <span class="badge builder"></span> [TimePickerBuilder](./builder-TimePickerBuilder.md)
