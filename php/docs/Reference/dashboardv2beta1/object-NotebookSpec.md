---
title: <span class="badge object-type-class"></span> NotebookSpec
---
# <span class="badge object-type-class"></span> NotebookSpec

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Definition

```php
class NotebookSpec implements \JsonSerializable
{
    public string $title;

    public ?string $description;

    /**
     * @var array<string>
     */
    public array $tags;

    public \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $timeSettings;

    /**
     * @var array<string, \Grafana\Foundation\Dashboardv2beta1\CellKind|\Grafana\Foundation\Dashboardv2beta1\V2PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>
     */
    public array $elements;

    public \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind $layout;

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

 * <span class="badge builder"></span> [NotebookSpecBuilder](./builder-NotebookSpecBuilder.md)
