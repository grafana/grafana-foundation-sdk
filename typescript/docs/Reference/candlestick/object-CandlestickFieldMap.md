---
title: <span class="badge object-type-interface"></span> CandlestickFieldMap
---
# <span class="badge object-type-interface"></span> CandlestickFieldMap

## Definition

```typescript
export interface CandlestickFieldMap {
	// Corresponds to the starting value of the given period
	open?: string;
	// Corresponds to the highest value of the given period
	high?: string;
	// Corresponds to the lowest value of the given period
	low?: string;
	// Corresponds to the final (end) value of the given period
	close?: string;
	// Corresponds to the sample count in the given period. (e.g. number of trades)
	volume?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CandlestickFieldMapBuilder](./builder-CandlestickFieldMapBuilder.md)
