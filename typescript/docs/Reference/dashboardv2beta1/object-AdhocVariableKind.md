---
title: <span class="badge object-type-interface"></span> AdhocVariableKind
---
# <span class="badge object-type-interface"></span> AdhocVariableKind

Adhoc variable kind

## Definition

```typescript
export interface AdhocVariableKind {
	kind: "AdhocVariable";
	group: string;
	datasource?: {
		name?: string;
	};
	spec: dashboardv2beta1.AdhocVariableSpec;
}

```
## See also

 * <span class="badge builder"></span> [AdhocVariableBuilder](./builder-AdhocVariableBuilder.md)
