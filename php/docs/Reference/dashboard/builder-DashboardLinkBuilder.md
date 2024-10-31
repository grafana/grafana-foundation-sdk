---
title: <span class="badge builder"></span> DashboardLinkBuilder
---
# <span class="badge builder"></span> DashboardLinkBuilder

## Constructor

```php
new DashboardLinkBuilder(string $title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> asDropdown

If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards

```php
asDropdown(bool $asDropdown)
```

### <span class="badge object-method"></span> icon

Icon name to be displayed with the link

```php
icon(string $icon)
```

### <span class="badge object-method"></span> includeVars

If true, includes current template variables values in the link as query params

```php
includeVars(bool $includeVars)
```

### <span class="badge object-method"></span> keepTime

If true, includes current time range in the link as query params

```php
keepTime(bool $keepTime)
```

### <span class="badge object-method"></span> tags

List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards

@param array<string> $tags

```php
tags(array $tags)
```

### <span class="badge object-method"></span> targetBlank

If true, the link will be opened in a new tab

```php
targetBlank(bool $targetBlank)
```

### <span class="badge object-method"></span> title

Title to display with the link

```php
title(string $title)
```

### <span class="badge object-method"></span> tooltip

Tooltip to display when the user hovers their mouse over it

```php
tooltip(string $tooltip)
```

### <span class="badge object-method"></span> type

Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

```php
type(\Grafana\Foundation\Dashboard\DashboardLinkType $type)
```

### <span class="badge object-method"></span> url

Link URL. Only required/valid if the type is link

```php
url(string $url)
```

## See also

 * <span class="badge object-type-class"></span> [DashboardLink](./object-DashboardLink.md)
