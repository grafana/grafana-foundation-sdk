---
title: <span class="badge object-type-interface"></span> FieldConfigSource
---
# <span class="badge object-type-interface"></span> FieldConfigSource

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```typescript
export interface FieldConfigSource {
	// Defaults are the options applied to all fields.
	defaults: dashboard.FieldConfig;
	// Overrides are the options applied to specific fields overriding the defaults.
	overrides: {
		matcher: dashboard.MatcherConfig;
		properties: dashboard.DynamicConfigValue[];
	}[];
}

```
## Methods

No methods.
