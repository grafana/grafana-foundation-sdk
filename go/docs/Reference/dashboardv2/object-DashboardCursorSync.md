---
title: <span class="badge object-type-enum"></span> DashboardCursorSync
---
# <span class="badge object-type-enum"></span> DashboardCursorSync

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

## Definition

```go
type DashboardCursorSync string
const (
	DashboardCursorSyncCrosshair DashboardCursorSync = "Crosshair"
	DashboardCursorSyncTooltip DashboardCursorSync = "Tooltip"
	DashboardCursorSyncOff DashboardCursorSync = "Off"
)

```
