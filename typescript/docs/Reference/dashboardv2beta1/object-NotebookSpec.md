---
title: <span class="badge object-type-interface"></span> NotebookSpec
---
# <span class="badge object-type-interface"></span> NotebookSpec

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Definition

```typescript
export interface NotebookSpec {
	title: string;
	description?: string;
	tags: string[];
	timeSettings: dashboardv2beta1.TimeSettingsSpec;
	elements: Record<string, dashboardv2beta1.NotebookElement>;
	layout: dashboardv2beta1.NotebookLayoutKind;
}

```
## See also

 * <span class="badge builder"></span> [NotebookSpecBuilder](./builder-NotebookSpecBuilder.md)
