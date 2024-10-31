---
title: <span class="badge object-type-class"></span> ContactPoint
---
# <span class="badge object-type-class"></span> ContactPoint

EmbeddedContactPoint is the contact point type that is used

by grafanas embedded alertmanager implementation.

## Definition

```php
class ContactPoint implements \JsonSerializable
{
    public ?bool $disableResolveMessage;

    /**
     * Name is used as grouping key in the UI. Contact points with the
     * same name will be grouped in the UI.
     */
    public ?string $name;

    public ?string $provenance;

    /**
     * @var mixed
     */
    public $settings;

    public \Grafana\Foundation\Alerting\ContactPointType $type;

    /**
     * UID is the unique identifier of the contact point. The UID can be
     * set by the user.
     */
    public ?string $uid;

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

 * <span class="badge builder"></span> [ContactPointBuilder](./builder-ContactPointBuilder.md)
