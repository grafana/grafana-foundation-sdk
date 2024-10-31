---
title: <span class="badge object-type-interface"></span> ResultAssertions
---
# <span class="badge object-type-interface"></span> ResultAssertions

## Definition

```typescript
export interface ResultAssertions {
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
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [ResultAssertionsBuilder](./builder-ResultAssertionsBuilder.md)
