// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

var _ cog.Builder[dashboardv2beta1.VizConfigKind] = (*VisualizationBuilder)(nil)

type VisualizationBuilder struct {
	internal *dashboardv2beta1.VizConfigKind
	errors   cog.BuildErrors
}

func NewVisualizationBuilder() *VisualizationBuilder {
	resource := dashboardv2beta1.NewVizConfigKind()
	builder := &VisualizationBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "VizConfig"
	builder.internal.Group = "xychart"

	return builder
}

func (builder *VisualizationBuilder) Build() (dashboardv2beta1.VizConfigKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.VizConfigKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.VizConfigKind{}, cog.MakeBuildErrors("xychart.visualization", builder.errors)
	}

	return *builder.internal, nil
}

// The display value for this field.  This supports template variables blank is auto
func (builder *VisualizationBuilder) DisplayName(displayName string) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.DisplayName = &displayName

	return builder
}

// This can be used by data sources that return and explicit naming structure for values and labels
// When this property is configured, this value is used rather than the default naming strategy.
func (builder *VisualizationBuilder) DisplayNameFromDS(displayNameFromDS string) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.DisplayNameFromDS = &displayNameFromDS

	return builder
}

// Human readable field metadata
func (builder *VisualizationBuilder) Description(description string) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Description = &description

	return builder
}

// An explicit path to the field in the datasource.  When the frame meta includes a path,
// This will default to `${frame.meta.path}/${field.name}
//
// When defined, this value can be used as an identifier within the datasource scope, and
// may be used to update the results
func (builder *VisualizationBuilder) Path(path string) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Path = &path

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
func (builder *VisualizationBuilder) Unit(unit string) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Unit = &unit

	return builder
}

// Specify the number of decimals Grafana includes in the rendered value.
// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
// To display all decimals, set the unit to `String`.
func (builder *VisualizationBuilder) Decimals(decimals float64) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Decimals = &decimals

	return builder
}

// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *VisualizationBuilder) Min(min float64) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Min = &min

	return builder
}

// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *VisualizationBuilder) Max(max float64) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Max = &max

	return builder
}

// Convert input values into a display string
func (builder *VisualizationBuilder) Mappings(mappings []dashboardv2beta1.ValueMapping) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Mappings = mappings

	return builder
}

// Map numeric values to states
func (builder *VisualizationBuilder) Thresholds(thresholds cog.Builder[dashboardv2beta1.ThresholdsConfig]) *VisualizationBuilder {
	thresholdsResource, err := thresholds.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Thresholds = &thresholdsResource

	return builder
}

// Panel color configuration
func (builder *VisualizationBuilder) ColorScheme(color cog.Builder[dashboardv2beta1.FieldColor]) *VisualizationBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Color = &colorResource

	return builder
}

// The behavior when clicking on a result
func (builder *VisualizationBuilder) DataLinks(links []any) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Links = links

	return builder
}

// Define interactive HTTP requests that can be triggered from data visualizations.
func (builder *VisualizationBuilder) Actions(actions []cog.Builder[dashboardv2beta1.Action]) *VisualizationBuilder {
	actionsResources := make([]dashboardv2beta1.Action, 0, len(actions))
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
func (builder *VisualizationBuilder) NoValue(noValue string) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Defaults.NoValue = &noValue

	return builder
}

// Overrides are the options applied to specific fields overriding the defaults.
func (builder *VisualizationBuilder) Overrides(overrides []cog.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]) *VisualizationBuilder {
	overridesResources := make([]dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides, 0, len(overrides))
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
func (builder *VisualizationBuilder) Override(matcher dashboardv2beta1.MatcherConfig, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher:    matcher,
		Properties: properties,
	})

	return builder
}

// Adds override rules for a specific field, referred to by its name.
func (builder *VisualizationBuilder) OverrideByName(name string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: dashboardv2beta1.MatcherConfig{
			Id:      "byName",
			Options: &name,
		},
		Properties: properties,
	})

	return builder
}

// Adds override rules for the fields whose name match the given regexp.
func (builder *VisualizationBuilder) OverrideByRegexp(regexp string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: dashboardv2beta1.MatcherConfig{
			Id:      "byRegexp",
			Options: &regexp,
		},
		Properties: properties,
	})

	return builder
}

