// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Dashboard struct {
	// Unique numeric identifier for the dashboard.
	// `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
	Id *int64 `json:"id,omitempty"`
	// Unique dashboard identifier that can be generated by anyone. string (8-40)
	Uid *string `json:"uid,omitempty"`
	// Title of dashboard.
	Title *string `json:"title,omitempty"`
	// Description of dashboard.
	Description *string `json:"description,omitempty"`
	// This property should only be used in dashboards defined by plugins.  It is a quick check
	// to see if the version has changed since the last time.
	Revision *int64 `json:"revision,omitempty"`
	// ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
	GnetId *string `json:"gnetId,omitempty"`
	// Tags associated with dashboard.
	Tags []string `json:"tags,omitempty"`
	// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
	Timezone *string `json:"timezone,omitempty"`
	// Whether a dashboard is editable or not.
	Editable *bool `json:"editable,omitempty"`
	// Configuration of dashboard cursor sync behavior.
	// Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
	GraphTooltip *DashboardCursorSync `json:"graphTooltip,omitempty"`
	// Time range for dashboard.
	// Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
	Time *struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"time,omitempty"`
	// Configuration of the time picker shown at the top of a dashboard.
	Timepicker *TimePickerConfig `json:"timepicker,omitempty"`
	// The month that the fiscal year starts on.  0 = January, 11 = December
	FiscalYearStartMonth *uint8 `json:"fiscalYearStartMonth,omitempty"`
	// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
	// This will keep data "moving left" regardless of the query refresh rate. This setting helps
	// avoid dashboards presenting stale live data
	LiveNow *bool `json:"liveNow,omitempty"`
	// Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
	WeekStart *string `json:"weekStart,omitempty"`
	// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
	Refresh *StringOrBool `json:"refresh,omitempty"`
	// Version of the JSON schema, incremented each time a Grafana update brings
	// changes to said schema.
	SchemaVersion uint16 `json:"schemaVersion"`
	// Version of the dashboard, incremented each time the dashboard is updated.
	Version *uint32 `json:"version,omitempty"`
	// List of dashboard panels
	Panels []PanelOrRowPanel `json:"panels,omitempty"`
	// Configured template variables
	Templating *struct {
		// List of configured template variables with their saved values along with some other metadata
		List []VariableModel `json:"list,omitempty"`
	} `json:"templating,omitempty"`
	// Contains the list of annotations that are associated with the dashboard.
	// Annotations are used to overlay event markers and overlay event tags on graphs.
	// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
	// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
	Annotations *AnnotationContainer `json:"annotations,omitempty"`
	// Links with references to other dashboards or external websites.
	Links []DashboardLink `json:"links,omitempty"`
	// Snapshot options. They are present only if the dashboard is a snapshot.
	Snapshot *Snapshot `json:"snapshot,omitempty"`
}

// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most dashboards
type AnnotationTarget struct {
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Limit int64 `json:"limit"`
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	MatchAny bool `json:"matchAny"`
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Tags []string `json:"tags"`
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Type string `json:"type"`
}

type AnnotationPanelFilter struct {
	// Should the specified panels be included or excluded
	Exclude *bool `json:"exclude,omitempty"`
	// Panel IDs that should be included or excluded
	Ids []uint8 `json:"ids"`
}

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
type AnnotationContainer struct {
	// List of annotations
	List []AnnotationQuery `json:"list,omitempty"`
}

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
type AnnotationQuery struct {
	// Name of annotation.
	Name string `json:"name"`
	// Datasource where the annotations data is
	Datasource DataSourceRef `json:"datasource"`
	// When enabled the annotation query is issued with every dashboard refresh
	Enable bool `json:"enable"`
	// Annotation queries can be toggled on or off at the top of the dashboard.
	// When hide is true, the toggle is not shown in the dashboard.
	Hide *bool `json:"hide,omitempty"`
	// Color to use for the annotation event markers
	IconColor string `json:"iconColor"`
	// Filters to apply when fetching annotations
	Filter *AnnotationPanelFilter `json:"filter,omitempty"`
	// TODO.. this should just be a normal query target
	Target *AnnotationTarget `json:"target,omitempty"`
	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	Type *string `json:"type,omitempty"`
	// Set to 1 for the standard annotation query all dashboards have by default.
	BuiltIn *float64 `json:"builtIn,omitempty"`
}

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type VariableModel struct {
	// Type of variable
	Type VariableType `json:"type"`
	// Name of variable
	Name string `json:"name"`
	// Optional display name
	Label *string `json:"label,omitempty"`
	// Visibility configuration for the variable
	Hide *VariableHide `json:"hide,omitempty"`
	// Whether the variable value should be managed by URL query params or not
	SkipUrlSync *bool `json:"skipUrlSync,omitempty"`
	// Description of variable. It can be defined but `null`.
	Description *string `json:"description,omitempty"`
	// Query used to fetch values for a variable
	Query *StringOrAny `json:"query,omitempty"`
	// Data source used to fetch values for a variable. It can be defined but `null`.
	Datasource *DataSourceRef `json:"datasource,omitempty"`
	// Shows current selected variable text/value on the dashboard
	Current *VariableOption `json:"current,omitempty"`
	// Whether multiple values can be selected or not from variable value list
	Multi *bool `json:"multi,omitempty"`
	// Options that can be selected for a variable.
	Options []VariableOption `json:"options,omitempty"`
	Refresh *VariableRefresh `json:"refresh,omitempty"`
	// Options sort order
	Sort *VariableSort `json:"sort,omitempty"`
}

// Option to be selected in a variable.
type VariableOption struct {
	// Whether the option is selected or not
	Selected *bool `json:"selected,omitempty"`
	// Text to be displayed for the option
	Text StringOrArrayOfString `json:"text"`
	// Value of the option
	Value StringOrArrayOfString `json:"value"`
}

// Options to config when to refresh a variable
// `0`: Never refresh the variable
// `1`: Queries the data source every time the dashboard loads.
// `2`: Queries the data source when the dashboard time range changes.
type VariableRefresh int64

const (
	VariableRefreshNever              VariableRefresh = 0
	VariableRefreshOnDashboardLoad    VariableRefresh = 1
	VariableRefreshOnTimeRangeChanged VariableRefresh = 2
)

// Determine if the variable shows on dashboard
// Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).
type VariableHide int64

const (
	VariableHideDontHide     VariableHide = 0
	VariableHideHideLabel    VariableHide = 1
	VariableHideHideVariable VariableHide = 2
)

// Sort variable options
// Accepted values are:
// `0`: No sorting
// `1`: Alphabetical ASC
// `2`: Alphabetical DESC
// `3`: Numerical ASC
// `4`: Numerical DESC
// `5`: Alphabetical Case Insensitive ASC
// `6`: Alphabetical Case Insensitive DESC
// `7`: Natural ASC
// `8`: Natural DESC
type VariableSort int64

const (
	VariableSortDisabled                        VariableSort = 0
	VariableSortAlphabeticalAsc                 VariableSort = 1
	VariableSortAlphabeticalDesc                VariableSort = 2
	VariableSortNumericalAsc                    VariableSort = 3
	VariableSortNumericalDesc                   VariableSort = 4
	VariableSortAlphabeticalCaseInsensitiveAsc  VariableSort = 5
	VariableSortAlphabeticalCaseInsensitiveDesc VariableSort = 6
	VariableSortNaturalAsc                      VariableSort = 7
	VariableSortNaturalDesc                     VariableSort = 8
)

// Ref to a DataSource instance
type DataSourceRef struct {
	// The plugin type-id
	Type *string `json:"type,omitempty"`
	// Specific datasource instance
	Uid *string `json:"uid,omitempty"`
}

// Links with references to other dashboards or external resources
type DashboardLink struct {
	// Title to display with the link
	Title string `json:"title"`
	// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
	Type DashboardLinkType `json:"type"`
	// Icon name to be displayed with the link
	Icon string `json:"icon"`
	// Tooltip to display when the user hovers their mouse over it
	Tooltip string `json:"tooltip"`
	// Link URL. Only required/valid if the type is link
	Url *string `json:"url,omitempty"`
	// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
	Tags []string `json:"tags"`
	// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
	AsDropdown bool `json:"asDropdown"`
	// If true, the link will be opened in a new tab
	TargetBlank bool `json:"targetBlank"`
	// If true, includes current template variables values in the link as query params
	IncludeVars bool `json:"includeVars"`
	// If true, includes current time range in the link as query params
	KeepTime bool `json:"keepTime"`
}

// Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
type DashboardLinkType string

const (
	DashboardLinkTypeLink       DashboardLinkType = "link"
	DashboardLinkTypeDashboards DashboardLinkType = "dashboards"
)

// Dashboard variable type
// `query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.
// `adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).
// `constant`: 	Define a hidden constant.
// `datasource`: Quickly change the data source for an entire dashboard.
// `interval`: Interval variables represent time spans.
// `textbox`: Display a free text input field with an optional default value.
// `custom`: Define the variable options manually using a comma-separated list.
// `system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables
type VariableType string

