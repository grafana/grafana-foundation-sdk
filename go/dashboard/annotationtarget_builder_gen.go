// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationTarget] = (*AnnotationTargetBuilder)(nil)

// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most dashboards
type AnnotationTargetBuilder struct {
	internal *AnnotationTarget
	errors   cog.BuildErrors
}

func NewAnnotationTargetBuilder() *AnnotationTargetBuilder {
	resource := NewAnnotationTarget()
	builder := &AnnotationTargetBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AnnotationTargetBuilder) Build() (AnnotationTarget, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationTarget{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationTarget{}, cog.MakeBuildErrors("dashboard.annotationTarget", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AnnotationTargetBuilder) RecordError(path string, err error) *AnnotationTargetBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func (builder *AnnotationTargetBuilder) Limit(limit int64) *AnnotationTargetBuilder {
	builder.internal.Limit = limit

	return builder
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func (builder *AnnotationTargetBuilder) MatchAny(matchAny bool) *AnnotationTargetBuilder {
	builder.internal.MatchAny = matchAny

	return builder
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func (builder *AnnotationTargetBuilder) Tags(tags []string) *AnnotationTargetBuilder {
	builder.internal.Tags = tags

	return builder
}

// Only required/valid for the grafana datasource...
// but code+tests is already depending on it so hard to change
func (builder *AnnotationTargetBuilder) Type(typeArg string) *AnnotationTargetBuilder {
	builder.internal.Type = typeArg

	return builder
}
