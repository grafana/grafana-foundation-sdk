---
title: <span class="badge object-type-enum"></span> VariableRegexApplyTo
---
# <span class="badge object-type-enum"></span> VariableRegexApplyTo

Determine whether regex applies to variable value or display text

Accepted values are `value` (apply to value used in queries) or `text` (apply to display text shown to users)

## Definition

```python
class VariableRegexApplyTo(enum.StrEnum):
    """
    Determine whether regex applies to variable value or display text
    Accepted values are `value` (apply to value used in queries) or `text` (apply to display text shown to users)
    """

    VALUE = "value"
    TEXT = "text"
```
