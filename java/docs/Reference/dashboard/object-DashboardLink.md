---
title: <span class="badge object-type-class"></span> DashboardLink
---
# <span class="badge object-type-class"></span> DashboardLink

Links with references to other dashboards or external resources

## Definition

```java
public class DashboardLink {
  public String title;
  public DashboardLinkType type;
  public String icon;
  public String tooltip;
  public String url;
  public List<String> tags;
  public Boolean asDropdown;
  public DashboardLinkPlacement placement;
  public Boolean targetBlank;
  public Boolean includeVars;
  public Boolean keepTime;
}
```
## See also

 * <span class="badge builder"></span> [DashboardLinkBuilder](./builder-DashboardLinkBuilder.md)
