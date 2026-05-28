---
title: <span class="badge object-type-interface"></span> DatasourceVariableSpec
---
# <span class="badge object-type-interface"></span> DatasourceVariableSpec

Datasource variable specification

## Definition

```typescript
export interface DatasourceVariableSpec {
	name: string;
	pluginId: string;
	refresh: dashboardv2.VariableRefresh;
	regex: string;
	current: dashboardv2.VariableOption;
	options: dashboardv2.VariableOption[];
	multi: boolean;
	includeAll: boolean;
	allValue?: string;
	label?: string;
	hide: dashboardv2.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	allowCustomValue: boolean;
	origin?: dashboardv2.ControlSourceRef;
}

```
