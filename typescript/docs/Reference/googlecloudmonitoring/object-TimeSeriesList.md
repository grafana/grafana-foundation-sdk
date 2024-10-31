---
title: <span class="badge object-type-interface"></span> TimeSeriesList
---
# <span class="badge object-type-interface"></span> TimeSeriesList

Time Series List sub-query properties.

## Definition

```typescript
export interface TimeSeriesList {
	// GCP project to execute the query against.
	projectName: string;
	// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	crossSeriesReducer: string;
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	alignmentPeriod?: string;
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	perSeriesAligner?: string;
	// Array of labels to group data by.
	groupBys?: string[];
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	filters?: string[];
	// Data view, defaults to FULL.
	view?: string;
	// Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	secondaryCrossSeriesReducer?: string;
	// Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	secondaryAlignmentPeriod?: string;
	// Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
	secondaryPerSeriesAligner?: string;
	// Only present if a preprocessor is selected. Array of labels to group data by.
	secondaryGroupBys?: string[];
	// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
	preprocessor?: googlecloudmonitoring.PreprocessorType;
	// Annotation title.
	title?: string;
	// Annotation text.
	text?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TimeSeriesListBuilder](./builder-TimeSeriesListBuilder.md)
