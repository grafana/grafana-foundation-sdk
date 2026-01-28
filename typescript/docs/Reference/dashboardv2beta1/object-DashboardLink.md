---
title: <span class="badge object-type-interface"></span> DashboardLink
---
# <span class="badge object-type-interface"></span> DashboardLink

Links with references to other dashboards or external resources

## Definition

```typescript
export interface DashboardLink {
	// Title to display with the link
	title: string;
	// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
	// FIXME: The type is generated as `type: DashboardLinkType | dashboardLinkType.Link;` but it should be `type: DashboardLinkType`
	type: dashboardv2beta1.DashboardLinkType;
	// Icon name to be displayed with the link
	icon: string;
	// Tooltip to display when the user hovers their mouse over it
	tooltip: string;
	// Link URL. Only required/valid if the type is link
	url?: string;
	// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
	tags: string[];
	// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
	asDropdown: boolean;
	// If true, the link will be opened in a new tab
	targetBlank: boolean;
	// If true, includes current template variables values in the link as query params
	includeVars: boolean;
	// If true, includes current time range in the link as query params
	keepTime: boolean;
	// Placement can be used to display the link somewhere else on the dashboard other than above the visualisations.
	placement?: "inControlsMenu";
}

```
## See also

 * <span class="badge builder"></span> [DashboardLinkBuilder](./builder-DashboardLinkBuilder.md)
