---
title: <span class="badge object-type-class"></span> Preferences
---
# <span class="badge object-type-class"></span> Preferences

Spec defines user, team or org Grafana preferences

swagger:model Preferences

## Definition

```java
public class Preferences {
  public String homeDashboardUID;
  public String timezone;
  public String weekStart;
  public String theme;
  public String language;
  public String locale;
  public QueryHistoryPreference queryHistory;
  public CookiePreferences cookiePreferences;
  public NavbarPreference navbar;
}
```
## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
