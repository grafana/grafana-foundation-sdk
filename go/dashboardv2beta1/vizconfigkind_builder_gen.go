// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VizConfigKind] = (*VizConfigKindBuilder)(nil)

type VizConfigKindBuilder struct {
	internal *VizConfigKind
	errors   cog.BuildErrors
}

func NewVizConfigKindBuilder() *VizConfigKindBuilder {
	resource := NewVizConfigKind()
	builder := &VizConfigKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "VizConfig"

	return builder
}

func (builder *VizConfigKindBuilder) Build() (VizConfigKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return VizConfigKind{}, err
	}

	if len(builder.errors) > 0 {
		return VizConfigKind{}, cog.MakeBuildErrors("dashboardv2beta1.vizConfigKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VizConfigKindBuilder) RecordError(path string, err error) *VizConfigKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The group is the plugin ID
func (builder *VizConfigKindBuilder) Group(group string) *VizConfigKindBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *VizConfigKindBuilder) Version(version string) *VizConfigKindBuilder {
	builder.internal.Version = version

	return builder
}

func (builder *VizConfigKindBuilder) Spec(spec VizConfigSpec) *VizConfigKindBuilder {
	builder.internal.Spec = spec

	return builder
}

func (builder *VizConfigKindBuilder) Options(options any) *VizConfigKindBuilder {
	builder.internal.Spec.Options = &options

	return builder
}

func (builder *VizConfigKindBuilder) FieldConfig(fieldConfig FieldConfigSource) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig = fieldConfig

	return builder
}

// The display value for this field.  This supports template variables blank is auto
func (builder *VizConfigKindBuilder) DisplayName(displayName string) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.DisplayName = &displayName

	return builder
}

// This can be used by data sources that return and explicit naming structure for values and labels
// When this property is configured, this value is used rather than the default naming strategy.
func (builder *VizConfigKindBuilder) DisplayNameFromDS(displayNameFromDS string) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.DisplayNameFromDS = &displayNameFromDS

	return builder
}

// Human readable field metadata
func (builder *VizConfigKindBuilder) Description(description string) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Description = &description

	return builder
}

// An explicit path to the field in the datasource.  When the frame meta includes a path,
// This will default to `${frame.meta.path}/${field.name}
//
// When defined, this value can be used as an identifier within the datasource scope, and
// may be used to update the results
func (builder *VizConfigKindBuilder) Path(path string) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Path = &path

	return builder
}

// True if data source can write a value to the path. Auth/authz are supported separately
func (builder *VizConfigKindBuilder) Writeable(writeable bool) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Writeable = &writeable

	return builder
}

// True if data source field supports ad-hoc filters
func (builder *VizConfigKindBuilder) Filterable(filterable bool) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Filterable = &filterable

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
func (builder *VizConfigKindBuilder) Unit(unit string) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Unit = &unit

	return builder
}

// Specify the number of decimals Grafana includes in the rendered value.
// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
// To display all decimals, set the unit to `String`.
func (builder *VizConfigKindBuilder) Decimals(decimals float64) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Decimals = &decimals

	return builder
}

// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *VizConfigKindBuilder) Min(min float64) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Min = &min

	return builder
}

// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *VizConfigKindBuilder) Max(max float64) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Max = &max

	return builder
}

// Convert input values into a display string
func (builder *VizConfigKindBuilder) Mappings(mappings []ValueMapping) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Mappings = mappings

	return builder
}

// Map numeric values to states
func (builder *VizConfigKindBuilder) Thresholds(thresholds cog.Builder[ThresholdsConfig]) *VizConfigKindBuilder {
	thresholdsResource, err := thresholds.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Thresholds = &thresholdsResource

	return builder
}

// Panel color configuration
func (builder *VizConfigKindBuilder) ColorScheme(color cog.Builder[FieldColor]) *VizConfigKindBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.FieldConfig.Defaults.Color = &colorResource

	return builder
}

// The behavior when clicking on a result
func (builder *VizConfigKindBuilder) DataLinks(links []any) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Links = links

	return builder
}

// Define interactive HTTP requests that can be triggered from data visualizations.
func (builder *VizConfigKindBuilder) Actions(actions []cog.Builder[Action]) *VizConfigKindBuilder {
	actionsResources := make([]Action, 0, len(actions))
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
func (builder *VizConfigKindBuilder) NoValue(noValue string) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.NoValue = &noValue

	return builder
}

// custom is specified by the FieldConfig field
// in panel plugin schemas.
func (builder *VizConfigKindBuilder) Custom(custom any) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults.Custom = &custom

	return builder
}

// Defaults are the options applied to all fields.
func (builder *VizConfigKindBuilder) Defaults(defaults FieldConfig) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Defaults = defaults

	return builder
}

// Overrides are the options applied to specific fields overriding the defaults.
func (builder *VizConfigKindBuilder) Overrides(overrides []cog.Builder[Dashboardv2beta1FieldConfigSourceOverrides]) *VizConfigKindBuilder {
	overridesResources := make([]Dashboardv2beta1FieldConfigSourceOverrides, 0, len(overrides))
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
func (builder *VizConfigKindBuilder) Override(matcher MatcherConfig, properties []DynamicConfigValue) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher:    matcher,
		Properties: properties,
	})

	return builder
}

// Adds override rules for a specific field, referred to by its name.
func (builder *VizConfigKindBuilder) OverrideByName(name string, properties []DynamicConfigValue) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: MatcherConfig{
			Id:      "byName",
			Options: &name,
		},
		Properties: properties,
	})

	return builder
}

// Adds override rules for the fields whose name match the given regexp.
func (builder *VizConfigKindBuilder) OverrideByRegexp(regexp string, properties []DynamicConfigValue) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: MatcherConfig{
			Id:      "byRegexp",
			Options: &regexp,
		},
		Properties: properties,
	})

	return builder
}

// Adds override rules for all the fields of the given type.
func (builder *VizConfigKindBuilder) OverrideByFieldType(fieldType string, properties []DynamicConfigValue) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: MatcherConfig{
			Id:      "byType",
			Options: &fieldType,
		},
		Properties: properties,
	})

	return builder
}

func (builder *VizConfigKindBuilder) OverrideByQuery(queryRefId string, properties []DynamicConfigValue) *VizConfigKindBuilder {
	builder.internal.Spec.FieldConfig.Overrides = append(builder.internal.Spec.FieldConfig.Overrides, Dashboardv2beta1FieldConfigSourceOverrides{
		Matcher: MatcherConfig{
			Id:      "byFrameRefID",
			Options: &queryRefId,
		},
		Properties: properties,
	})

	return builder
}