const (
	VariableTypeQuery      VariableType = "query"
	VariableTypeAdhoc      VariableType = "adhoc"
	VariableTypeConstant   VariableType = "constant"
	VariableTypeDatasource VariableType = "datasource"
	VariableTypeInterval   VariableType = "interval"
	VariableTypeTextbox    VariableType = "textbox"
	VariableTypeCustom     VariableType = "custom"
	VariableTypeSystem     VariableType = "system"
)

// Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.
// Continuous color interpolates a color using the percentage of a value relative to min and max.
// Accepted values are:
// `thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold
// `palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations
// `palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations
// `continuous-GrYlRd`: ontinuous Green-Yellow-Red palette mode
// `continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode
// `continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode
// `continuous-YlRd`: Continuous Yellow-Red palette mode
// `continuous-BlPu`: Continuous Blue-Purple palette mode
// `continuous-YlBl`: Continuous Yellow-Blue palette mode
// `continuous-blues`: Continuous Blue palette mode
// `continuous-reds`: Continuous Red palette mode
// `continuous-greens`: Continuous Green palette mode
// `continuous-purples`: Continuous Purple palette mode
// `shades`: Shades of a single color. Specify a single color, useful in an override rule.
// `fixed`: Fixed color mode. Specify a single color, useful in an override rule.
type FieldColorModeId string