// Adds override rules for all the fields of the given type.
func (builder *VisualizationBuilder) OverrideByFieldType(fieldType string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: dashboardv2beta1.MatcherConfig{
			Id:      "byType",
			Options: &fieldType,
		},
		Properties: properties,
	})

	return builder
}

func (builder *VisualizationBuilder) OverrideByQuery(queryRefId string, properties []dashboardv2beta1.DynamicConfigValue) *VisualizationBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: dashboardv2beta1.MatcherConfig{
			Id:      "byFrameRefID",
			Options: &queryRefId,
		},
		Properties: properties,
	})

	return builder
}

func (builder *VisualizationBuilder) Show(show XYShowMode) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).Show = &show

	return builder
}

func (builder *VisualizationBuilder) PointSize(pointSize cog.Builder[XychartFieldConfigPointSize]) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	pointSizeResource, err := pointSize.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).PointSize = &pointSizeResource

	return builder
}

func (builder *VisualizationBuilder) PointShape(pointShape PointShape) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).PointShape = &pointShape

	return builder
}

func (builder *VisualizationBuilder) PointStrokeWidth(pointStrokeWidth int32) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).PointStrokeWidth = &pointStrokeWidth

	return builder
}

func (builder *VisualizationBuilder) FillOpacity(fillOpacity uint32) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).FillOpacity = &fillOpacity

	return builder
}

func (builder *VisualizationBuilder) LineWidth(lineWidth int32) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).LineWidth = &lineWidth

	return builder
}

func (builder *VisualizationBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *VisualizationBuilder {
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

func (builder *VisualizationBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisPlacement = &axisPlacement

	return builder
}

func (builder *VisualizationBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisColorMode = &axisColorMode

	return builder
}

func (builder *VisualizationBuilder) AxisLabel(axisLabel string) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisLabel = &axisLabel

	return builder
}

func (builder *VisualizationBuilder) AxisWidth(axisWidth float64) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisWidth = &axisWidth

	return builder
}

func (builder *VisualizationBuilder) AxisSoftMin(axisSoftMin float64) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMin = &axisSoftMin

	return builder
}

func (builder *VisualizationBuilder) AxisSoftMax(axisSoftMax float64) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisSoftMax = &axisSoftMax

	return builder
}

func (builder *VisualizationBuilder) AxisGridShow(axisGridShow bool) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisGridShow = &axisGridShow

	return builder
}

func (builder *VisualizationBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *VisualizationBuilder {
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

func (builder *VisualizationBuilder) AxisCenteredZero(axisCenteredZero bool) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisCenteredZero = &axisCenteredZero

	return builder
}

func (builder *VisualizationBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *VisualizationBuilder {
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

func (builder *VisualizationBuilder) AxisBorderShow(axisBorderShow bool) *VisualizationBuilder {
	if builder.internal.Spec.FieldConfig.Defaults.Custom == nil {
		builder.internal.Spec.FieldConfig.Defaults.Custom = NewFieldConfig()
	}
	builder.internal.Spec.FieldConfig.Defaults.Custom.(*FieldConfig).AxisBorderShow = &axisBorderShow

	return builder
}

func (builder *VisualizationBuilder) Mapping(mapping SeriesMapping) *VisualizationBuilder {
	if builder.internal.Spec.Options == nil {
		builder.internal.Spec.Options = NewOptions()
	}
	builder.internal.Spec.Options.(*Options).Mapping = mapping

	return builder
}

func (builder *VisualizationBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *VisualizationBuilder {
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

func (builder *VisualizationBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *VisualizationBuilder {
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

func (builder *VisualizationBuilder) Series(series []cog.Builder[XYSeriesConfig]) *VisualizationBuilder {
	if builder.internal.Spec.Options == nil {
		builder.internal.Spec.Options = NewOptions()
	}
	seriesResources := make([]XYSeriesConfig, 0, len(series))
	for _, r1 := range series {
		seriesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		seriesResources = append(seriesResources, seriesDepth1)
	}
	builder.internal.Spec.Options.(*Options).Series = seriesResources

	return builder
}
