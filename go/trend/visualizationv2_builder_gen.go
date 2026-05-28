// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package trend

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

var _ cog.Builder[dashboardv2.VizConfigKind] = (*VisualizationV2Builder)(nil)

type VisualizationV2Builder struct {
	internal *dashboardv2.VizConfigKind
	errors   cog.BuildErrors
}

func NewVisualizationV2Builder() *VisualizationV2Builder {
	resource := dashboardv2.NewVizConfigKind()
	builder := &VisualizationV2Builder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "VizConfig"
	builder.internal.Group = "trend"

	return builder
}

func (builder *VisualizationV2Builder) Build() (dashboardv2.VizConfigKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.VizConfigKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.VizConfigKind{}, cog.MakeBuildErrors("trend.visualizationV2", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VisualizationV2Builder) RecordError(path string, err error) *VisualizationV2Builder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The display value for this field.  This supports template variables blank is auto
func (builder *VisualizationV2Builder) DisplayName(displayName string) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.DisplayName = &displayName

	return builder
}

// This can be used by data sources that return and explicit naming structure for values and labels
// When this property is configured, this value is used rather than the default naming strategy.
func (builder *VisualizationV2Builder) DisplayNameFromDS(displayNameFromDS string) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.DisplayNameFromDS = &displayNameFromDS

	return builder
}

// Human readable field metadata
func (builder *VisualizationV2Builder) Description(description string) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Description = &description

	return builder
}

// An explicit path to the field in the datasource.  When the frame meta includes a path,
// This will default to `${frame.meta.path}/${field.name}
//
// When defined, this value can be used as an identifier within the datasource scope, and
// may be used to update the results
func (builder *VisualizationV2Builder) Path(path string) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Path = &path

	return builder
}

// Unit a field should use. The unit you select is applied to all fields except time.
// You can use the units ID available in Grafana or a custom unit.
// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
// As custom unit, you can use the following formats:
// `suffix:<suffix>` for custom unit that should go after value.
// `prefix:<prefix>` for custom unit that should go before value.
// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
// `count:<unit>` for a custom count unit.
// `currency:<unit>` for custom a currency unit.
func (builder *VisualizationV2Builder) Unit(unit string) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Unit = &unit

	return builder
}

// Specify the number of decimals Grafana includes in the rendered value.
// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
// To display all decimals, set the unit to `String`.
func (builder *VisualizationV2Builder) Decimals(decimals float64) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Decimals = &decimals

	return builder
}

// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *VisualizationV2Builder) Min(min float64) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Min = &min

	return builder
}

// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *VisualizationV2Builder) Max(max float64) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Max = &max

	return builder
}

// Convert input values into a display string
func (builder *VisualizationV2Builder) Mappings(mappings []dashboardv2.ValueMapping) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Mappings = mappings

	return builder
}

// Map numeric values to states
func (builder *VisualizationV2Builder) Thresholds(thresholds cog.Builder[dashboardv2.ThresholdsConfig]) *VisualizationV2Builder {
	thresholdsResource, err := thresholds.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Thresholds = &thresholdsResource

	return builder
}

// Panel color configuration
func (builder *VisualizationV2Builder) ColorScheme(color cog.Builder[dashboardv2.FieldColor]) *VisualizationV2Builder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Color = &colorResource

	return builder
}

// The behavior when clicking on a result
func (builder *VisualizationV2Builder) DataLinks(links []any) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.Links = links

	return builder
}

// Define interactive HTTP requests that can be triggered from data visualizations.
func (builder *VisualizationV2Builder) Actions(actions []cog.Builder[dashboardv2.Action]) *VisualizationV2Builder {
	actionsResources := make([]dashboardv2.Action, 0, len(actions))
	for _, r1 := range actions {
		actionsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		actionsResources = append(actionsResources, actionsDepth1)
	}
	builder.internal.Spec.FieldConfig.Defaults.Actions = actionsResources

	return builder
}

// Alternative to empty string
func (builder *VisualizationV2Builder) NoValue(noValue string) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.NoValue = &noValue

	return builder
}

// Calculate min max per field
func (builder *VisualizationV2Builder) FieldMinMax(fieldMinMax bool) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.FieldMinMax = &fieldMinMax

	return builder
}

