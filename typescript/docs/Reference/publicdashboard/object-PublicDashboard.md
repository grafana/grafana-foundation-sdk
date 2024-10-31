---
title: <span class="badge object-type-interface"></span> PublicDashboard
---
# <span class="badge object-type-interface"></span> PublicDashboard

## Definition

```typescript
export interface PublicDashboard {
	// Unique public dashboard identifier
	uid: string;
	// Dashboard unique identifier referenced by this public dashboard
	dashboardUid: string;
	// Unique public access token
	accessToken?: string;
	// Flag that indicates if the public dashboard is enabled
	isEnabled: boolean;
	// Flag that indicates if annotations are enabled
	annotationsEnabled: boolean;
	// Flag that indicates if the time range picker is enabled
	timeSelectionEnabled: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [PublicDashboardBuilder](./builder-PublicDashboardBuilder.md)
