---
title: <span class="badge object-type-enum"></span> VariableRefresh
---
# <span class="badge object-type-enum"></span> VariableRefresh

Options to config when to refresh a variable

`never`: Never refresh the variable

`onDashboardLoad`: Queries the data source every time the dashboard loads.

`onTimeRangeChanged`: Queries the data source when the dashboard time range changes.

## Definition

```typescript
export enum VariableRefresh {
	Never = "never",
	OnDashboardLoad = "onDashboardLoad",
	OnTimeRangeChanged = "onTimeRangeChanged",
}

```
