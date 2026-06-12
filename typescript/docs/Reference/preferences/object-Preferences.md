---
title: <span class="badge object-type-interface"></span> <span class="badge deprecated"></span> Preferences
---
# <span class="badge object-type-interface"></span> <span class="badge deprecated"></span> Preferences

<span class="badge deprecated"></span>Prefer using preferencesv1alpha1.Preferences instead.

Spec defines user, team or org Grafana preferences

swagger:model Preferences

## Definition

```typescript
export interface Preferences {
	// UID for the home dashboard
	homeDashboardUID?: string;
	// The timezone selection
	// TODO: this should use the timezone defined in common
	timezone?: string;
	// day of the week (sunday, monday, etc)
	weekStart?: string;
	// light, dark, empty is default
	theme?: string;
	// Selected language (beta)
	language?: string;
	// Explore query history preferences
	queryHistory?: preferences.QueryHistoryPreference;
	// Cookie preferences
	cookiePreferences?: preferences.CookiePreferences;
	// Navigation preferences
	navbar?: preferences.NavbarPreference;
}

```
## See also

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
