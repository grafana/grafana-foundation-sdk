---
title: <span class="badge object-type-interface"></span> AnnotationQuerySpec
---
# <span class="badge object-type-interface"></span> AnnotationQuerySpec

## Definition

```typescript
export interface AnnotationQuerySpec {
	query: dashboardv2.DataQueryKind;
	enable: boolean;
	hide: boolean;
	iconColor: string;
	name: string;
	builtIn?: boolean;
	filter?: dashboardv2.AnnotationPanelFilter;
	// Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
	placement?: "inControlsMenu";
	// Mappings define how to convert data frame fields to annotation event fields.
	mappings?: Record<string, dashboardv2.AnnotationEventFieldMapping>;
	// Catch-all field for datasource-specific properties. Should not be available in as code tooling.
	legacyOptions?: Record<string, any>;
}

```
