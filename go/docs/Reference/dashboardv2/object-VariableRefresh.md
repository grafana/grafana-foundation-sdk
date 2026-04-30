---
title: <span class="badge object-type-enum"></span> VariableRefresh
---
# <span class="badge object-type-enum"></span> VariableRefresh

Options to config when to refresh a variable

`never`: Never refresh the variable

`onDashboardLoad`: Queries the data source every time the dashboard loads.

`onTimeRangeChanged`: Queries the data source when the dashboard time range changes.

## Definition

```go
type VariableRefresh string
const (
	VariableRefreshNever VariableRefresh = "never"
	VariableRefreshOnDashboardLoad VariableRefresh = "onDashboardLoad"
	VariableRefreshOnTimeRangeChanged VariableRefresh = "onTimeRangeChanged"
)

```
