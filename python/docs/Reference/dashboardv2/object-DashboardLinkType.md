---
title: <span class="badge object-type-enum"></span> DashboardLinkType
---
# <span class="badge object-type-enum"></span> DashboardLinkType

Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

## Definition

```python
class DashboardLinkType(enum.StrEnum):
    """
    Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    """

    LINK = "link"
    DASHBOARDS = "dashboards"
```
