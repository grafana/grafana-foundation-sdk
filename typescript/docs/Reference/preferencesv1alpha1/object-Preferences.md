---
title: <span class="badge object-type-interface"></span> Preferences
---
# <span class="badge object-type-interface"></span> Preferences

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
	// Selected locale (beta)
	regionalFormat?: string;
	// Explore query history preferences
	queryHistory?: preferencesv1alpha1.QueryHistoryPreference;
	// Cookie preferences
	cookiePreferences?: preferencesv1alpha1.CookiePreferences;
	// Navigation preferences
	navbar?: preferencesv1alpha1.NavbarPreference;
}

```
## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
