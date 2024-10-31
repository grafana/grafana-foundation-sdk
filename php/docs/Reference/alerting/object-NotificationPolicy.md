---
title: <span class="badge object-type-class"></span> NotificationPolicy
---
# <span class="badge object-type-class"></span> NotificationPolicy

A Route is a node that contains definitions of how to handle alerts. This is modified

from the upstream alertmanager in that it adds the ObjectMatchers property.

## Definition

```php
class NotificationPolicy implements \JsonSerializable
{
    public ?bool $continue;

    /**
     * @var array<string>|null
     */
    public ?array $groupBy;

    public ?string $groupInterval;

    public ?string $groupWait;

    /**
     * Deprecated. Remove before v1.0 release.
     * @var array<string, string>|null
     */
    public ?array $match;

    /**
     * @var array<string, string>
     */
    public array $matchRe;

    /**
     * Matchers is a slice of Matchers that is sortable, implements Stringer, and
     * provides a Matches method to match a LabelSet against all Matchers in the
     * slice. Note that some users of Matchers might require it to be sorted.
     * @var array<\Grafana\Foundation\Alerting\Matcher>
     */
    public array $matchers;

    /**
     * @var array<string>|null
     */
    public ?array $muteTimeIntervals;

    /**
     * Matchers is a slice of Matchers that is sortable, implements Stringer, and
     * provides a Matches method to match a LabelSet against all Matchers in the
     * slice. Note that some users of Matchers might require it to be sorted.
     * @var array<\Grafana\Foundation\Alerting\Matcher>
     */
    public array $objectMatchers;

    public string $provenance;

    public ?string $receiver;

    public ?string $repeatInterval;

    /**
     * @var array<\Grafana\Foundation\Alerting\NotificationPolicy>|null
     */
    public ?array $routes;

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

 * <span class="badge builder"></span> [NotificationPolicyBuilder](./builder-NotificationPolicyBuilder.md)
