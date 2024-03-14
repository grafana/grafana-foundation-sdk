// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationContainer] = (*AnnotationContainerBuilder)(nil)

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
type AnnotationContainerBuilder struct {
	internal *AnnotationContainer
	errors   map[string]cog.BuildErrors
}

func NewAnnotationContainerBuilder() *AnnotationContainerBuilder {
	resource := &AnnotationContainer{}
	builder := &AnnotationContainerBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AnnotationContainerBuilder) Build() (AnnotationContainer, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("AnnotationContainer", err)...)
	}

	if len(errs) != 0 {
		return AnnotationContainer{}, errs
	}

	return *builder.internal, nil
}

// List of annotations
func (builder *AnnotationContainerBuilder) List(list []cog.Builder[AnnotationQuery]) *AnnotationContainerBuilder {
	listResources := make([]AnnotationQuery, 0, len(list))
	for _, r1 := range list {
		listDepth1, err := r1.Build()
		if err != nil {
			builder.errors["list"] = err.(cog.BuildErrors)
			return builder
		}
		listResources = append(listResources, listDepth1)
	}
	builder.internal.List = listResources

	return builder
}

func (builder *AnnotationContainerBuilder) applyDefaults() {
}
