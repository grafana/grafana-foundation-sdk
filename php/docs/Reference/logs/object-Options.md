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

    public ?bool $enableInfiniteScrolling;

    /**
     * Display controls to jump to the last or first log line, and filters by log level.
     */
    public ?bool $showControls;

    /**
     * Show a component to manage the displayed fields from the logs.
     */
    public ?bool $showFieldSelector;

    /**
     * Use a predefined coloring scheme to highlight relevant parts of the log lines.
     */
    public ?bool $syntaxHighlighting;

    public ?\Grafana\Foundation\Logs\OptionsFontSize $fontSize;

    public ?\Grafana\Foundation\Logs\OptionsDetailsMode $detailsMode;

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

