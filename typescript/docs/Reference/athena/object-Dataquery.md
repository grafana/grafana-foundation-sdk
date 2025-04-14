---
title: <span class="badge object-type-interface"></span> Dataquery
---
# <span class="badge object-type-interface"></span> Dataquery

Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts

## Definition

```typescript
export interface Dataquery {
	format: athena.FormatOptions;
	connectionArgs: athena.ConnectionArgs;
	table?: string;
	column?: string;
	queryID?: string;
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
	rawSQL: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	_implementsDataqueryVariant(): void;
}

```
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
