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
	refresh?: dashboard.VariableRefresh;
	// Options sort order
	sort?: dashboard.VariableSort;
	// Dynamically calculates interval by dividing time range by the count specified.
	auto?: boolean;
	// The minimum threshold below which the step count intervals will not divide the time.
	auto_min?: string;
	// How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
	// For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
	auto_count?: number;
}

```
## See also

 * <span class="badge builder"></span> [AdHocVariableBuilder](./builder-AdHocVariableBuilder.md)
 * <span class="badge builder"></span> [ConstantVariableBuilder](./builder-ConstantVariableBuilder.md)
 * <span class="badge builder"></span> [CustomVariableBuilder](./builder-CustomVariableBuilder.md)
 * <span class="badge builder"></span> [DatasourceVariableBuilder](./builder-DatasourceVariableBuilder.md)
 * <span class="badge builder"></span> [IntervalVariableBuilder](./builder-IntervalVariableBuilder.md)
 * <span class="badge builder"></span> [QueryVariableBuilder](./builder-QueryVariableBuilder.md)
 * <span class="badge builder"></span> [TextBoxVariableBuilder](./builder-TextBoxVariableBuilder.md)
