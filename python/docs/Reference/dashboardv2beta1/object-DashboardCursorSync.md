---
title: <span class="badge object-type-enum"></span> DashboardCursorSync
---
# <span class="badge object-type-enum"></span> DashboardCursorSync

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

## Definition

```python
class DashboardCursorSync(enum.StrEnum):
    """
    "Off" for no shared crosshair or tooltip (default).
    "Crosshair" for shared crosshair.
    "Tooltip" for shared crosshair AND shared tooltip.
    """

    CROSSHAIR = "Crosshair"
    TOOLTIP = "Tooltip"
    OFF = "Off"
```