const (
	FieldColorModeIdThresholds           FieldColorModeId = "thresholds"
	FieldColorModeIdPaletteClassic       FieldColorModeId = "palette-classic"
	FieldColorModeIdPaletteClassicByName FieldColorModeId = "palette-classic-by-name"
	FieldColorModeIdContinuousGrYlRd     FieldColorModeId = "continuous-GrYlRd"
	FieldColorModeIdContinuousRdYlGr     FieldColorModeId = "continuous-RdYlGr"
	FieldColorModeIdContinuousBlYlRd     FieldColorModeId = "continuous-BlYlRd"
	FieldColorModeIdContinuousYlRd       FieldColorModeId = "continuous-YlRd"
	FieldColorModeIdContinuousBlPu       FieldColorModeId = "continuous-BlPu"
	FieldColorModeIdContinuousYlBl       FieldColorModeId = "continuous-YlBl"
	FieldColorModeIdContinuousBlues      FieldColorModeId = "continuous-blues"
	FieldColorModeIdContinuousReds       FieldColorModeId = "continuous-reds"
	FieldColorModeIdContinuousGreens     FieldColorModeId = "continuous-greens"
	FieldColorModeIdContinuousPurples    FieldColorModeId = "continuous-purples"
	FieldColorModeIdFixed                FieldColorModeId = "fixed"
	FieldColorModeIdShades               FieldColorModeId = "shades"
)

// Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
type FieldColorSeriesByMode string

const (
	FieldColorSeriesByModeMin  FieldColorSeriesByMode = "min"
	FieldColorSeriesByModeMax  FieldColorSeriesByMode = "max"
	FieldColorSeriesByModeLast FieldColorSeriesByMode = "last"
)

// Map a field to a color.
type FieldColor struct {
	// The main color scheme mode.
	Mode FieldColorModeId `json:"mode"`
	// The fixed color value for fixed or shades color modes.
	FixedColor *string `json:"fixedColor,omitempty"`
	// Some visualizations need to know how to assign a series color from by value color schemes.
	SeriesBy *FieldColorSeriesByMode `json:"seriesBy,omitempty"`
}

