---
title: <span class="badge object-type-enum"></span> DashboardCursorSync
---
# <span class="badge object-type-enum"></span> DashboardCursorSync

0 for no shared crosshair or tooltip (default).

1 for shared crosshair.

2 for shared crosshair AND shared tooltip.

## Definition

```python
class DashboardCursorSync(enum.IntEnum):
    """
    0 for no shared crosshair or tooltip (default).
    1 for shared crosshair.
    2 for shared crosshair AND shared tooltip.
    """

    OFF = 0
    CROSSHAIR = 1
    TOOLTIP = 2
```
