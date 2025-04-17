---
title: <span class="badge object-type-interface"></span> Manifest
---
# <span class="badge object-type-interface"></span> Manifest

## Definition

```typescript
export interface Manifest {
	apiVersion: string;
	kind: string;
	metadata: resource.Metadata;
	spec: any;
}

```
## See also

 * <span class="badge builder"></span> [ManifestBuilder](./builder-ManifestBuilder.md)
