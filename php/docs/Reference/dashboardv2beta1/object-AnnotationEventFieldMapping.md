---
title: <span class="badge object-type-class"></span> AnnotationEventFieldMapping
---
# <span class="badge object-type-class"></span> AnnotationEventFieldMapping

Annotation event field mapping. Defines how to map a data frame field to an annotation event field.

## Definition

```php
class AnnotationEventFieldMapping implements \JsonSerializable
{
    /**
     * Source type for the field value
     */
    public ?string $source;

    /**
     * Constant value to use when source is "text"
     */
    public ?string $value;

    /**
     * Regular expression to apply to the field value
     */
    public ?string $regex;

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

 * <span class="badge builder"></span> [AnnotationEventFieldMappingBuilder](./builder-AnnotationEventFieldMappingBuilder.md)
