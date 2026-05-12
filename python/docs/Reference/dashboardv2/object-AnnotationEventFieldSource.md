---
title: <span class="badge object-type-enum"></span> AnnotationEventFieldSource
---
# <span class="badge object-type-enum"></span> AnnotationEventFieldSource

Annotation event field source. Defines how to obtain the value for an annotation event field.

- "field": Find the value with a matching key (default)

- "text": Write a constant string into the value

- "skip": Do not include the field

## Definition

```python
class AnnotationEventFieldSource(enum.StrEnum):
    """
    Annotation event field source. Defines how to obtain the value for an annotation event field.
    - "field": Find the value with a matching key (default)
    - "text": Write a constant string into the value
    - "skip": Do not include the field
    """

    FIELD = "field"
    TEXT = "text"
    SKIP = "skip"
```
