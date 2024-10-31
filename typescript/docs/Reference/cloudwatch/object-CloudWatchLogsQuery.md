---
title: <span class="badge object-type-interface"></span> CloudWatchLogsQuery
---
# <span class="badge object-type-interface"></span> CloudWatchLogsQuery

Shape of a CloudWatch Logs query

## Definition

```typescript
export interface CloudWatchLogsQuery {
	// Whether a query is a Metrics, Logs, or Annotations query
	queryMode: cloudwatch.CloudWatchQueryMode;
	id: string;
	// AWS region to query for the logs
	region: string;
	// The CloudWatch Logs Insights query to execute
	expression?: string;
	// Fields to group the results by, this field is automatically populated whenever the query is updated
	statsGroups?: string[];
	// Log groups to query
	logGroups?: cloudwatch.LogGroup[];
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// @deprecated use logGroups
	logGroupNames?: string[];
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CloudWatchLogsQueryBuilder](./builder-CloudWatchLogsQueryBuilder.md)
