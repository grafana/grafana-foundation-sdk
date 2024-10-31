---
title: <span class="badge object-type-interface"></span> DataTransformerConfig
---
# <span class="badge object-type-interface"></span> DataTransformerConfig

Transformations allow to manipulate data returned by a query before the system applies a visualization.

Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,

use the output of one transformation as the input to another transformation, etc.

## Definition

```typescript
export interface DataTransformerConfig {
	// Unique identifier of transformer
	id: string;
	// Disabled transformations are skipped
	disabled?: boolean;
	// Optional frame matcher. When missing it will be applied to all results
	filter?: dashboard.MatcherConfig;
	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	options: any;
}

```
## Methods

No methods.
