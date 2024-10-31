---
title: <span class="badge object-type-interface"></span> SLOQuery
---
# <span class="badge object-type-interface"></span> SLOQuery

SLO sub-query properties.

## Definition

```typescript
export interface SLOQuery {
	// GCP project to execute the query against.
	projectName: string;
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	perSeriesAligner?: string;
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	alignmentPeriod?: string;
	// SLO selector.
	selectorName: string;
	// ID for the service the SLO is in.
	serviceId: string;
	// Name for the service the SLO is in.
	serviceName: string;
	// ID for the SLO.
	sloId: string;
	// Name of the SLO.
	sloName: string;
	// SLO goal value.
	goal?: number;
	// Specific lookback period for the SLO.
	lookbackPeriod?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [SLOQueryBuilder](./builder-SLOQueryBuilder.md)
