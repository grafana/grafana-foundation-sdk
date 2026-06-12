---
title: <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Dashboard
---
# <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Dashboard

<span class="badge deprecated"></span>Prefer using dashboardv2.Dashboard instead.

## Definition

```java
public class Dashboard {
  public List<AnnotationQueryKind> annotations;
  public DashboardCursorSync cursorSync;
  public String description;
  public Boolean editable;
  public Map<String, Element> elements;
  public GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind layout;
  public List<DashboardLink> links;
  public Boolean liveNow;
  public Boolean preload;
  public Short revision;
  public List<String> tags;
  public TimeSettingsSpec timeSettings;
  public String title;
  public List<VariableKind> variables;
}
```
## See also

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [DashboardBuilder](./builder-DashboardBuilder.md)
