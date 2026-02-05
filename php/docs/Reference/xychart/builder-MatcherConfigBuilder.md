---
title: <span class="badge builder"></span> MatcherConfigBuilder
---
# <span class="badge builder"></span> MatcherConfigBuilder

## Constructor

```php
new MatcherConfigBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> id

The matcher id. This is used to find the matcher implementation from registry.

```php
id(string $id)
```

### <span class="badge object-method"></span> options

The matcher options. This is specific to the matcher implementation.

@param mixed $options

```php
options($options)
```

## See also

 * <span class="badge object-type-class"></span> [MatcherConfig](./object-MatcherConfig.md)
