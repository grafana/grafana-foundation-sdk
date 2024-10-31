// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeReduceSettings] = (*ExprTypeReduceSettingsBuilder)(nil)

type ExprTypeReduceSettingsBuilder struct {
	internal *ExprTypeReduceSettings
	errors   map[string]cog.BuildErrors
}

func NewExprTypeReduceSettingsBuilder() *ExprTypeReduceSettingsBuilder {
	resource := &ExprTypeReduceSettings{}
	builder := &ExprTypeReduceSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeReduceSettingsBuilder) Build() (ExprTypeReduceSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeReduceSettings{}, err
	}

	return *builder.internal, nil
}

// Non-number reduce behavior
// Possible enum values:
//   - `"dropNN"` Drop non-numbers
//   - `"replaceNN"` Replace non-numbers
func (builder *ExprTypeReduceSettingsBuilder) Mode(mode TypeReduceMode) *ExprTypeReduceSettingsBuilder {
	builder.internal.Mode = mode

	return builder
}

// Only valid when mode is replace
func (builder *ExprTypeReduceSettingsBuilder) ReplaceWithValue(replaceWithValue float64) *ExprTypeReduceSettingsBuilder {
	builder.internal.ReplaceWithValue = &replaceWithValue

	return builder
}

func (builder *ExprTypeReduceSettingsBuilder) applyDefaults() {
}
