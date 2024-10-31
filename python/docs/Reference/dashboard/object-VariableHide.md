---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).

## Definition

```python
class VariableHide(enum.IntEnum):
    """
    Determine if the variable shows on dashboard
    Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).
    """

    DONT_HIDE = 0
    HIDE_LABEL = 1
    HIDE_VARIABLE = 2
```
