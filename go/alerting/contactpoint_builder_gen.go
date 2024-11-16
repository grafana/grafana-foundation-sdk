// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ContactPoint] = (*ContactPointBuilder)(nil)

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
type ContactPointBuilder struct {
	internal *ContactPoint
	errors   map[string]cog.BuildErrors
}

func NewContactPointBuilder() *ContactPointBuilder {
	resource := NewContactPoint()
	builder := &ContactPointBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ContactPointBuilder) Build() (ContactPoint, error) {
	if err := builder.internal.Validate(); err != nil {
		return ContactPoint{}, err
	}

	return *builder.internal, nil
}

func (builder *ContactPointBuilder) DisableResolveMessage(disableResolveMessage bool) *ContactPointBuilder {
	builder.internal.DisableResolveMessage = &disableResolveMessage

	return builder
}

// Name is used as grouping key in the UI. Contact points with the
// same name will be grouped in the UI.
func (builder *ContactPointBuilder) Name(name string) *ContactPointBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *ContactPointBuilder) Provenance(provenance string) *ContactPointBuilder {
	builder.internal.Provenance = &provenance

	return builder
}

func (builder *ContactPointBuilder) Settings(settings Json) *ContactPointBuilder {
	builder.internal.Settings = settings

	return builder
}

func (builder *ContactPointBuilder) Type(typeArg ContactPointType) *ContactPointBuilder {
	builder.internal.Type = typeArg

	return builder
}

// UID is the unique identifier of the contact point. The UID can be
// set by the user.
func (builder *ContactPointBuilder) Uid(uid string) *ContactPointBuilder {
	builder.internal.Uid = &uid

	return builder
}
