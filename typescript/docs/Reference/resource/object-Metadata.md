---
title: <span class="badge object-type-interface"></span> Metadata
---
# <span class="badge object-type-interface"></span> Metadata

## Definition

```typescript
export interface Metadata {
	name: string;
	namespace?: string;
	labels?: Record<string, string>;
	annotations?: Record<string, string>;
	uid?: string;
	resourceVersion?: string;
	generation?: number;
	creationTimestamp?: string;
	updateTimestamp?: string;
	deletionTimestamp?: string;
}

```
## See also

 * <span class="badge builder"></span> [MetadataBuilder](./builder-MetadataBuilder.md)
