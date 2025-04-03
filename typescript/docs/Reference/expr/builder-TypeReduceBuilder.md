---
title: <span class="badge builder"></span> TypeReduceBuilder
---
# <span class="badge builder"></span> TypeReduceBuilder

## Constructor

```typescript
new TypeReduceBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasource

The datasource

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> expression

Reference to single query result

```typescript
expression(expression: string)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```typescript
intervalMs(intervalMs: number)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```typescript
maxDataPoints(maxDataPoints: number)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> reducer

The reducer

Possible enum values:

 - `"sum"` 

 - `"mean"` 

 - `"min"` 

 - `"max"` 

 - `"count"` 

 - `"last"` 

 - `"median"` 

```typescript
reducer(reducer: "sum" | "mean" | "min" | "max" | "count" | "last" | "median")
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```typescript
resultAssertions(resultAssertions: {
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
})
```

### <span class="badge object-method"></span> settings

Reducer Options

```typescript
settings(settings: {
	// Non-number reduce behavior
	// Possible enum values:
	//  - `"dropNN"` Drop non-numbers
	//  - `"replaceNN"` Replace non-numbers
	mode: "dropNN" | "replaceNN";
	// Only valid when mode is replace
	replaceWithValue?: number;
})
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```typescript
timeRange(timeRange: {
	// From is the start time of the query.
	from: string;
	// To is the end time of the query.
	to: string;
})
```

## See also

 * <span class="badge object-type-interface"></span> [TypeReduce](./object-TypeReduce.md)
