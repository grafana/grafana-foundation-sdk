// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[dashboard.Panel] = (*PanelBuilder)(nil)

// Dashboard panels are the basic visualization building blocks.
type PanelBuilder struct {
	internal *dashboard.Panel
	errors   map[string]cog.BuildErrors
}

func NewPanelBuilder() *PanelBuilder {
	resource := &dashboard.Panel{}
	builder := &PanelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "heatmap"

	return builder
}

func (builder *PanelBuilder) Build() (dashboard.Panel, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Panel", err)...)
	}

	if len(errs) != 0 {
		return dashboard.Panel{}, errs
	}

	return *builder.internal, nil
}

// Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
func (builder *PanelBuilder) Id(id uint32) *PanelBuilder {
	builder.internal.Id = &id

	return builder
}

// Depends on the panel plugin. See the plugin documentation for details.
func (builder *PanelBuilder) WithTarget(targets cog.Builder[cogvariants.Dataquery]) *PanelBuilder {
	targetsResource, err := targets.Build()
	if err != nil {
		builder.errors["targets"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Targets = append(builder.internal.Targets, targetsResource)

	return builder
}

// Panel title.
func (builder *PanelBuilder) Title(title string) *PanelBuilder {
	builder.internal.Title = &title

	return builder
}

// Panel description.
func (builder *PanelBuilder) Description(description string) *PanelBuilder {
	builder.internal.Description = &description

	return builder
}

// Whether to display the panel without a background.
func (builder *PanelBuilder) Transparent(transparent bool) *PanelBuilder {
	builder.internal.Transparent = &transparent

	return builder
}

// The datasource used in all targets.
func (builder *PanelBuilder) Datasource(datasource dashboard.DataSourceRef) *PanelBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Panel height. The height is the number of rows from the top edge of the panel.
func (builder *PanelBuilder) Height(h uint32) *PanelBuilder {
	if !(h > 0) {
		builder.errors["h"] = cog.MakeBuildErrors("h", errors.New("h must be > 0"))
		return builder
	}
	if builder.internal.GridPos == nil {
		builder.internal.GridPos = &dashboard.GridPos{}
	}
	builder.internal.GridPos.H = h

	return builder
}

// Panel width. The width is the number of columns from the left edge of the panel.
func (builder *PanelBuilder) Span(w uint32) *PanelBuilder {
	if !(w > 0) {
		builder.errors["w"] = cog.MakeBuildErrors("w", errors.New("w must be > 0"))
		return builder
	}
	if !(w <= 24) {
		builder.errors["w"] = cog.MakeBuildErrors("w", errors.New("w must be <= 24"))
		return builder
	}
	if builder.internal.GridPos == nil {
		builder.internal.GridPos = &dashboard.GridPos{}
	}
	builder.internal.GridPos.W = w

	return builder
}

// Panel links.
func (builder *PanelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *PanelBuilder {
	linksResources := make([]dashboard.DashboardLink, 0, len(links))
	for _, r1 := range links {
		linksDepth1, err := r1.Build()
		if err != nil {
			builder.errors["links"] = err.(cog.BuildErrors)
			return builder
		}
		linksResources = append(linksResources, linksDepth1)
	}
	builder.internal.Links = linksResources

	return builder
}

// Name of template variable to repeat for.
func (builder *PanelBuilder) Repeat(repeat string) *PanelBuilder {
	builder.internal.Repeat = &repeat

	return builder
}

// Direction to repeat in if 'repeat' is set.
// `h` for horizontal, `v` for vertical.
func (builder *PanelBuilder) RepeatDirection(repeatDirection dashboard.PanelRepeatDirection) *PanelBuilder {
	builder.internal.RepeatDirection = &repeatDirection

	return builder
}

// Option for repeated panels that controls max items per row
// Only relevant for horizontally repeated panels
func (builder *PanelBuilder) MaxPerRow(maxPerRow float64) *PanelBuilder {
	builder.internal.MaxPerRow = &maxPerRow

	return builder
}

// The maximum number of data points that the panel queries are retrieving.
func (builder *PanelBuilder) MaxDataPoints(maxDataPoints float64) *PanelBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// List of transformations that are applied to the panel data before rendering.
// When there are multiple transformations, Grafana applies them in the order they are listed.
// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
func (builder *PanelBuilder) WithTransformation(transformations dashboard.DataTransformerConfig) *PanelBuilder {
	builder.internal.Transformations = append(builder.internal.Transformations, transformations)

	return builder
}

// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
// This value must be formatted as a number followed by a valid time
// identifier like: "40s", "3d", etc.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelBuilder) Interval(interval string) *PanelBuilder {
	builder.internal.Interval = &interval

	return builder
}

// Overrides the relative time range for individual panels,
// which causes them to be different than what is selected in
// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
// time periods or days on the same dashboard.
// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelBuilder) TimeFrom(timeFrom string) *PanelBuilder {
	builder.internal.TimeFrom = &timeFrom

	return builder
}

// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelBuilder) TimeShift(timeShift string) *PanelBuilder {
	builder.internal.TimeShift = &timeShift

	return builder
}

