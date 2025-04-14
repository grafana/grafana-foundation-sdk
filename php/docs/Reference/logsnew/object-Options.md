---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    public bool $showControls;

    public bool $showTime;

    public bool $wrapLogMessage;

    public bool $enableLogDetails;

    public bool $syntaxHighlighting;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    /**
     * @var mixed|null
     */
    public $grammar;

    public ?bool $enableInfiniteScrolling;

    /**
     * @var mixed|null
     */
    public $onLogOptionsChange;

    /**
     * @var mixed|null
     */
    public $onNewLogsReceived;

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

