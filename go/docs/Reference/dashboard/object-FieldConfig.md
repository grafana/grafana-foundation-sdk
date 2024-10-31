---
title: <span class="badge object-type-struct"></span> FieldConfig
---
# <span class="badge object-type-struct"></span> FieldConfig

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```go
type FieldConfig struct {
    // The display value for this field.  This supports template variables blank is auto
    DisplayName *string `json:"displayName,omitempty"`
    // This can be used by data sources that return and explicit naming structure for values and labels
    // When this property is configured, this value is used rather than the default naming strategy.
    DisplayNameFromDS *string `json:"displayNameFromDS,omitempty"`
    // Human readable field metadata
    Description *string `json:"description,omitempty"`
    // An explicit path to the field in the datasource.  When the frame meta includes a path,
    // This will default to `${frame.meta.path}/${field.name}
    // 
    // When defined, this value can be used as an identifier within the datasource scope, and
    // may be used to update the results
    Path *string `json:"path,omitempty"`
    // True if data source can write a value to the path. Auth/authz are supported separately
    Writeable *bool `json:"writeable,omitempty"`
    // True if data source field supports ad-hoc filters
    Filterable *bool `json:"filterable,omitempty"`
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
    Unit *string `json:"unit,omitempty"`
    // Specify the number of decimals Grafana includes in the rendered value.
    // If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    // For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    // To display all decimals, set the unit to `String`.
    Decimals *float64 `json:"decimals,omitempty"`
    // The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    Min *float64 `json:"min,omitempty"`
    // The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    Max *float64 `json:"max,omitempty"`
    // Convert input values into a display string
    Mappings []dashboard.ValueMapping `json:"mappings,omitempty"`
    // Map numeric values to states
    Thresholds *dashboard.ThresholdsConfig `json:"thresholds,omitempty"`
    // Panel color configuration
    Color *dashboard.FieldColor `json:"color,omitempty"`
    // The behavior when clicking on a result
    Links []any `json:"links,omitempty"`
    // Alternative to empty string
    NoValue *string `json:"noValue,omitempty"`
    // custom is specified by the FieldConfig field
    // in panel plugin schemas.
    Custom any `json:"custom,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fieldConfig *FieldConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FieldConfig` objects.

```go
func (fieldConfig *FieldConfig) Equals(other FieldConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.

```go
func (fieldConfig *FieldConfig) Validate() error
```