// Controls if the timeFrom or timeShift overrides are shown in the panel header
func (builder *PanelBuilder) HideTimeOverride(hideTimeOverride bool) *PanelBuilder {
	builder.internal.HideTimeOverride = &hideTimeOverride

	return builder
}

// Dynamically load the panel
func (builder *PanelBuilder) LibraryPanel(libraryPanel dashboard.LibraryPanelRef) *PanelBuilder {
	builder.internal.LibraryPanel = &libraryPanel

	return builder
}

// Sets panel queries cache timeout.
func (builder *PanelBuilder) CacheTimeout(cacheTimeout string) *PanelBuilder {
	builder.internal.CacheTimeout = &cacheTimeout

	return builder
}

// Overrides the data source configured time-to-live for a query cache item in milliseconds
func (builder *PanelBuilder) QueryCachingTTL(queryCachingTTL float64) *PanelBuilder {
	builder.internal.QueryCachingTTL = &queryCachingTTL

	return builder
}

// The display value for this field.  This supports template variables blank is auto
func (builder *PanelBuilder) DisplayName(displayName string) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.DisplayName = &displayName

	return builder
}

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
func (builder *PanelBuilder) Unit(unit string) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Unit = &unit

	return builder
}

// Specify the number of decimals Grafana includes in the rendered value.
// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
// To display all decimals, set the unit to `String`.
func (builder *PanelBuilder) Decimals(decimals float64) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Decimals = &decimals

	return builder
}

// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *PanelBuilder) Min(min float64) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Min = &min

	return builder
}

// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *PanelBuilder) Max(max float64) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Max = &max

	return builder
}

// Convert input values into a display string
func (builder *PanelBuilder) Mappings(mappings []dashboard.ValueMapping) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Mappings = mappings

	return builder
}

// Map numeric values to states
func (builder *PanelBuilder) Thresholds(thresholds cog.Builder[dashboard.ThresholdsConfig]) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	thresholdsResource, err := thresholds.Build()
	if err != nil {
		builder.errors["fieldConfig.defaults.thresholds"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.FieldConfig.Defaults.Thresholds = &thresholdsResource

	return builder
}

// Alternative to empty string
func (builder *PanelBuilder) NoValue(noValue string) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.NoValue = &noValue

	return builder
}

// Defaults are the options applied to all fields.
func (builder *PanelBuilder) Defaults(defaults dashboard.FieldConfig) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults = defaults

	return builder
}

// Overrides are the options applied to specific fields overriding the defaults.
func (builder *PanelBuilder) WithOverride(matcher dashboard.MatcherConfig, properties []dashboard.DynamicConfigValue) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Overrides = append(builder.internal.FieldConfig.Overrides, struct {
		Matcher    dashboard.MatcherConfig        `json:"matcher"`
		Properties []dashboard.DynamicConfigValue `json:"properties"`
	}{
		Matcher:    matcher,
		Properties: properties,
	})

	return builder
}

// Controls if the heatmap should be calculated from data
func (builder *PanelBuilder) Calculate(calculate bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).Calculate = &calculate

	return builder
}

