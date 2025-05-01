---
title: <span class="badge builder"></span> DashboardLinkBuilder
---
# <span class="badge builder"></span> DashboardLinkBuilder

## Constructor

```java
new DashboardLinkBuilder(String title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public DashboardLink build()
```

### <span class="badge object-method"></span> asDropdown

If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards

```java
public DashboardLinkBuilder asDropdown(Boolean asDropdown)
```

### <span class="badge object-method"></span> icon

Icon name to be displayed with the link

```java
public DashboardLinkBuilder icon(String icon)
```

### <span class="badge object-method"></span> includeVars

If true, includes current template variables values in the link as query params

```java
public DashboardLinkBuilder includeVars(Boolean includeVars)
```

### <span class="badge object-method"></span> keepTime

If true, includes current time range in the link as query params

```java
public DashboardLinkBuilder keepTime(Boolean keepTime)
```

### <span class="badge object-method"></span> tags

List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards

```java
public DashboardLinkBuilder tags(List<String> tags)
```

### <span class="badge object-method"></span> targetBlank

If true, the link will be opened in a new tab

```java
public DashboardLinkBuilder targetBlank(Boolean targetBlank)
```

### <span class="badge object-method"></span> title

Title to display with the link

```java
public DashboardLinkBuilder title(String title)
```

### <span class="badge object-method"></span> tooltip

Tooltip to display when the user hovers their mouse over it

```java
public DashboardLinkBuilder tooltip(String tooltip)
```

### <span class="badge object-method"></span> type

Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

```java
public DashboardLinkBuilder type(DashboardLinkType type)
```

### <span class="badge object-method"></span> url

Link URL. Only required/valid if the type is link

```java
public DashboardLinkBuilder url(String url)
```

## See also

 * <span class="badge object-type-class"></span> [DashboardLink](./object-DashboardLink.md)
