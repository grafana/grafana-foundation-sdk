---
title: <span class="badge object-type-enum"></span> SpecialValueMatch
---
# <span class="badge object-type-enum"></span> SpecialValueMatch

Special value types supported by the `SpecialValueMap`

## Definition

```python
class SpecialValueMatch(enum.StrEnum):
    """
    Special value types supported by the `SpecialValueMap`
    """

    TRUE = "true"
    FALSE = "false"
    NULL = "null"
    NA_N = "nan"
    NULL_AND_NAN = "null+nan"
    EMPTY = "empty"
```
