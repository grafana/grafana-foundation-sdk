---
title: <span class="badge object-type-interface"></span> Query
---
# <span class="badge object-type-interface"></span> Query

## Definition

```typescript
export interface Query {
	// Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
	datasourceUid?: string;
	// JSON is the raw JSON query and includes the above properties as well as custom properties.
	model?: cog.Dataquery;
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	queryType?: string;
	// RefID is the unique identifier of the query, set by the frontend call.
	refId?: string;
	// RelativeTimeRange is the per query start and end time
	// for requests.
	relativeTimeRange?: alerting.RelativeTimeRange;
}

```
## See also

 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
