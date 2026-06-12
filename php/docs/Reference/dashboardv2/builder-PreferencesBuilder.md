---
title: <span class="badge builder"></span> PreferencesBuilder
---
# <span class="badge builder"></span> PreferencesBuilder

Dashboard specific preferences (applied per dashboard = all users using the dashboard)

## Constructor

```php
new PreferencesBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> layout

default layout template to be used when new containers are created

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutKind> $layout

```php
layout($layout)
```

## See also

 * <span class="badge object-type-class"></span> [Preferences](./object-Preferences.md)
