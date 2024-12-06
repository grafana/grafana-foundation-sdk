---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    public bool $showLabels;

    public bool $showCommonLabels;

    public bool $showTime;

    public bool $showLogContextToggle;

    public bool $wrapLogMessage;

    public bool $prettifyLogMessage;

    public bool $enableLogDetails;

    public \Grafana\Foundation\Common\LogsSortOrder $sortOrder;

    public \Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy;

    /**
     * TODO: figure out how to define callbacks
     * @var mixed|null
     */
    public $onClickFilterLabel;

    /**
     * @var mixed|null
     */
    public $onClickFilterOutLabel;

    /**
     * @var mixed|null
     */
    public $isFilterLabelActive;

    /**
     * @var mixed|null
     */
    public $onClickFilterString;

    /**
     * @var mixed|null
     */
    public $onClickFilterOutString;

    /**
     * @var mixed|null
     */
    public $onClickShowField;

    /**
     * @var mixed|null
     */
    public $onClickHideField;

    /**
     * @var array<string>|null
     */
    public ?array $displayedFields;

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

