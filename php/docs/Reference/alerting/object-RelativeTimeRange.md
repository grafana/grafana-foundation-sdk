---
title: <span class="badge object-type-class"></span> RelativeTimeRange
---
# <span class="badge object-type-class"></span> RelativeTimeRange

RelativeTimeRange is the per query start and end time

for requests.

## Definition

```php
class RelativeTimeRange implements \JsonSerializable
{
    /**
     * A Duration represents the elapsed time between two instants
     * as an int64 nanosecond count. The representation limits the
     * largest representable duration to approximately 290 years.
     */
    public int $from;

    /**
     * A Duration represents the elapsed time between two instants
     * as an int64 nanosecond count. The representation limits the
     * largest representable duration to approximately 290 years.
     */
    public int $to;

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