// Position and dimensions of a panel in the grid
type GridPos struct {
	// Panel height. The height is the number of rows from the top edge of the panel.
	H uint32 `json:"h"`
	// Panel width. The width is the number of columns from the left edge of the panel.
	W uint32 `json:"w"`
	// Panel x. The x coordinate is the number of columns from the left edge of the grid
	X uint32 `json:"x"`
	// Panel y. The y coordinate is the number of rows from the top edge of the grid
	Y uint32 `json:"y"`
	// Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
	Static *bool `json:"static,omitempty"`
}

// User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded
// They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.
type Threshold struct {
	// Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
	// Nulls currently appear here when serializing -Infinity to JSON.
	Value *float64 `json:"value"`
	// Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
	Color string `json:"color"`
}

// Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).
type ThresholdsMode string

const (
	ThresholdsModeAbsolute   ThresholdsMode = "absolute"
	ThresholdsModePercentage ThresholdsMode = "percentage"
)

// Thresholds configuration for the panel
type ThresholdsConfig struct {
	// Thresholds mode.
	Mode ThresholdsMode `json:"mode"`
	// Must be sorted by 'value', first value is always -Infinity
	Steps []Threshold `json:"steps"`
}

// Allow to transform the visual representation of specific data values in a visualization, irrespective of their original units
type ValueMapping ValueMapOrRangeMapOrRegexMapOrSpecialValueMap

// Supported value mapping types
// `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
type MappingType string

const (
	MappingTypeValueToText  MappingType = "value"
	MappingTypeRangeToText  MappingType = "range"
	MappingTypeRegexToText  MappingType = "regex"
	MappingTypeSpecialValue MappingType = "special"
)

// Maps text values to a color or different display text and color.
// For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
type ValueMap struct {
	Type string `json:"type"`
	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	Options map[string]ValueMappingResult `json:"options"`
}

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
type RangeMap struct {
	Type string `json:"type"`
	// Range to match against and the result to apply when the value is within the range
	Options struct {
		// Min value of the range. It can be null which means -Infinity
		From *float64 `json:"from"`
		// Max value of the range. It can be null which means +Infinity
		To *float64 `json:"to"`
		// Config to apply when the value is within the range
		Result ValueMappingResult `json:"result"`
	} `json:"options"`
}

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
type RegexMap struct {
	Type string `json:"type"`
	// Regular expression to match against and the result to apply when the value matches the regex
	Options struct {
		// Regular expression to match against
		Pattern string `json:"pattern"`
		// Config to apply when the value matches the regex
		Result ValueMappingResult `json:"result"`
	} `json:"options"`
}

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
type SpecialValueMap struct {
	Type    string `json:"type"`
	Options struct {
		// Special value to match against
		Match SpecialValueMatch `json:"match"`
		// Config to apply when the value matches the special value
		Result ValueMappingResult `json:"result"`
	} `json:"options"`
}

// Special value types supported by the `SpecialValueMap`
type SpecialValueMatch string

const (
	SpecialValueMatchTrue       SpecialValueMatch = "true"
	SpecialValueMatchFalse      SpecialValueMatch = "false"
	SpecialValueMatchNull       SpecialValueMatch = "null"
	SpecialValueMatchNaN        SpecialValueMatch = "nan"
	SpecialValueMatchNullAndNan SpecialValueMatch = "null+nan"
	SpecialValueMatchEmpty      SpecialValueMatch = "empty"
)

// Result used as replacement with text and color when the value matches
type ValueMappingResult struct {
	// Text to display when the value matches
	Text *string `json:"text,omitempty"`
	// Text to use when the value matches
	Color *string `json:"color,omitempty"`
	// Icon to display when the value matches. Only specific visualizations.
	Icon *string `json:"icon,omitempty"`
	// Position in the mapping array. Only used internally.
	Index *int32 `json:"index,omitempty"`
}

