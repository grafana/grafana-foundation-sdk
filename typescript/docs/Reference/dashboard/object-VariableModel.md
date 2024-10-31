---
title: <span class="badge object-type-interface"></span> VariableModel
---
# <span class="badge object-type-interface"></span> VariableModel

A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.

## Definition

```typescript
export interface VariableModel {
	// Type of variable
	type: dashboard.VariableType;
	// Name of variable
	name: string;
	// Optional display name
	label?: string;
	// Visibility configuration for the variable
	hide?: dashboard.VariableHide;
	// Whether the variable value should be managed by URL query params or not
	skipUrlSync?: boolean;
	// Description of variable. It can be defined but `null`.
	description?: string;
	// Query used to fetch values for a variable
	query?: string | Record<string, any>;
	// Data source used to fetch values for a variable. It can be defined but `null`.
	datasource?: dashboard.DataSourceRef;
	// Shows current selected variable text/value on the dashboard
	current?: dashboard.VariableOption;
	// Whether multiple values can be selected or not from variable value list
	multi?: boolean;
	// Options that can be selected for a variable.
	options?: dashboard.VariableOption[];
	// Options to config when to refresh a variable
	refresh?: dashboard.VariableRefresh;
	// Options sort order
	sort?: dashboard.VariableSort;
	// Whether all value option is available or not
	includeAll?: boolean;
	// Custom all value
	allValue?: string;
	// Optional field, if you want to extract part of a series name or metric node segment.
	// Named capture groups can be used to separate the display text and value.
	regex?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AdHocVariableBuilder](./builder-AdHocVariableBuilder.md)
 * <span class="badge builder"></span> [ConstantVariableBuilder](./builder-ConstantVariableBuilder.md)
 * <span class="badge builder"></span> [CustomVariableBuilder](./builder-CustomVariableBuilder.md)
 * <span class="badge builder"></span> [DatasourceVariableBuilder](./builder-DatasourceVariableBuilder.md)
 * <span class="badge builder"></span> [IntervalVariableBuilder](./builder-IntervalVariableBuilder.md)
 * <span class="badge builder"></span> [QueryVariableBuilder](./builder-QueryVariableBuilder.md)
 * <span class="badge builder"></span> [TextBoxVariableBuilder](./builder-TextBoxVariableBuilder.md)
