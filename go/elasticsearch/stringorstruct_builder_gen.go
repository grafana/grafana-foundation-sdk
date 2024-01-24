// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrStruct] = (*StringOrStructBuilder)(nil)

type StringOrStructBuilder struct {
	internal *StringOrStruct
	errors   map[string]cog.BuildErrors
}

func NewStringOrStructBuilder() *StringOrStructBuilder {
	resource := &StringOrStruct{}
	builder := &StringOrStructBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *StringOrStructBuilder) Build() (StringOrStruct, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("StringOrStruct", err)...)
	}

	if len(errs) != 0 {
		return StringOrStruct{}, errs
	}

	return *builder.internal, nil
}

func (builder *StringOrStructBuilder) String(string string) *StringOrStructBuilder {
	builder.internal.String = &string

	return builder
}

func (builder *StringOrStructBuilder) Struct(structArg struct {
	Inline string `json:"inline,omitempty"`
}) *StringOrStructBuilder {
	builder.internal.Struct = &structArg

	return builder
}

func (builder *StringOrStructBuilder) applyDefaults() {
}
