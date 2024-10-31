---
title: <span class="badge builder"></span> DashboardLink
---
# <span class="badge builder"></span> DashboardLink

## Constructor

```python
DashboardLink(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboard.DashboardLink
```

### <span class="badge object-method"></span> as_dropdown

If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards

```python
def as_dropdown(as_dropdown: bool) -> typing.Self
```

### <span class="badge object-method"></span> icon

Icon name to be displayed with the link

```python
def icon(icon: str) -> typing.Self
```

### <span class="badge object-method"></span> include_vars

If true, includes current template variables values in the link as query params

```python
def include_vars(include_vars: bool) -> typing.Self
```

### <span class="badge object-method"></span> keep_time

If true, includes current time range in the link as query params

```python
def keep_time(keep_time: bool) -> typing.Self
```

### <span class="badge object-method"></span> tags

List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards

```python
def tags(tags: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> target_blank

If true, the link will be opened in a new tab

```python
def target_blank(target_blank: bool) -> typing.Self
```

### <span class="badge object-method"></span> title

Title to display with the link

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> tooltip

Tooltip to display when the user hovers their mouse over it

```python
def tooltip(tooltip: str) -> typing.Self
```

### <span class="badge object-method"></span> type_val

Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

```python
def type_val(type_val: dashboard.DashboardLinkType) -> typing.Self
```

### <span class="badge object-method"></span> url

Link URL. Only required/valid if the type is link

```python
def url(url: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [DashboardLink](./object-DashboardLink.md)