// Transformations allow to manipulate data returned by a query before the system applies a visualization.
// Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
// use the output of one transformation as the input to another transformation, etc.
type DataTransformerConfig struct {
	// Unique identifier of transformer
	Id string `json:"id"`
	// Disabled transformations are skipped
	Disabled *bool `json:"disabled,omitempty"`
	// Optional frame matcher. When missing it will be applied to all results
	Filter *MatcherConfig `json:"filter,omitempty"`
	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	Options any `json:"options"`
}

// Time picker configuration
// It defines the default config for the time picker and the refresh picker for the specific dashboard.
type TimePickerConfig struct {
	// Whether timepicker is visible or not.
	Hidden bool `json:"hidden"`
	// Interval options available in the refresh picker dropdown.
	RefreshIntervals []string `json:"refresh_intervals"`
	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	TimeOptions []string `json:"time_options"`
}

// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
type DashboardCursorSync int64

const (
	DashboardCursorSyncOff       DashboardCursorSync = 0
	DashboardCursorSyncCrosshair DashboardCursorSync = 1
	DashboardCursorSyncTooltip   DashboardCursorSync = 2
)

// A dashboard snapshot shares an interactive dashboard publicly.
// It is a read-only version of a dashboard, and is not editable.
// It is possible to create a snapshot of a snapshot.
// Grafana strips away all sensitive information from the dashboard.
// Sensitive information stripped: queries (metric, template,annotation) and panel links.
type Snapshot struct {
	// Time when the snapshot was created
	Created string `json:"created"`
	// Time when the snapshot expires, default is never to expire
	Expires string `json:"expires"`
	// Is the snapshot saved in an external grafana instance
	External bool `json:"external"`
	// external url, if snapshot was shared in external grafana instance
	ExternalUrl string `json:"externalUrl"`
	// Unique identifier of the snapshot
	Id uint32 `json:"id"`
	// Optional, defined the unique key of the snapshot, required if external is true
	Key string `json:"key"`
	// Optional, name of the snapshot
	Name string `json:"name"`
	// org id of the snapshot
	OrgId uint32 `json:"orgId"`
	// last time when the snapshot was updated
	Updated string `json:"updated"`
	// url of the snapshot, if snapshot was shared internally
	Url *string `json:"url,omitempty"`
	// user id of the snapshot creator
	UserId uint32 `json:"userId"`
}

