---
title: <span class="badge object-type-interface"></span> LegacyCloudMonitoringAnnotationQuery
---
# <span class="badge object-type-interface"></span> LegacyCloudMonitoringAnnotationQuery

@deprecated Use AnnotationQuery instead. Legacy annotation query properties for migration purposes.

## Definition

```typescript
export interface LegacyCloudMonitoringAnnotationQuery {
	// GCP project to execute the query against.
	projectName: string;
	metricType: string;
	// Query refId.
	refId: string;
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	filters: string[];
	metricKind: googlecloudmonitoring.MetricKind;
	valueType: string;
	// Annotation title.
	title: string;
	// Annotation text.
	text: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [LegacyCloudMonitoringAnnotationQueryBuilder](./builder-LegacyCloudMonitoringAnnotationQueryBuilder.md)
