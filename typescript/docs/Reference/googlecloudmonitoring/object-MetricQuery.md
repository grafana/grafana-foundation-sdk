---
title: <span class="badge object-type-interface"></span> MetricQuery
---
# <span class="badge object-type-interface"></span> MetricQuery

@deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.

## Definition

```typescript
export interface MetricQuery {
	// GCP project to execute the query against.
	projectName: string;
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	perSeriesAligner?: string;
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	alignmentPeriod?: string;
	// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
	aliasBy?: string;
	editorMode: string;
	metricType: string;
	// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	crossSeriesReducer: string;
	// Array of labels to group data by.
	groupBys?: string[];
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	filters?: string[];
	metricKind?: googlecloudmonitoring.MetricKind;
	valueType?: string;
	view?: string;
	// MQL query to be executed.
	query: string;
	// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
	preprocessor?: googlecloudmonitoring.PreprocessorType;
	// To disable the graphPeriod, it should explictly be set to 'disabled'.
	graphPeriod?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
