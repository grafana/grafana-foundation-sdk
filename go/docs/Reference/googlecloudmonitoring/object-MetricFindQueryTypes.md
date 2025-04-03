---
title: <span class="badge object-type-enum"></span> MetricFindQueryTypes
---
# <span class="badge object-type-enum"></span> MetricFindQueryTypes

## Definition

```go
type MetricFindQueryTypes string
const (
	MetricFindQueryTypesProjects MetricFindQueryTypes = "projects"
	MetricFindQueryTypesServices MetricFindQueryTypes = "services"
	MetricFindQueryTypesDefaultProject MetricFindQueryTypes = "defaultProject"
	MetricFindQueryTypesMetricTypes MetricFindQueryTypes = "metricTypes"
	MetricFindQueryTypesLabelKeys MetricFindQueryTypes = "labelKeys"
	MetricFindQueryTypesLabelValues MetricFindQueryTypes = "labelValues"
	MetricFindQueryTypesResourceTypes MetricFindQueryTypes = "resourceTypes"
	MetricFindQueryTypesAggregations MetricFindQueryTypes = "aggregations"
	MetricFindQueryTypesAligners MetricFindQueryTypes = "aligners"
	MetricFindQueryTypesAlignmentPeriods MetricFindQueryTypes = "alignmentPeriods"
	MetricFindQueryTypesSelectors MetricFindQueryTypes = "selectors"
	MetricFindQueryTypesSLOServices MetricFindQueryTypes = "sloServices"
	MetricFindQueryTypesSLO MetricFindQueryTypes = "slo"
)

```