// How null values should be handled when calculating field stats
// "null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero
func (builder *VisualizationV2Builder) NullValueMode(nullValueMode dashboardv2.NullValueMode) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Defaults.NullValueMode = &nullValueMode

	return builder
}

// Overrides are the options applied to specific fields overriding the defaults.
func (builder *VisualizationV2Builder) Overrides(overrides []cog.Builder[dashboardv2.Dashboardv2FieldConfigSourceOverrides]) *VisualizationV2Builder {
	overridesResources := make([]dashboardv2.Dashboardv2FieldConfigSourceOverrides, 0, len(overrides))
	for _, r1 := range overrides {
		overridesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		overridesResources = append(overridesResources, overridesDepth1)
	}
	builder.internal.Spec.FieldConfig.Overrides = overridesResources

	return builder
}

// Overrides are the options applied to specific fields overriding the defaults.
func (builder *VisualizationV2Builder) Override(matcher dashboardv2.MatcherConfig, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2.Dashboardv2FieldConfigSourceOverrides{
		Matcher:    matcher,
		Properties: properties,
	})

	return builder
}

func (builder *VisualizationV2Builder) Legend(legend cog.Builder[common.VizLegendOptions]) *VisualizationV2Builder {
	if builder.internal.Spec.Options == nil {
		builder.internal.Spec.Options = NewOptions()
	}
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Options.(*Options).Legend = legendResource

	return builder
}

func (builder *VisualizationV2Builder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *VisualizationV2Builder {
	if builder.internal.Spec.Options == nil {
		builder.internal.Spec.Options = NewOptions()
	}
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Options.(*Options).Tooltip = tooltipResource

	return builder
}

// Name of the x field to use (defaults to first number)
func (builder *VisualizationV2Builder) XField(xField string) *VisualizationV2Builder {
	if builder.internal.Spec.Options == nil {
		builder.internal.Spec.Options = NewOptions()
	}
	builder.internal.Spec.Options.(*Options).XField = &xField

	return builder
}

func (builder *VisualizationV2Builder) DrawStyle(drawStyle common.GraphDrawStyle) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).DrawStyle = &drawStyle

	return builder
}

func (builder *VisualizationV2Builder) GradientMode(gradientMode common.GraphGradientMode) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).GradientMode = &gradientMode

	return builder
}

func (builder *VisualizationV2Builder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	thresholdsStyleResource, err := thresholdsStyle.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ThresholdsStyle = &thresholdsStyleResource

	return builder
}

func (builder *VisualizationV2Builder) Transform(transform common.GraphTransform) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).Transform = &transform

	return builder
}

func (builder *VisualizationV2Builder) LineColor(lineColor string) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineColor = &lineColor

	return builder
}

func (builder *VisualizationV2Builder) LineWidth(lineWidth float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineWidth = &lineWidth

	return builder
}

func (builder *VisualizationV2Builder) LineInterpolation(lineInterpolation common.LineInterpolation) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineInterpolation = &lineInterpolation

	return builder
}

func (builder *VisualizationV2Builder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	lineStyleResource, err := lineStyle.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineStyle = &lineStyleResource

	return builder
}

func (builder *VisualizationV2Builder) FillColor(fillColor string) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillColor = &fillColor

	return builder
}

func (builder *VisualizationV2Builder) FillOpacity(fillOpacity float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillOpacity = &fillOpacity

	return builder
}

func (builder *VisualizationV2Builder) ShowPoints(showPoints common.VisibilityMode) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ShowPoints = &showPoints

	return builder
}

func (builder *VisualizationV2Builder) PointSize(pointSize float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).PointSize = &pointSize

	return builder
}

func (builder *VisualizationV2Builder) PointColor(pointColor string) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).PointColor = &pointColor

	return builder
}

func (builder *VisualizationV2Builder) AxisPlacement(axisPlacement common.AxisPlacement) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisPlacement = &axisPlacement

	return builder
}

func (builder *VisualizationV2Builder) AxisColorMode(axisColorMode common.AxisColorMode) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisColorMode = &axisColorMode

	return builder
}

