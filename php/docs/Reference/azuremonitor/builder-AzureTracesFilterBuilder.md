---
title: <span class="badge builder"></span> AzureTracesFilterBuilder
---
# <span class="badge builder"></span> AzureTracesFilterBuilder

## Constructor

```php
new AzureTracesFilterBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> filters

Values to filter by.

@param array<string> $filters

```php
filters(array $filters)
```

### <span class="badge object-method"></span> operation

Comparison operator to use. Either equals or not equals.

```php
operation(string $operation)
```

### <span class="badge object-method"></span> property

Property name, auto-populated based on available traces.

```php
property(string $property)
```

## See also

 * <span class="badge object-type-class"></span> [AzureTracesFilter](./object-AzureTracesFilter.md)
