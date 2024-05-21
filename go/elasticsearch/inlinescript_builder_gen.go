// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[InlineScript] = (*InlineScriptBuilder)(nil)

type InlineScriptBuilder struct {
	internal *InlineScript
	errors   map[string]cog.BuildErrors
}

func NewInlineScriptBuilder() *InlineScriptBuilder {
	resource := &InlineScript{}
	builder := &InlineScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *InlineScriptBuilder) Build() (InlineScript, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("InlineScript", err)...)
	}

	if len(errs) != 0 {
		return InlineScript{}, errs
	}

	return *builder.internal, nil
}

func (builder *InlineScriptBuilder) String(string string) *InlineScriptBuilder {
	builder.internal.String = &string

	return builder
}

func (builder *InlineScriptBuilder) Struct(structArg struct {
	Inline *string `json:"inline,omitempty"`
}) *InlineScriptBuilder {
	builder.internal.Struct = &structArg

	return builder
}

func (builder *InlineScriptBuilder) applyDefaults() {
}
