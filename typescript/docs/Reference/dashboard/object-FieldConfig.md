---
title: <span class="badge object-type-interface"></span> FieldConfig
---
# <span class="badge object-type-interface"></span> FieldConfig

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```typescript
export interface FieldConfig {
	// The display value for this field.  This supports template variables blank is auto
	displayName?: string;
	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	displayNameFromDS?: string;
	// Human readable field metadata
	description?: string;
	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	// 
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	path?: string;
	// True if data source can write a value to the path. Auth/authz are supported separately
	writeable?: boolean;
	// True if data source field supports ad-hoc filters
	filterable?: boolean;
	// Unit a field should use. The unit you select is applied to all fields except time.
	// You can use the units ID availables in Grafana or a custom unit.
	// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
	// As custom unit, you can use the following formats:
	// `suffix:<suffix>` for custom unit that should go after value.
	// `prefix:<prefix>` for custom unit that should go before value.
	// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
	// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
	// `count:<unit>` for a custom count unit.
	// `currency:<unit>` for custom a currency unit.
	unit?: string;
	// Specify the number of decimals Grafana includes in the rendered value.
	// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
	// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
	// To display all decimals, set the unit to `String`.
	decimals?: number;
	// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	min?: number;
	// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	max?: number;
	// Convert input values into a display string
	mappings?: dashboard.ValueMapping[];
	// Map numeric values to states
	thresholds?: dashboard.ThresholdsConfig;
	// Panel color configuration
	color?: dashboard.FieldColor;
	// The behavior when clicking on a result
	links?: any[];
	// Alternative to empty string
	noValue?: string;
	// custom is specified by the FieldConfig field
	// in panel plugin schemas.
	custom?: any;
}

```
## Methods

No methods.