func (builder *VisualizationV2Builder) AxisLabel(axisLabel string) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisLabel = &axisLabel

	return builder
}

func (builder *VisualizationV2Builder) AxisWidth(axisWidth float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisWidth = &axisWidth

	return builder
}

func (builder *VisualizationV2Builder) AxisSoftMin(axisSoftMin float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMin = &axisSoftMin

	return builder
}

func (builder *VisualizationV2Builder) AxisSoftMax(axisSoftMax float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMax = &axisSoftMax

	return builder
}

func (builder *VisualizationV2Builder) AxisGridShow(axisGridShow bool) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisGridShow = &axisGridShow

	return builder
}

func (builder *VisualizationV2Builder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	scaleDistributionResource, err := scaleDistribution.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).ScaleDistribution = &scaleDistributionResource

	return builder
}

func (builder *VisualizationV2Builder) AxisCenteredZero(axisCenteredZero bool) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisCenteredZero = &axisCenteredZero

	return builder
}

func (builder *VisualizationV2Builder) BarAlignment(barAlignment common.BarAlignment) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).BarAlignment = &barAlignment

	return builder
}

func (builder *VisualizationV2Builder) BarWidthFactor(barWidthFactor float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).BarWidthFactor = &barWidthFactor

	return builder
}

func (builder *VisualizationV2Builder) Stacking(stacking cog.Builder[common.StackingConfig]) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	stackingResource, err := stacking.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).Stacking = &stackingResource

	return builder
}

func (builder *VisualizationV2Builder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	hideFromResource, err := hideFrom.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).HideFrom = &hideFromResource

	return builder
}

func (builder *VisualizationV2Builder) InsertNulls(insertNulls common.BoolOrFloat64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).InsertNulls = &insertNulls

	return builder
}

// Indicate if null values should be treated as gaps or connected.
// When the value is a number, it represents the maximum delta in the
// X axis that should be considered connected.  For timeseries, this is milliseconds
func (builder *VisualizationV2Builder) SpanNulls(spanNulls common.BoolOrFloat64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).SpanNulls = &spanNulls

	return builder
}

func (builder *VisualizationV2Builder) FillBelowTo(fillBelowTo string) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillBelowTo = &fillBelowTo

	return builder
}

func (builder *VisualizationV2Builder) PointSymbol(pointSymbol string) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).PointSymbol = &pointSymbol

	return builder
}

func (builder *VisualizationV2Builder) AxisBorderShow(axisBorderShow bool) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisBorderShow = &axisBorderShow

	return builder
}

func (builder *VisualizationV2Builder) BarMaxWidth(barMaxWidth float64) *VisualizationV2Builder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).BarMaxWidth = &barMaxWidth

	return builder
}

// Adds override rules for a specific field, referred to by its name.
func (builder *VisualizationV2Builder) OverrideByName(name string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2.Dashboardv2FieldConfigSourceOverrides{
		Matcher: dashboardv2.MatcherConfig{
			Id:      "byName",
			Options: &name,
		},
		Properties: properties,
	})

	return builder
}

// Adds override rules for the fields whose name match the given regexp.
func (builder *VisualizationV2Builder) OverrideByRegexp(regexp string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2.Dashboardv2FieldConfigSourceOverrides{
		Matcher: dashboardv2.MatcherConfig{
			Id:      "byRegexp",
			Options: &regexp,
		},
		Properties: properties,
	})

	return builder
}

// Adds override rules for all the fields of the given type.
func (builder *VisualizationV2Builder) OverrideByFieldType(fieldType string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2.Dashboardv2FieldConfigSourceOverrides{
		Matcher: dashboardv2.MatcherConfig{
			Id:      "byType",
			Options: &fieldType,
		},
		Properties: properties,
	})

	return builder
}

func (builder *VisualizationV2Builder) OverrideByQuery(queryRefId string, properties []dashboardv2.DynamicConfigValue) *VisualizationV2Builder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2.Dashboardv2FieldConfigSourceOverrides{
		Matcher: dashboardv2.MatcherConfig{
			Id:      "byFrameRefID",
			Options: &queryRefId,
		},
		Properties: properties,
	})

	return builder
}
