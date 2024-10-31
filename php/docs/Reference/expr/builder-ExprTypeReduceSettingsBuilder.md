---
title: <span class="badge builder"></span> ExprTypeReduceSettingsBuilder
---
# <span class="badge builder"></span> ExprTypeReduceSettingsBuilder

## Constructor

```php
new ExprTypeReduceSettingsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> mode

Non-number reduce behavior

Possible enum values:

 - `"dropNN"` Drop non-numbers

 - `"replaceNN"` Replace non-numbers

```php
mode(\Grafana\Foundation\Expr\TypeReduceMode $mode)
```

### <span class="badge object-method"></span> replaceWithValue

Only valid when mode is replace

```php
replaceWithValue(float $replaceWithValue)
```

## See also

 * <span class="badge object-type-class"></span> [ExprTypeReduceSettings](./object-ExprTypeReduceSettings.md)
