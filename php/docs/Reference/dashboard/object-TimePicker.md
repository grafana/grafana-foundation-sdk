---
title: <span class="badge object-type-class"></span> TimePicker
---
# <span class="badge object-type-class"></span> TimePicker

## Definition

```php
class TimePicker implements \JsonSerializable
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
     * Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
     */
    public bool $collapse;

    /**
     * Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
     */
    public bool $enable;

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
