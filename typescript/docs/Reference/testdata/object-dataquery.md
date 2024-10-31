---
title: <span class="badge object-type-interface"></span> dataquery
---
# <span class="badge object-type-interface"></span> dataquery

## Definition

```typescript
export interface dataquery {
	alias?: string;
	scenarioId?: testdata.TestDataQueryType;
	stringInput?: string;
	stream?: testdata.StreamingQuery;
	pulseWave?: testdata.PulseWaveQuery;
	sim?: testdata.SimulationQuery;
	csvWave?: testdata.CSVWave[];
	labels?: string;
	lines?: number;
	levelColumn?: boolean;
	channel?: string;
	nodes?: testdata.NodesQuery;
	csvFileName?: string;
	csvContent?: string;
	rawFrameContent?: string;
	seriesCount?: number;
	usa?: testdata.USAQuery;
	errorType?: "server_panic" | "frontend_exception" | "frontend_observable";
	spanCount?: number;
	points?: (string | number)[][];
	// Drop percentage (the chance we will lose a point 0-100)
	dropPercent?: number;
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

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
