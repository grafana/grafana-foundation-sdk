---
title: <span class="badge object-type-class"></span> MatcherConfig
---
# <span class="badge object-type-class"></span> MatcherConfig

Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.

It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.

## Definition

```php
class MatcherConfig implements \JsonSerializable
{
    /**
     * The matcher id. This is used to find the matcher implementation from registry.
     */
    public string $id;

    /**
     * The matcher options. This is specific to the matcher implementation.
     * @var mixed|null
     */
    public $options;

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

