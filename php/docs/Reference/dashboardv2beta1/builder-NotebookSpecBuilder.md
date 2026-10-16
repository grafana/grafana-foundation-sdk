---
title: <span class="badge builder"></span> NotebookSpecBuilder
---
# <span class="badge builder"></span> NotebookSpecBuilder

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Constructor

```php
new NotebookSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> description

```php
description(string $description)
```

### <span class="badge object-method"></span> elements

@param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CellKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2PanelKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>> $elements

```php
elements(array $elements)
```

### <span class="badge object-method"></span> layout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind> $layout

```php
layout(\Grafana\Foundation\Cog\Builder $layout)
```

### <span class="badge object-method"></span> tags

@param array<string> $tags

```php
tags(array $tags)
```

### <span class="badge object-method"></span> timeSettings

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec> $timeSettings

```php
timeSettings(\Grafana\Foundation\Cog\Builder $timeSettings)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

## See also

 * <span class="badge object-type-class"></span> [NotebookSpec](./object-NotebookSpec.md)
