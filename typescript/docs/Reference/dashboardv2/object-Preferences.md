---
title: <span class="badge object-type-interface"></span> Preferences
---
# <span class="badge object-type-interface"></span> Preferences

Dashboard specific preferences (applied per dashboard = all users using the dashboard)

## Definition

```typescript
export interface Preferences {
	// default layout template to be used when new containers are created
	layout?: dashboardv2.AutoGridLayoutKind | dashboardv2.GridLayoutKind;
}

```
## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
