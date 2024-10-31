---
title: <span class="badge object-type-class"></span> DashboardLink
---
# <span class="badge object-type-class"></span> DashboardLink

Links with references to other dashboards or external resources

## Definition

```python
class DashboardLink:
    """
    Links with references to other dashboards or external resources
    """

    # Title to display with the link
    title: str
    # Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    type_val: dashboard.DashboardLinkType
    # Icon name to be displayed with the link
    icon: str
    # Tooltip to display when the user hovers their mouse over it
    tooltip: str
    # Link URL. Only required/valid if the type is link
    url: str
    # List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
    tags: list[str]
    # If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
    as_dropdown: bool
    # If true, the link will be opened in a new tab
    target_blank: bool
    # If true, includes current template variables values in the link as query params
    include_vars: bool
    # If true, includes current time range in the link as query params
    keep_time: bool
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

 * <span class="badge builder"></span> [DashboardLink](./builder-DashboardLink.md)
