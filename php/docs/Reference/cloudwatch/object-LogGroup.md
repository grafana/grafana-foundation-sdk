---
title: <span class="badge object-type-class"></span> LogGroup
---
# <span class="badge object-type-class"></span> LogGroup

## Definition

```php
class LogGroup implements \JsonSerializable
{
    /**
     * ARN of the log group
     */
    public string $arn;

    /**
     * Name of the log group
     */
    public string $name;

    /**
     * AccountId of the log group
     */
    public ?string $accountId;

    /**
     * Label of the log group
     */
    public ?string $accountLabel;

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

 * <span class="badge builder"></span> [LogGroupBuilder](./builder-LogGroupBuilder.md)
