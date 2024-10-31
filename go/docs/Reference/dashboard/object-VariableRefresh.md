---
title: <span class="badge object-type-enum"></span> VariableRefresh
---
# <span class="badge object-type-enum"></span> VariableRefresh

Options to config when to refresh a variable

`0`: Never refresh the variable

`1`: Queries the data source every time the dashboard loads.

`2`: Queries the data source when the dashboard time range changes.

## Definition

```go
type VariableRefresh int64
const (
	VariableRefreshNever VariableRefresh = 0
	VariableRefreshOnDashboardLoad VariableRefresh = 1
	VariableRefreshOnTimeRangeChanged VariableRefresh = 2
)

```
