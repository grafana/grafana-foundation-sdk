---
title: <span class="badge object-type-interface"></span> Preferences
---
# <span class="badge object-type-interface"></span> Preferences

## Definition

```typescript
export interface Preferences {
	// Explicit home URL (NOTE: this can only be modified in the system settings)
	homeURL?: string;
	// UID for the home dashboard
	homeDashboardUID?: string;
	// The timezone selection
	timezone?: string;
	// day of the week (sunday, monday, etc)
	weekStart?: string;
	// user interface theme
	theme?: string;
	// Selected language
	language?: string;
	// Explore query history preferences
	queryHistory?: preferencesv1alpha1.QueryHistoryPreference;
	// Navigation preferences
	navbar?: preferencesv1alpha1.NavbarPreference;
}

```
## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
