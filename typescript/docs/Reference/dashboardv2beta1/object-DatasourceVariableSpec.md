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
	refresh: dashboardv2beta1.VariableRefresh;
	regex: string;
	current: dashboardv2beta1.VariableOption;
	options: dashboardv2beta1.VariableOption[];
	multi: boolean;
	includeAll: boolean;
	allValue?: string;
	label?: string;
	hide: dashboardv2beta1.VariableHide;
	skipUrlSync: boolean;
	description?: string;
	allowCustomValue: boolean;
}

```
