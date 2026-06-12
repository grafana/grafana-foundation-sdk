---
title: <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Preferences
---
# <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Preferences

<span class="badge deprecated"></span>Prefer using preferencesv1alpha1.Preferences instead.

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
  public QueryHistoryPreference queryHistory;
  public CookiePreferences cookiePreferences;
  public NavbarPreference navbar;
}
```
## See also

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
