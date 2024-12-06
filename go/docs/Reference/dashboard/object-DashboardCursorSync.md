---
title: <span class="badge object-type-enum"></span> DashboardCursorSync
---
# <span class="badge object-type-enum"></span> DashboardCursorSync

0 for no shared crosshair or tooltip (default).

1 for shared crosshair.

2 for shared crosshair AND shared tooltip.

## Definition

```go
type DashboardCursorSync int64
const (
	DashboardCursorSyncOff DashboardCursorSync = 0
	DashboardCursorSyncCrosshair DashboardCursorSync = 1
	DashboardCursorSyncTooltip DashboardCursorSync = 2
)

```
