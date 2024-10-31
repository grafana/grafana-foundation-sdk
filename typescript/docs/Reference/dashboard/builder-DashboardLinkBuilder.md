---
title: <span class="badge builder"></span> DashboardLinkBuilder
---
# <span class="badge builder"></span> DashboardLinkBuilder

## Constructor

```typescript
new DashboardLinkBuilder(title: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> asDropdown

If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards

```typescript
asDropdown(asDropdown: boolean)
```

### <span class="badge object-method"></span> icon

Icon name to be displayed with the link

```typescript
icon(icon: string)
```

### <span class="badge object-method"></span> includeVars

If true, includes current template variables values in the link as query params

```typescript
includeVars(includeVars: boolean)
```

### <span class="badge object-method"></span> keepTime

If true, includes current time range in the link as query params

```typescript
keepTime(keepTime: boolean)
```

### <span class="badge object-method"></span> tags

List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards

```typescript
tags(tags: string[])
```

### <span class="badge object-method"></span> targetBlank

If true, the link will be opened in a new tab

```typescript
targetBlank(targetBlank: boolean)
```

### <span class="badge object-method"></span> title

Title to display with the link

```typescript
title(title: string)
```

### <span class="badge object-method"></span> tooltip

Tooltip to display when the user hovers their mouse over it

```typescript
tooltip(tooltip: string)
```

### <span class="badge object-method"></span> type

Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

```typescript
type(type: dashboard.DashboardLinkType)
```

### <span class="badge object-method"></span> url

Link URL. Only required/valid if the type is link

```typescript
url(url: string)
```

## See also

 * <span class="badge object-type-interface"></span> [DashboardLink](./object-DashboardLink.md)
