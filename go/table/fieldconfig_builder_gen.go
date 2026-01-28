// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package table

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[FieldConfig] = (*FieldConfigBuilder)(nil)

type FieldConfigBuilder struct {
	internal *FieldConfig
	errors   cog.BuildErrors
}

func NewFieldConfigBuilder() *FieldConfigBuilder {
	resource := NewFieldConfig()
	builder := &FieldConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FieldConfigBuilder) Build() (FieldConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return FieldConfig{}, err
	}

	if len(builder.errors) > 0 {
		return FieldConfig{}, cog.MakeBuildErrors("table.fieldConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FieldConfigBuilder) Width(width float64) *FieldConfigBuilder {
	builder.internal.Width = &width

	return builder
}

func (builder *FieldConfigBuilder) MinWidth(minWidth float64) *FieldConfigBuilder {
	builder.internal.MinWidth = &minWidth

	return builder
}

func (builder *FieldConfigBuilder) Align(align common.FieldTextAlignment) *FieldConfigBuilder {
	builder.internal.Align = align

	return builder
}

// This field is deprecated in favor of using cellOptions
func (builder *FieldConfigBuilder) DisplayMode(displayMode common.TableCellDisplayMode) *FieldConfigBuilder {
	builder.internal.DisplayMode = &displayMode

	return builder
}

func (builder *FieldConfigBuilder) CellOptions(cellOptions common.TableCellOptions) *FieldConfigBuilder {
	builder.internal.CellOptions = &cellOptions

	return builder
}

// ?? default is missing or false ??
func (builder *FieldConfigBuilder) Hidden(hidden bool) *FieldConfigBuilder {
	builder.internal.Hidden = &hidden

	return builder
}

func (builder *FieldConfigBuilder) Inspect(inspect bool) *FieldConfigBuilder {
	builder.internal.Inspect = inspect

	return builder
}

func (builder *FieldConfigBuilder) Filterable(filterable bool) *FieldConfigBuilder {
	builder.internal.Filterable = &filterable

	return builder
}

// Hides any header for a column, useful for columns that show some static content or buttons.
func (builder *FieldConfigBuilder) HideHeader(hideHeader bool) *FieldConfigBuilder {
	builder.internal.HideHeader = &hideHeader

	return builder
}
