---
title: <span class="badge object-type-enum"></span> VariableHide
---
# <span class="badge object-type-enum"></span> VariableHide

Determine if the variable shows on dashboard

Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).

## Definition

```python
class VariableHide(enum.StrEnum):
    """
    Determine if the variable shows on dashboard
    Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).
    """

    DONT_HIDE = "dontHide"
    HIDE_LABEL = "hideLabel"
    HIDE_VARIABLE = "hideVariable"
    IN_CONTROLS_MENU = "inControlsMenu"
```
