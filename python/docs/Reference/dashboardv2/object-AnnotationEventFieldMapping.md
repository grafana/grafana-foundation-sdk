---
title: <span class="badge object-type-class"></span> AnnotationEventFieldMapping
---
# <span class="badge object-type-class"></span> AnnotationEventFieldMapping

Annotation event field mapping. Defines how to map a data frame field to an annotation event field.

## Definition

```python
class AnnotationEventFieldMapping:
    """
    Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
    """

    # Source type for the field value
    source: typing.Optional[str]
    # Constant value to use when source is "text"
    value: typing.Optional[str]
    # Regular expression to apply to the field value
    regex: typing.Optional[str]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [AnnotationEventFieldMapping](./builder-AnnotationEventFieldMapping.md)