// Dashboard panels are the basic visualization building blocks.
type Panel struct {
	// The panel plugin type id. This is used to find the plugin to display the panel.
	Type string `json:"type"`
	// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
	Id *uint32 `json:"id,omitempty"`
	// The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
	PluginVersion *string `json:"pluginVersion,omitempty"`
	// Tags for the panel.
	Tags []string `json:"tags,omitempty"`
	// Depends on the panel plugin. See the plugin documentation for details.
	Targets []cogvariants.Dataquery `json:"targets,omitempty"`
	// Panel title.
	Title *string `json:"title,omitempty"`
	// Panel description.
	Description *string `json:"description,omitempty"`
	// Whether to display the panel without a background.
	Transparent *bool `json:"transparent,omitempty"`
	// The datasource used in all targets.
	Datasource *DataSourceRef `json:"datasource,omitempty"`
	// Grid position.
	GridPos *GridPos `json:"gridPos,omitempty"`
	// Panel links.
	Links []DashboardLink `json:"links,omitempty"`
	// Name of template variable to repeat for.
	Repeat *string `json:"repeat,omitempty"`
	// Direction to repeat in if 'repeat' is set.
	// `h` for horizontal, `v` for vertical.
	RepeatDirection *PanelRepeatDirection `json:"repeatDirection,omitempty"`
	// Option for repeated panels that controls max items per row
	// Only relevant for horizontally repeated panels
	MaxPerRow *float64 `json:"maxPerRow,omitempty"`
	// The maximum number of data points that the panel queries are retrieving.
	MaxDataPoints *float64 `json:"maxDataPoints,omitempty"`
	// List of transformations that are applied to the panel data before rendering.
	// When there are multiple transformations, Grafana applies them in the order they are listed.
	// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
	Transformations []DataTransformerConfig `json:"transformations,omitempty"`
	// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
	// This value must be formatted as a number followed by a valid time
	// identifier like: "40s", "3d", etc.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	Interval *string `json:"interval,omitempty"`
	// Overrides the relative time range for individual panels,
	// which causes them to be different than what is selected in
	// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
	// time periods or days on the same dashboard.
	// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
	// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	TimeFrom *string `json:"timeFrom,omitempty"`
	// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
	// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
	// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
	// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
	TimeShift *string `json:"timeShift,omitempty"`
	// Controls if the timeFrom or timeShift overrides are shown in the panel header
	HideTimeOverride *bool `json:"hideTimeOverride,omitempty"`
	// Dynamically load the panel
	LibraryPanel *LibraryPanelRef `json:"libraryPanel,omitempty"`
	// It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
	Options any `json:"options,omitempty"`
	// Field options allow you to change how the data is displayed in your visualizations.
	FieldConfig *FieldConfigSource `json:"fieldConfig,omitempty"`
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
type FieldConfigSource struct {
	// Defaults are the options applied to all fields.
	Defaults FieldConfig `json:"defaults"`
	// Overrides are the options applied to specific fields overriding the defaults.
	Overrides []struct {
		Matcher    MatcherConfig        `json:"matcher"`
		Properties []DynamicConfigValue `json:"properties"`
	} `json:"overrides"`
}

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
type LibraryPanelRef struct {
	// Library panel name
	Name string `json:"name"`
	// Library panel uid
	Uid string `json:"uid"`
}

// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
type MatcherConfig struct {
	// The matcher id. This is used to find the matcher implementation from registry.
	Id string `json:"id"`
	// The matcher options. This is specific to the matcher implementation.
	Options any `json:"options,omitempty"`
}

type DynamicConfigValue struct {
	Id    string `json:"id"`
	Value any    `json:"value,omitempty"`
}

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
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
	Mappings []ValueMapping `json:"mappings,omitempty"`
	// Map numeric values to states
	Thresholds *ThresholdsConfig `json:"thresholds,omitempty"`
	// Panel color configuration
	Color *FieldColor `json:"color,omitempty"`
	// The behavior when clicking on a result
	Links []any `json:"links,omitempty"`
	// Alternative to empty string
	NoValue *string `json:"noValue,omitempty"`
	// custom is specified by the FieldConfig field
	// in panel plugin schemas.
	Custom any `json:"custom,omitempty"`
}

// Row panel
type RowPanel struct {
	// The panel type
	Type string `json:"type"`
	// Whether this row should be collapsed or not.
	Collapsed bool `json:"collapsed"`
	// Row title
	Title *string `json:"title,omitempty"`
	// Name of default datasource for the row
	Datasource *DataSourceRef `json:"datasource,omitempty"`
	// Row grid position
	GridPos *GridPos `json:"gridPos,omitempty"`
	// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
	Id uint32 `json:"id"`
	// List of panels in the row
	Panels []Panel `json:"panels"`
	// Name of template variable to repeat for.
	Repeat *string `json:"repeat,omitempty"`
}

type PanelRepeatDirection string

const (
	PanelRepeatDirectionH PanelRepeatDirection = "h"
	PanelRepeatDirectionV PanelRepeatDirection = "v"
)

type PanelOrRowPanel struct {
	Panel    *Panel    `json:"Panel,omitempty"`
	RowPanel *RowPanel `json:"RowPanel,omitempty"`
}

type StringOrAny struct {
	String *string `json:"String,omitempty"`
	Any    any     `json:"Any,omitempty"`
}

type StringOrArrayOfString struct {
	String        *string  `json:"String,omitempty"`
	ArrayOfString []string `json:"ArrayOfString,omitempty"`
}

type StringOrBool struct {
	String *string `json:"String,omitempty"`
	Bool   *bool   `json:"Bool,omitempty"`
}

type ValueMapOrRangeMapOrRegexMapOrSpecialValueMap struct {
	ValueMap        *ValueMap        `json:"ValueMap,omitempty"`
	RangeMap        *RangeMap        `json:"RangeMap,omitempty"`
	RegexMap        *RegexMap        `json:"RegexMap,omitempty"`
	SpecialValueMap *SpecialValueMap `json:"SpecialValueMap,omitempty"`
}
