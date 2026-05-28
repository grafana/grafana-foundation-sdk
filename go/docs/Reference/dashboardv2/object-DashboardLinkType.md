---
title: <span class="badge object-type-enum"></span> DashboardLinkType
---
# <span class="badge object-type-enum"></span> DashboardLinkType

Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

## Definition

```go
type DashboardLinkType string
const (
	DashboardLinkTypeLink DashboardLinkType = "link"
	DashboardLinkTypeDashboards DashboardLinkType = "dashboards"
)

```
