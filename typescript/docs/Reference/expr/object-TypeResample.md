---
title: <span class="badge object-type-interface"></span> TypeResample
---
# <span class="badge object-type-interface"></span> TypeResample

## Definition

```typescript
export interface TypeResample {
	// The datasource
	datasource?: dashboard.DataSourceRef;
	// The downsample function
	// Possible enum values:
	//  - `"sum"` 
	//  - `"mean"` 
	//  - `"min"` 
	//  - `"max"` 
	//  - `"count"` 
	//  - `"last"` 
	//  - `"median"` 
	downsampler: "sum" | "mean" | "min" | "max" | "count" | "last" | "median";
	// The math expression
	expression: string;
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	intervalMs?: number;
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	maxDataPoints?: number;
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	queryType?: string;
	// RefID is the unique identifier of the query, set by the frontend call.
	refId: string;
	// Optionally define expected query result behavior
	resultAssertions?: {
		// Maximum frame count
		maxFrames?: number;
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""` 
		//  - `"timeseries-wide"` 
		//  - `"timeseries-long"` 
		//  - `"timeseries-many"` 
		//  - `"timeseries-multi"` 
		//  - `"directory-listing"` 
		//  - `"table"` 
		//  - `"numeric-wide"` 
		//  - `"numeric-multi"` 
		//  - `"numeric-long"` 
		//  - `"log-lines"` 
		type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		typeVersion: number[];
	};
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	timeRange?: {
		// From is the start time of the query.
		from: string;
		// To is the end time of the query.
		to: string;
	};
	type: "resample";
	// The upsample function
	// Possible enum values:
	//  - `"pad"` Use the last seen value
	//  - `"backfilling"` backfill
	//  - `"fillna"` Do not fill values (nill)
	upsampler: "pad" | "backfilling" | "fillna";
	// The time duration
	window: string;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [TypeResampleBuilder](./builder-TypeResampleBuilder.md)
