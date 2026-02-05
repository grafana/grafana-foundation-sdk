---
title: <span class="badge object-type-enum"></span> VariableRefresh
---
# <span class="badge object-type-enum"></span> VariableRefresh

Options to config when to refresh a variable

`0`: Never refresh the variable

`1`: Queries the data source every time the dashboard loads.

`2`: Queries the data source when the dashboard time range changes.

## Definition

```python
class VariableRefresh(enum.IntEnum):
    """
    Options to config when to refresh a variable
    `0`: Never refresh the variable
    `1`: Queries the data source every time the dashboard loads.
    `2`: Queries the data source when the dashboard time range changes.
    """

    NEVER = 0
    ON_DASHBOARD_LOAD = 1
    ON_TIME_RANGE_CHANGED = 2
```
