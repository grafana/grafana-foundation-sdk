---
title: <span class="badge object-type-class"></span> NotebookLayoutItemSpec
---
# <span class="badge object-type-class"></span> NotebookLayoutItemSpec

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Definition

```php
class NotebookLayoutItemSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ElementReference $element;

    public \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource $source;

    public ?bool $collapsed;

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

 * <span class="badge builder"></span> [NotebookLayoutItemSpecBuilder](./builder-NotebookLayoutItemSpecBuilder.md)
