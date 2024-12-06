---
title: <span class="badge object-type-class"></span> DashboardLink
---
# <span class="badge object-type-class"></span> DashboardLink

Links with references to other dashboards or external resources

## Definition

```php
class DashboardLink implements \JsonSerializable
{
    /**
     * Title to display with the link
     */
    public string $title;

    /**
     * Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
     */
    public \Grafana\Foundation\Dashboard\DashboardLinkType $type;

    /**
     * Icon name to be displayed with the link
     */
    public string $icon;

    /**
     * Tooltip to display when the user hovers their mouse over it
     */
    public string $tooltip;

    /**
     * Link URL. Only required/valid if the type is link
     */
    public ?string $url;

    /**
     * List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
     * @var array<string>
     */
    public array $tags;

    /**
     * If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
     */
    public bool $asDropdown;

    /**
     * If true, the link will be opened in a new tab
     */
    public bool $targetBlank;

    /**
     * If true, includes current template variables values in the link as query params
     */
    public bool $includeVars;

    /**
     * If true, includes current time range in the link as query params
     */
    public bool $keepTime;

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

 * <span class="badge builder"></span> [DashboardLinkBuilder](./builder-DashboardLinkBuilder.md)