// Calculation options for the heatmap
func (builder *PanelBuilder) Calculation(calculation cog.Builder[common.HeatmapCalculationOptions]) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	calculationResource, err := calculation.Build()
	if err != nil {
		builder.errors["options.calculation"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Options.(*Options).Calculation = &calculationResource

	return builder
}

// Controls the color options
func (builder *PanelBuilder) Color(color HeatmapColorOptions) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).Color = color

	return builder
}

// Filters values between a given range
func (builder *PanelBuilder) FilterValues(filterValues FilterValueRange) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).FilterValues = &filterValues

	return builder
}

// Controls tick alignment and value name when not calculating from data
func (builder *PanelBuilder) RowsFrame(rowsFrame RowsHeatmapOptions) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).RowsFrame = &rowsFrame

	return builder
}

//	| *{
//		layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
//	}
//
// Controls the display of the value in the cell
func (builder *PanelBuilder) ShowValue(showValue common.VisibilityMode) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).ShowValue = showValue

	return builder
}

// Controls gap between cells
func (builder *PanelBuilder) CellGap(cellGap uint8) *PanelBuilder {
	if !(cellGap <= 25) {
		builder.errors["cellGap"] = cog.MakeBuildErrors("cellGap", errors.New("cellGap must be <= 25"))
		return builder
	}
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).CellGap = &cellGap

	return builder
}

// Controls cell radius
func (builder *PanelBuilder) CellRadius(cellRadius float32) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).CellRadius = &cellRadius

	return builder
}

// Controls cell value unit
func (builder *PanelBuilder) CellValues(cellValues CellValues) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).CellValues = &cellValues

	return builder
}

// Controls yAxis placement
func (builder *PanelBuilder) YAxis(yAxis YAxisConfig) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).YAxis = yAxis

	return builder
}

//	| *{
//		axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
//	}
//
// Controls legend options
func (builder *PanelBuilder) Legend(legend HeatmapLegend) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).Legend = legend

	return builder
}

// Controls tooltip options
func (builder *PanelBuilder) Tooltip(tooltip HeatmapTooltip) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).Tooltip = tooltip

	return builder
}

// Controls exemplar options
func (builder *PanelBuilder) ExemplarsColor(exemplars ExemplarConfig) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).Exemplars = exemplars

	return builder
}

func (builder *PanelBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	if builder.internal.FieldConfig.Defaults.Custom == nil {
		builder.internal.FieldConfig.Defaults.Custom = &FieldConfig{}
	}
	scaleDistributionResource, err := scaleDistribution.Build()
	if err != nil {
		builder.errors["fieldConfig.defaults.custom.scaleDistribution"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.FieldConfig.Defaults.Custom.(*FieldConfig).ScaleDistribution = &scaleDistributionResource

	return builder
}

func (builder *PanelBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	if builder.internal.FieldConfig.Defaults.Custom == nil {
		builder.internal.FieldConfig.Defaults.Custom = &FieldConfig{}
	}
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors["fieldConfig.defaults.custom.hideFrom"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.FieldConfig.Defaults.Custom.(*FieldConfig).HideFrom = &hideFromResource

	return builder
}

func (builder *PanelBuilder) applyDefaults() {
	builder.Transparent(false)
	builder.Height(9)
	builder.Span(12)
	builder.Calculate(false)
	builder.Color(HeatmapColorOptions{
		Exponent: 0.5,
		Fill:     "dark-orange",
		Reverse:  false,
		Scheme:   "Oranges",
		Steps:    64,
	})
	builder.FilterValues(FilterValueRange{
		Le: cog.ToPtr[float32](1e-09),
	})
	builder.ShowValue("auto")
	builder.CellGap(1)
	builder.Legend(HeatmapLegend{
		Show: true,
	})
	builder.Tooltip(HeatmapTooltip{
		Mode:           "single",
		ShowColorScale: cog.ToPtr[bool](false),
		YHistogram:     cog.ToPtr[bool](false),
	})
	builder.ExemplarsColor(ExemplarConfig{
		Color: "rgba(255,0,255,0.7)",
	})
}
